package osd

import (
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type OSD struct {
	osd // interface com.deepin.dde.osd
	proxy.Object
}

func NewOSD(conn *dbus.Conn) *OSD {
	obj := new(OSD)
	obj.Object.Init_(conn, "com.deepin.dde.osd", "/")
	return obj
}

type osd struct{}

func (v *osd) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*osd) GetInterfaceName_() string {
	return "com.deepin.dde.osd"
}

// method ShowOSD

func (v *osd) GoShowOSD(flags dbus.Flags, ch chan *dbus.Call, osd string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowOSD", flags, ch, osd)
}

func (v *osd) ShowOSD(flags dbus.Flags, osd string) error {
	return (<-v.GoShowOSD(flags, make(chan *dbus.Call, 1), osd).Done).Err
}
