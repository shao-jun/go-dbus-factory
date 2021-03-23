package systeminfo

import (
	"unsafe"

	"github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type SystemInfo struct {
	systemInfo // interface com.deepin.system.SystemInfo
	proxy.Object
}

func NewSystemInfo(conn *dbus.Conn) *SystemInfo {
	obj := new(SystemInfo)
	obj.Object.Init_(conn, "com.deepin.system.SystemInfo", "/com/deepin/system/SystemInfo")
	return obj
}

type systemInfo struct{}

func (v *systemInfo) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*systemInfo) GetInterfaceName_() string {
	return "com.deepin.system.SystemInfo"
}

// property MemorySizeHuman s

func (v *systemInfo) MemorySizeHuman() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MemorySizeHuman",
	}
}

// property MemorySize t

func (v *systemInfo) MemorySize() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "MemorySize",
	}
}

// property CurrentSpeed t

func (v *systemInfo) CurrentSpeed() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "CurrentSpeed",
	}
}
