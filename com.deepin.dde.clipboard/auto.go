package clipboard

import (
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type Clipboard struct {
	clipboard // interface com.deepin.dde.Clipboard
	proxy.Object
}

func NewClipboard(conn *dbus.Conn) *Clipboard {
	obj := new(Clipboard)
	obj.Object.Init_(conn, "com.deepin.dde.Clipboard", "/com/deepin/dde/Clipboard")
	return obj
}

type clipboard struct{}

func (v *clipboard) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*clipboard) GetInterfaceName_() string {
	return "com.deepin.dde.Clipboard"
}

// method Toggle

func (v *clipboard) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *clipboard) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *clipboard) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *clipboard) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Hide

func (v *clipboard) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *clipboard) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}
