package robot

import (
	"github.com/kbinani/robot/key"
	"github.com/kbinani/win"
)

func setKeyboardStatus(nativeKeyCode int, down bool) {
	status := 0
	if !down {
		status = win.KEYEVENTF_KEYUP
	}
	win.Keybd_event(win.BYTE(nativeKeyCode), 0, win.DWORD(status), nil)
}

func nativeKeyCode(code key.Code) int {
	// key.* is equal to VK_* in Windows.
	return int(code)
}
