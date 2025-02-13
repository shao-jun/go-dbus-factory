// Code generated by "./generator ./com.deepin.system.systeminfo"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package systeminfo

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type SystemInfo interface {
	systemInfo // interface com.deepin.system.SystemInfo
	proxy.Object
}

type objectSystemInfo struct {
	interfaceSystemInfo // interface com.deepin.system.SystemInfo
	proxy.ImplObject
}

func NewSystemInfo(conn *dbus.Conn) SystemInfo {
	obj := new(objectSystemInfo)
	obj.ImplObject.Init_(conn, "com.deepin.system.SystemInfo", "/com/deepin/system/SystemInfo")
	return obj
}

type systemInfo interface {
	MemorySizeHuman() proxy.PropString
	MemorySize() proxy.PropUint64
	CurrentSpeed() proxy.PropUint64
}

type interfaceSystemInfo struct{}

func (v *interfaceSystemInfo) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSystemInfo) GetInterfaceName_() string {
	return "com.deepin.system.SystemInfo"
}

// property MemorySizeHuman s

func (v *interfaceSystemInfo) MemorySizeHuman() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "MemorySizeHuman",
	}
}

// property MemorySize t

func (v *interfaceSystemInfo) MemorySize() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "MemorySize",
	}
}

// property CurrentSpeed t

func (v *interfaceSystemInfo) CurrentSpeed() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "CurrentSpeed",
	}
}
