// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	catalog "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/catalog"

	handler "github.com/flyteorg/flytepropeller/pkg/controller/nodes/handler"

	interfaces "github.com/flyteorg/flytepropeller/pkg/controller/nodes/interfaces"

	io "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/io"

	ioutils "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/ioutils"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// TaskNodeHandler is an autogenerated mock type for the TaskNodeHandler type
type TaskNodeHandler struct {
	mock.Mock
}

type TaskNodeHandler_Abort struct {
	*mock.Call
}

func (_m TaskNodeHandler_Abort) Return(_a0 error) *TaskNodeHandler_Abort {
	return &TaskNodeHandler_Abort{Call: _m.Call.Return(_a0)}
}

func (_m *TaskNodeHandler) OnAbort(ctx context.Context, executionContext interfaces.NodeExecutionContext, reason string) *TaskNodeHandler_Abort {
	c_call := _m.On("Abort", ctx, executionContext, reason)
	return &TaskNodeHandler_Abort{Call: c_call}
}

func (_m *TaskNodeHandler) OnAbortMatch(matchers ...interface{}) *TaskNodeHandler_Abort {
	c_call := _m.On("Abort", matchers...)
	return &TaskNodeHandler_Abort{Call: c_call}
}

// Abort provides a mock function with given fields: ctx, executionContext, reason
func (_m *TaskNodeHandler) Abort(ctx context.Context, executionContext interfaces.NodeExecutionContext, reason string) error {
	ret := _m.Called(ctx, executionContext, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.NodeExecutionContext, string) error); ok {
		r0 = rf(ctx, executionContext, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type TaskNodeHandler_Finalize struct {
	*mock.Call
}

func (_m TaskNodeHandler_Finalize) Return(_a0 error) *TaskNodeHandler_Finalize {
	return &TaskNodeHandler_Finalize{Call: _m.Call.Return(_a0)}
}

func (_m *TaskNodeHandler) OnFinalize(ctx context.Context, executionContext interfaces.NodeExecutionContext) *TaskNodeHandler_Finalize {
	c_call := _m.On("Finalize", ctx, executionContext)
	return &TaskNodeHandler_Finalize{Call: c_call}
}

func (_m *TaskNodeHandler) OnFinalizeMatch(matchers ...interface{}) *TaskNodeHandler_Finalize {
	c_call := _m.On("Finalize", matchers...)
	return &TaskNodeHandler_Finalize{Call: c_call}
}

// Finalize provides a mock function with given fields: ctx, executionContext
func (_m *TaskNodeHandler) Finalize(ctx context.Context, executionContext interfaces.NodeExecutionContext) error {
	ret := _m.Called(ctx, executionContext)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.NodeExecutionContext) error); ok {
		r0 = rf(ctx, executionContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type TaskNodeHandler_FinalizeRequired struct {
	*mock.Call
}

func (_m TaskNodeHandler_FinalizeRequired) Return(_a0 bool) *TaskNodeHandler_FinalizeRequired {
	return &TaskNodeHandler_FinalizeRequired{Call: _m.Call.Return(_a0)}
}

func (_m *TaskNodeHandler) OnFinalizeRequired() *TaskNodeHandler_FinalizeRequired {
	c_call := _m.On("FinalizeRequired")
	return &TaskNodeHandler_FinalizeRequired{Call: c_call}
}

func (_m *TaskNodeHandler) OnFinalizeRequiredMatch(matchers ...interface{}) *TaskNodeHandler_FinalizeRequired {
	c_call := _m.On("FinalizeRequired", matchers...)
	return &TaskNodeHandler_FinalizeRequired{Call: c_call}
}

// FinalizeRequired provides a mock function with given fields:
func (_m *TaskNodeHandler) FinalizeRequired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type TaskNodeHandler_Handle struct {
	*mock.Call
}

func (_m TaskNodeHandler_Handle) Return(_a0 handler.Transition, _a1 error) *TaskNodeHandler_Handle {
	return &TaskNodeHandler_Handle{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskNodeHandler) OnHandle(ctx context.Context, executionContext interfaces.NodeExecutionContext) *TaskNodeHandler_Handle {
	c_call := _m.On("Handle", ctx, executionContext)
	return &TaskNodeHandler_Handle{Call: c_call}
}

func (_m *TaskNodeHandler) OnHandleMatch(matchers ...interface{}) *TaskNodeHandler_Handle {
	c_call := _m.On("Handle", matchers...)
	return &TaskNodeHandler_Handle{Call: c_call}
}

// Handle provides a mock function with given fields: ctx, executionContext
func (_m *TaskNodeHandler) Handle(ctx context.Context, executionContext interfaces.NodeExecutionContext) (handler.Transition, error) {
	ret := _m.Called(ctx, executionContext)

	var r0 handler.Transition
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.NodeExecutionContext) handler.Transition); ok {
		r0 = rf(ctx, executionContext)
	} else {
		r0 = ret.Get(0).(handler.Transition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.NodeExecutionContext) error); ok {
		r1 = rf(ctx, executionContext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type TaskNodeHandler_Setup struct {
	*mock.Call
}

func (_m TaskNodeHandler_Setup) Return(_a0 error) *TaskNodeHandler_Setup {
	return &TaskNodeHandler_Setup{Call: _m.Call.Return(_a0)}
}

func (_m *TaskNodeHandler) OnSetup(ctx context.Context, setupContext interfaces.SetupContext) *TaskNodeHandler_Setup {
	c_call := _m.On("Setup", ctx, setupContext)
	return &TaskNodeHandler_Setup{Call: c_call}
}

func (_m *TaskNodeHandler) OnSetupMatch(matchers ...interface{}) *TaskNodeHandler_Setup {
	c_call := _m.On("Setup", matchers...)
	return &TaskNodeHandler_Setup{Call: c_call}
}

// Setup provides a mock function with given fields: ctx, setupContext
func (_m *TaskNodeHandler) Setup(ctx context.Context, setupContext interfaces.SetupContext) error {
	ret := _m.Called(ctx, setupContext)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.SetupContext) error); ok {
		r0 = rf(ctx, setupContext)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type TaskNodeHandler_ValidateOutputAndCacheAdd struct {
	*mock.Call
}

func (_m TaskNodeHandler_ValidateOutputAndCacheAdd) Return(_a0 catalog.Status, _a1 *io.ExecutionError, _a2 error) *TaskNodeHandler_ValidateOutputAndCacheAdd {
	return &TaskNodeHandler_ValidateOutputAndCacheAdd{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *TaskNodeHandler) OnValidateOutputAndCacheAdd(ctx context.Context, nodeID string, i io.InputReader, r io.OutputReader, outputCommitter io.OutputWriter, executionConfig v1alpha1.ExecutionConfig, tr ioutils.SimpleTaskReader, m catalog.Metadata) *TaskNodeHandler_ValidateOutputAndCacheAdd {
	c_call := _m.On("ValidateOutputAndCacheAdd", ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m)
	return &TaskNodeHandler_ValidateOutputAndCacheAdd{Call: c_call}
}

func (_m *TaskNodeHandler) OnValidateOutputAndCacheAddMatch(matchers ...interface{}) *TaskNodeHandler_ValidateOutputAndCacheAdd {
	c_call := _m.On("ValidateOutputAndCacheAdd", matchers...)
	return &TaskNodeHandler_ValidateOutputAndCacheAdd{Call: c_call}
}

// ValidateOutputAndCacheAdd provides a mock function with given fields: ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m
func (_m *TaskNodeHandler) ValidateOutputAndCacheAdd(ctx context.Context, nodeID string, i io.InputReader, r io.OutputReader, outputCommitter io.OutputWriter, executionConfig v1alpha1.ExecutionConfig, tr ioutils.SimpleTaskReader, m catalog.Metadata) (catalog.Status, *io.ExecutionError, error) {
	ret := _m.Called(ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m)

	var r0 catalog.Status
	if rf, ok := ret.Get(0).(func(context.Context, string, io.InputReader, io.OutputReader, io.OutputWriter, v1alpha1.ExecutionConfig, ioutils.SimpleTaskReader, catalog.Metadata) catalog.Status); ok {
		r0 = rf(ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m)
	} else {
		r0 = ret.Get(0).(catalog.Status)
	}

	var r1 *io.ExecutionError
	if rf, ok := ret.Get(1).(func(context.Context, string, io.InputReader, io.OutputReader, io.OutputWriter, v1alpha1.ExecutionConfig, ioutils.SimpleTaskReader, catalog.Metadata) *io.ExecutionError); ok {
		r1 = rf(ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*io.ExecutionError)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, io.InputReader, io.OutputReader, io.OutputWriter, v1alpha1.ExecutionConfig, ioutils.SimpleTaskReader, catalog.Metadata) error); ok {
		r2 = rf(ctx, nodeID, i, r, outputCommitter, executionConfig, tr, m)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
