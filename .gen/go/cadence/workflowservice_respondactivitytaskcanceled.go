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

// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package cadence

import (
	"errors"
	"fmt"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// WorkflowService_RespondActivityTaskCanceled_Args represents the arguments for the WorkflowService.RespondActivityTaskCanceled function.
//
// The arguments for RespondActivityTaskCanceled are sent and received over the wire as this struct.
type WorkflowService_RespondActivityTaskCanceled_Args struct {
	CanceledRequest *shared.RespondActivityTaskCanceledRequest `json:"canceledRequest,omitempty"`
}

// ToWire translates a WorkflowService_RespondActivityTaskCanceled_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_RespondActivityTaskCanceled_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.CanceledRequest != nil {
		w, err = v.CanceledRequest.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RespondActivityTaskCanceledRequest_Read(w wire.Value) (*shared.RespondActivityTaskCanceledRequest, error) {
	var v shared.RespondActivityTaskCanceledRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a WorkflowService_RespondActivityTaskCanceled_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RespondActivityTaskCanceled_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RespondActivityTaskCanceled_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RespondActivityTaskCanceled_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.CanceledRequest, err = _RespondActivityTaskCanceledRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RespondActivityTaskCanceled_Args
// struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.CanceledRequest != nil {
		fields[i] = fmt.Sprintf("CanceledRequest: %v", v.CanceledRequest)
		i++
	}

	return fmt.Sprintf("WorkflowService_RespondActivityTaskCanceled_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RespondActivityTaskCanceled_Args match the
// provided WorkflowService_RespondActivityTaskCanceled_Args.
//
// This function performs a deep comparison.
func (v *WorkflowService_RespondActivityTaskCanceled_Args) Equals(rhs *WorkflowService_RespondActivityTaskCanceled_Args) bool {
	if !((v.CanceledRequest == nil && rhs.CanceledRequest == nil) || (v.CanceledRequest != nil && rhs.CanceledRequest != nil && v.CanceledRequest.Equals(rhs.CanceledRequest))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "RespondActivityTaskCanceled" for this struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Args) MethodName() string {
	return "RespondActivityTaskCanceled"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// WorkflowService_RespondActivityTaskCanceled_Helper provides functions that aid in handling the
// parameters and return values of the WorkflowService.RespondActivityTaskCanceled
// function.
var WorkflowService_RespondActivityTaskCanceled_Helper = struct {
	// Args accepts the parameters of RespondActivityTaskCanceled in-order and returns
	// the arguments struct for the function.
	Args func(
		canceledRequest *shared.RespondActivityTaskCanceledRequest,
	) *WorkflowService_RespondActivityTaskCanceled_Args

	// IsException returns true if the given error can be thrown
	// by RespondActivityTaskCanceled.
	//
	// An error can be thrown by RespondActivityTaskCanceled only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for RespondActivityTaskCanceled
	// given the error returned by it. The provided error may
	// be nil if RespondActivityTaskCanceled did not fail.
	//
	// This allows mapping errors returned by RespondActivityTaskCanceled into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// RespondActivityTaskCanceled
	//
	//   err := RespondActivityTaskCanceled(args)
	//   result, err := WorkflowService_RespondActivityTaskCanceled_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from RespondActivityTaskCanceled: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*WorkflowService_RespondActivityTaskCanceled_Result, error)

	// UnwrapResponse takes the result struct for RespondActivityTaskCanceled
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if RespondActivityTaskCanceled threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := WorkflowService_RespondActivityTaskCanceled_Helper.UnwrapResponse(result)
	UnwrapResponse func(*WorkflowService_RespondActivityTaskCanceled_Result) error
}{}

func init() {
	WorkflowService_RespondActivityTaskCanceled_Helper.Args = func(
		canceledRequest *shared.RespondActivityTaskCanceledRequest,
	) *WorkflowService_RespondActivityTaskCanceled_Args {
		return &WorkflowService_RespondActivityTaskCanceled_Args{
			CanceledRequest: canceledRequest,
		}
	}

	WorkflowService_RespondActivityTaskCanceled_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *shared.BadRequestError:
			return true
		case *shared.EntityNotExistsError:
			return true
		case *shared.DomainNotActiveError:
			return true
		case *shared.LimitExceededError:
			return true
		case *shared.ServiceBusyError:
			return true
		case *shared.ClientVersionNotSupportedError:
			return true
		default:
			return false
		}
	}

	WorkflowService_RespondActivityTaskCanceled_Helper.WrapResponse = func(err error) (*WorkflowService_RespondActivityTaskCanceled_Result, error) {
		if err == nil {
			return &WorkflowService_RespondActivityTaskCanceled_Result{}, nil
		}

		switch e := err.(type) {
		case *shared.BadRequestError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.BadRequestError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{BadRequestError: e}, nil
		case *shared.EntityNotExistsError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.EntityNotExistError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{EntityNotExistError: e}, nil
		case *shared.DomainNotActiveError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.DomainNotActiveError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{DomainNotActiveError: e}, nil
		case *shared.LimitExceededError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.LimitExceededError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{LimitExceededError: e}, nil
		case *shared.ServiceBusyError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.ServiceBusyError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{ServiceBusyError: e}, nil
		case *shared.ClientVersionNotSupportedError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for WorkflowService_RespondActivityTaskCanceled_Result.ClientVersionNotSupportedError")
			}
			return &WorkflowService_RespondActivityTaskCanceled_Result{ClientVersionNotSupportedError: e}, nil
		}

		return nil, err
	}
	WorkflowService_RespondActivityTaskCanceled_Helper.UnwrapResponse = func(result *WorkflowService_RespondActivityTaskCanceled_Result) (err error) {
		if result.BadRequestError != nil {
			err = result.BadRequestError
			return
		}
		if result.EntityNotExistError != nil {
			err = result.EntityNotExistError
			return
		}
		if result.DomainNotActiveError != nil {
			err = result.DomainNotActiveError
			return
		}
		if result.LimitExceededError != nil {
			err = result.LimitExceededError
			return
		}
		if result.ServiceBusyError != nil {
			err = result.ServiceBusyError
			return
		}
		if result.ClientVersionNotSupportedError != nil {
			err = result.ClientVersionNotSupportedError
			return
		}
		return
	}

}

// WorkflowService_RespondActivityTaskCanceled_Result represents the result of a WorkflowService.RespondActivityTaskCanceled function call.
//
// The result of a RespondActivityTaskCanceled execution is sent and received over the wire as this struct.
type WorkflowService_RespondActivityTaskCanceled_Result struct {
	BadRequestError                *shared.BadRequestError                `json:"badRequestError,omitempty"`
	EntityNotExistError            *shared.EntityNotExistsError           `json:"entityNotExistError,omitempty"`
	DomainNotActiveError           *shared.DomainNotActiveError           `json:"domainNotActiveError,omitempty"`
	LimitExceededError             *shared.LimitExceededError             `json:"limitExceededError,omitempty"`
	ServiceBusyError               *shared.ServiceBusyError               `json:"serviceBusyError,omitempty"`
	ClientVersionNotSupportedError *shared.ClientVersionNotSupportedError `json:"clientVersionNotSupportedError,omitempty"`
}

// ToWire translates a WorkflowService_RespondActivityTaskCanceled_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *WorkflowService_RespondActivityTaskCanceled_Result) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.BadRequestError != nil {
		w, err = v.BadRequestError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.EntityNotExistError != nil {
		w, err = v.EntityNotExistError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.DomainNotActiveError != nil {
		w, err = v.DomainNotActiveError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.LimitExceededError != nil {
		w, err = v.LimitExceededError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.ServiceBusyError != nil {
		w, err = v.ServiceBusyError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if v.ClientVersionNotSupportedError != nil {
		w, err = v.ClientVersionNotSupportedError.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 7, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("WorkflowService_RespondActivityTaskCanceled_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a WorkflowService_RespondActivityTaskCanceled_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a WorkflowService_RespondActivityTaskCanceled_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v WorkflowService_RespondActivityTaskCanceled_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *WorkflowService_RespondActivityTaskCanceled_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BadRequestError, err = _BadRequestError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.EntityNotExistError, err = _EntityNotExistsError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.DomainNotActiveError, err = _DomainNotActiveError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.LimitExceededError, err = _LimitExceededError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.ServiceBusyError, err = _ServiceBusyError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TStruct {
				v.ClientVersionNotSupportedError, err = _ClientVersionNotSupportedError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.BadRequestError != nil {
		count++
	}
	if v.EntityNotExistError != nil {
		count++
	}
	if v.DomainNotActiveError != nil {
		count++
	}
	if v.LimitExceededError != nil {
		count++
	}
	if v.ServiceBusyError != nil {
		count++
	}
	if v.ClientVersionNotSupportedError != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("WorkflowService_RespondActivityTaskCanceled_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a WorkflowService_RespondActivityTaskCanceled_Result
// struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [6]string
	i := 0
	if v.BadRequestError != nil {
		fields[i] = fmt.Sprintf("BadRequestError: %v", v.BadRequestError)
		i++
	}
	if v.EntityNotExistError != nil {
		fields[i] = fmt.Sprintf("EntityNotExistError: %v", v.EntityNotExistError)
		i++
	}
	if v.DomainNotActiveError != nil {
		fields[i] = fmt.Sprintf("DomainNotActiveError: %v", v.DomainNotActiveError)
		i++
	}
	if v.LimitExceededError != nil {
		fields[i] = fmt.Sprintf("LimitExceededError: %v", v.LimitExceededError)
		i++
	}
	if v.ServiceBusyError != nil {
		fields[i] = fmt.Sprintf("ServiceBusyError: %v", v.ServiceBusyError)
		i++
	}
	if v.ClientVersionNotSupportedError != nil {
		fields[i] = fmt.Sprintf("ClientVersionNotSupportedError: %v", v.ClientVersionNotSupportedError)
		i++
	}

	return fmt.Sprintf("WorkflowService_RespondActivityTaskCanceled_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this WorkflowService_RespondActivityTaskCanceled_Result match the
// provided WorkflowService_RespondActivityTaskCanceled_Result.
//
// This function performs a deep comparison.
func (v *WorkflowService_RespondActivityTaskCanceled_Result) Equals(rhs *WorkflowService_RespondActivityTaskCanceled_Result) bool {
	if !((v.BadRequestError == nil && rhs.BadRequestError == nil) || (v.BadRequestError != nil && rhs.BadRequestError != nil && v.BadRequestError.Equals(rhs.BadRequestError))) {
		return false
	}
	if !((v.EntityNotExistError == nil && rhs.EntityNotExistError == nil) || (v.EntityNotExistError != nil && rhs.EntityNotExistError != nil && v.EntityNotExistError.Equals(rhs.EntityNotExistError))) {
		return false
	}
	if !((v.DomainNotActiveError == nil && rhs.DomainNotActiveError == nil) || (v.DomainNotActiveError != nil && rhs.DomainNotActiveError != nil && v.DomainNotActiveError.Equals(rhs.DomainNotActiveError))) {
		return false
	}
	if !((v.LimitExceededError == nil && rhs.LimitExceededError == nil) || (v.LimitExceededError != nil && rhs.LimitExceededError != nil && v.LimitExceededError.Equals(rhs.LimitExceededError))) {
		return false
	}
	if !((v.ServiceBusyError == nil && rhs.ServiceBusyError == nil) || (v.ServiceBusyError != nil && rhs.ServiceBusyError != nil && v.ServiceBusyError.Equals(rhs.ServiceBusyError))) {
		return false
	}
	if !((v.ClientVersionNotSupportedError == nil && rhs.ClientVersionNotSupportedError == nil) || (v.ClientVersionNotSupportedError != nil && rhs.ClientVersionNotSupportedError != nil && v.ClientVersionNotSupportedError.Equals(rhs.ClientVersionNotSupportedError))) {
		return false
	}

	return true
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "RespondActivityTaskCanceled" for this struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Result) MethodName() string {
	return "RespondActivityTaskCanceled"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *WorkflowService_RespondActivityTaskCanceled_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
