package robot

import (
	"github.com/kbinani/robot/key"
)

// Kbd changes key statuses of keyboaard.
func Kbd(code key.Code, op Op) {
	nativeKeyCode := nativeKeyCode(code)
	if nativeKeyCode <= 0 {
		return
	}
	if op != Up {
		setKeyboardStatus(nativeKeyCode, true)
	}
	if op != Down {
		setKeyboardStatus(nativeKeyCode, false)
	}
}
