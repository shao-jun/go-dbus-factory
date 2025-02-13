// Code generated by "./generator ./org.freedesktop.miracle.wfd"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package wfd

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-dbus-factory/object_manager"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockWfd struct {
	object_manager.MockInterfaceObjectManager // interface org.freedesktop.DBus.ObjectManager
	MockInterfaceWfd                          // interface org.freedesktop.miracle.wfd
	proxy.MockObject
}

type MockInterfaceWfd struct {
	mock.Mock
}

// method Shutdown

func (v *MockInterfaceWfd) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWfd) Shutdown(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

type MockSink struct {
	MockInterfaceSink // interface org.freedesktop.miracle.wfd.Sink
	proxy.MockObject
}

type MockInterfaceSink struct {
	mock.Mock
}

// method StartSession

func (v *MockInterfaceSink) GoStartSession(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) *dbus.Call {
	mockArgs := v.Called(flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSink) StartSession(flags dbus.Flags, arg0 string, arg1 string, arg2 uint32, arg3 uint32, arg4 uint32, arg5 uint32, arg6 string) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Session o

func (v *MockInterfaceSink) Session() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Peer o

func (v *MockInterfaceSink) Peer() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockSession struct {
	MockInterfaceSession // interface org.freedesktop.miracle.wfd.Session
	proxy.MockObject
}

type MockInterfaceSession struct {
	mock.Mock
}

// method Resume

func (v *MockInterfaceSession) GoResume(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSession) Resume(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Pause

func (v *MockInterfaceSession) GoPause(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSession) Pause(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Teardown

func (v *MockInterfaceSession) GoTeardown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSession) Teardown(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// property Sink o

func (v *MockInterfaceSession) Sink() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Url s

func (v *MockInterfaceSession) Url() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *MockInterfaceSession) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
