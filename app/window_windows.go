package app

import (
	"github.com/kbinani/win"
	"unsafe"
	// "fmt"
)

type Window struct {
	hWnd win.HWND
}

func newWindow(hWnd win.HWND) *Window {
	w := new(Window)
	w.hWnd = hWnd
	return w
}

func (w *Window) activate() {
	if w.isMinimized() {
		win.ShowWindow(w.hWnd, win.SW_RESTORE)
	}
	win.SetForegroundWindow(w.hWnd)
}

func (w *Window) isMinimized() bool {
	p := w.placement()
	minimized := p.ShowCmd == win.SW_SHOWMINIMIZED
	return minimized
}

func (w *Window) placement() win.WINDOWPLACEMENT {
	var p win.WINDOWPLACEMENT
	p.Length = uint32(unsafe.Sizeof(p))
	win.GetWindowPlacement(w.hWnd, &p)
	return p
}
