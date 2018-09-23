package robot

import (
	"github.com/kbinani/win"
	"image"
)

func mmv(pos image.Point) error {
	win.SetCursorPos(int32(pos.X), int32(pos.Y))
	return nil
}

func mpos() (image.Point, error) {
	var pos win.POINT
	win.GetCursorPos(&pos)
	return image.Pt(int(pos.X), int(pos.Y)), nil
}

func btn(button Button, op Op, pos image.Point) {
	Mmv(pos)

	if op != Up {
		var down win.DWORD
		if button == Left {
			down = win.MOUSEEVENTF_LEFTDOWN
		} else if button == Right {
			down = win.MOUSEEVENTF_RIGHTDOWN
		} else if button == Middle {
			down = win.MOUSEEVENTF_MIDDLEDOWN
		}
		win.Mouse_event(down, 0, 0, 0, nil)
	}
	if op != Down {
		var up win.DWORD
		if button == Left {
			up = win.MOUSEEVENTF_LEFTUP
		} else if button == Right {
			up = win.MOUSEEVENTF_RIGHTUP
		} else if button == Middle {
			up = win.MOUSEEVENTF_MIDDLEUP
		}
		win.Mouse_event(up, 0, 0, 0, nil)
	}
}
