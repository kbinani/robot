package app

import (
	"github.com/kbinani/robot/ax"
)

type Window struct {
	axWindow *ax.UIElement
}

func newWindow(axWindow *ax.UIElement) *Window {
	w := new(Window)
	w.axWindow = axWindow
	return w
}

func (w *Window) activate() {
	w.axWindow.Perform(ax.RaiseAction)
}

func (w *Window) minimize() {
	children := w.axWindow.Children()
	for _, child := range children {
		role := child.Role()
		subrole := child.Subrole()
		if role == ax.ButtonRole && subrole == ax.MinimizeButtonSubrole {
			child.Perform(ax.PressAction)
			break
		}
	}
}

func (w *Window) isMinimized() bool {
	return w.axWindow.IsMinimized()
}
