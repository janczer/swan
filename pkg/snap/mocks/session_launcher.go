package mocks

import (
	"github.com/intelsdi-x/swan/pkg/executor"
	"github.com/intelsdi-x/swan/pkg/experiment/phase"
	"github.com/intelsdi-x/swan/pkg/snap"
	"github.com/stretchr/testify/mock"
)

// SessionLauncher is an autogenerated mock type for the SessionLauncher type
type SessionLauncher struct {
	mock.Mock
}

// LaunchSession provides a mock function with given fields: _a0, _a1
func (_m *SessionLauncher) LaunchSession(_a0 executor.TaskInfo, _a1 phase.Session) (snap.SessionHandle, error) {
	ret := _m.Called(_a0, _a1)

	var r0 snap.SessionHandle
	if rf, ok := ret.Get(0).(func(executor.TaskInfo, phase.Session) snap.SessionHandle); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(snap.SessionHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(executor.TaskInfo, phase.Session) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}