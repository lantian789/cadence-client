// Copyright (c) 2017-2021 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package internal

import (
	"context"
	"errors"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shadower"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/internal/common"
	"go.uber.org/cadence/internal/common/backoff"
	"go.uber.org/zap"
)

type (
	shadowWorker struct {
		activityWorker *activityWorker

		service  workflowserviceclient.Interface
		domain   string
		taskList string
		options  *ShadowOptions
		logger   *zap.Logger
	}
)

func newShadowWorker(
	service workflowserviceclient.Interface,
	domain string,
	shadowOptions *ShadowOptions,
	params workerExecutionParameters,
	registry *registry,
) *shadowWorker {
	registry.RegisterActivityWithOptions(scanWorkflowActivity, RegisterActivityOptions{
		Name: shadower.ScanWorkflowActivityName,
	})
	registry.RegisterActivityWithOptions(replayWorkflowActivity, RegisterActivityOptions{
		Name: shadower.ReplayWorkflowActivityName,
	})

	replayer := NewWorkflowReplayer()
	replayer.registry = registry

	params.UserContext = context.WithValue(params.UserContext, serviceClientContextKey, service)
	params.UserContext = context.WithValue(params.UserContext, workflowReplayerContextKey, replayer)

	activityWorker := newActivityWorker(
		service,
		shadower.LocalDomainName, // note: this is the system domain for all shadow workflows
		params,
		nil,
		registry,
		nil,
	)
	return &shadowWorker{
		activityWorker: activityWorker,

		service:  service,
		domain:   domain,
		taskList: params.TaskList,
		options:  shadowOptions,
		logger:   params.Logger,
	}
}

func (sw *shadowWorker) Start() error {
	if sw.options == nil {
		return errors.New("shadowerOptions must be specified when shadow worker is enabled")
	}

	if err := sw.options.validateAndPopulateFields(); err != nil {
		return err
	}

	if err := verifyDomainExist(sw.service, sw.domain, sw.logger); err != nil {
		return err
	}

	if len(sw.taskList) == 0 {
		return errTaskListNotSet
	}

	if err := sw.startShadowWorkflow(); err != nil {
		return err
	}

	return sw.activityWorker.Start()
}

func (sw *shadowWorker) Stop() {
	sw.activityWorker.Stop()
}

func (sw *shadowWorker) startShadowWorkflow() error {
	workflowParams := shadower.WorkflowParams{
		Domain:        common.StringPtr(sw.domain),
		TaskList:      common.StringPtr(sw.taskList),
		WorkflowQuery: common.StringPtr(sw.options.WorkflowQuery),
		SamplingRate:  common.Float64Ptr(sw.options.SamplingRate),
		ShadowMode:    sw.options.Mode.toThriftPtr(),
		ExitCondition: sw.options.ExitCondition.toThriftPtr(),
		Concurrency:   common.Int32Ptr(int32(sw.options.Concurrency)),
	}

	ctx := context.Background()

	workflowType, input, err := getValidatedWorkflowFunction(shadower.WorkflowName, []interface{}{workflowParams}, getDefaultDataConverter(), nil)
	if err != nil {
		return err
	}

	startWorkflowRequest := &shared.StartWorkflowExecutionRequest{
		Domain:       common.StringPtr(shadower.LocalDomainName),
		WorkflowId:   common.StringPtr(sw.domain + shadower.WorkflowIDSuffix),
		WorkflowType: workflowTypePtr(*workflowType),
		TaskList: &shared.TaskList{
			Name: common.StringPtr(shadower.TaskList),
		},
		Input:                               input,
		ExecutionStartToCloseTimeoutSeconds: common.Int32Ptr(864000),
		TaskStartToCloseTimeoutSeconds:      common.Int32Ptr(60),
		RequestId:                           common.StringPtr(uuid.New()),
		WorkflowIdReusePolicy:               shared.WorkflowIdReusePolicyAllowDuplicate.Ptr(),
	}

	startWorkflowOp := func() error {
		tchCtx, cancel, opt := newChannelContext(ctx)
		defer cancel()
		_, err := sw.service.StartWorkflowExecution(tchCtx, startWorkflowRequest, opt...)
		if err != nil {
			if _, ok := err.(*shared.WorkflowExecutionAlreadyStartedError); ok {
				return nil
			}
		}

		return err
	}

	return backoff.Retry(ctx, startWorkflowOp, createDynamicServiceRetryPolicy(ctx), isServiceTransientError)
}
