package robot

import (
	u "github.com/kbinani/robot/unsafe"
	"github.com/lxn/win"
	"unsafe"
)

func pw(op PwOp) {
	switch op {
	case MonitorOn:
		// SendInput and SendMessage are both necessary to ensure turning on monitors.
		inputs := make([]win.MOUSE_INPUT, 2)
		inputs[0].Mi.Dx = 0
		inputs[0].Mi.DwFlags = win.MOUSEEVENTF_MOVE
		inputs[1].Mi.Dx = 0
		inputs[1].Mi.DwFlags = win.MOUSEEVENTF_MOVE
		win.SendInput(uint32(len(inputs)), u.Ptr(inputs), int32(unsafe.Sizeof(inputs[0])))

		displayPowerOn := -1
		win.SendMessage(win.HWND_BROADCAST, win.WM_SYSCOMMAND, win.SC_MONITORPOWER, uintptr(displayPowerOn))
	case MonitorOff:
		displayPowerOff := 2
		win.SendMessage(win.HWND_BROADCAST, win.WM_SYSCOMMAND, win.SC_MONITORPOWER, uintptr(displayPowerOff))
	}
}
