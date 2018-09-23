package robot

import (
	"image"
)

// Mmv moves mouse cursor to specified position.
func Mmv(pos image.Point) error {
	return mmv(pos)
}

// Mpos returns the position of mouse cursor.
func Mpos() (image.Point, error) {
	return mpos()
}

// Btn operates mouse buttons.
func Btn(button Button, operation Op, pos image.Point) {
	btn(button, operation, pos)
}
