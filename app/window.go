package app

func (w *Window) Activate() {
	w.activate()
}

func (w *Window) Minimize() {
	w.minimize()
}

func (w *Window) IsMinimized() bool {
	return w.isMinimized()
}
