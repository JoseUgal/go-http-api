// Code generated by mockery v2.12.3. DO NOT EDIT.
package eventmocks

import (
	context "context"

	event "github.com/JoseUgal/go-http-api/kit/event"
	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0, _a1
func (_m *Bus) Publish(_a0 context.Context, _a1 []event.Event) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []event.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: _a0, _a1
func (_m *Bus) Subscribe(_a0 event.Type, _a1 event.Handler) {
	_m.Called(_a0, _a1)
}

type NewBusT interface {
	mock.TestingT
	Cleanup(func())
}

// NewBus creates a new instance of Bus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBus(t NewBusT) *Bus {
	mock := &Bus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
