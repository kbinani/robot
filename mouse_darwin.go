package robot

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>

static void releaseCGEvent(CGEventRef o) {
	CFRelease(o);
}
*/
import "C"

import (
	"errors"
	"image"
)

func mmv(pos image.Point) error {
	p := C.CGPointMake(C.CGFloat(pos.X), C.CGFloat(pos.Y))
	move := C.CGEventCreateMouseEvent(
		nil, C.kCGEventMouseMoved,
		p,
		C.kCGMouseButtonLeft)
	if move == nil {
		return errors.New("cannot create mouse event")
	}
	defer C.releaseCGEvent(move)

	C.CGEventPost(C.kCGHIDEventTap, move)
	return nil
}

func mpos() (image.Point, error) {
	event := C.CGEventCreate(nil)
	if event == nil {
		return image.Pt(0, 0), errors.New("cannot create CGEvent")
	}
	defer C.releaseCGEvent(event)

	loc := C.CGEventGetLocation(event)
	return image.Pt(int(loc.x), int(loc.y)), nil
}

func btn(btn Button, operation Op, pos image.Point) {
	var mouseButton C.CGMouseButton
	mouseButton = C.kCGMouseButtonLeft
	switch btn {
	case Left:
		mouseButton = C.kCGMouseButtonLeft
	case Right:
		mouseButton = C.kCGMouseButtonRight
	case Middle:
		mouseButton = C.kCGMouseButtonCenter
	}
	p := C.CGPointMake(C.CGFloat(pos.X), C.CGFloat(pos.Y))
	if operation != Down {
		var eventType C.CGEventType
		if btn == Right {
			eventType = C.kCGEventRightMouseDown
		} else {
			eventType = C.kCGEventLeftMouseDown
		}
		down := C.CGEventCreateMouseEvent(nil, eventType, p, mouseButton)
		defer C.releaseCGEvent(down)
		C.CGEventPost(C.kCGHIDEventTap, down)
	}
	if operation != Up {
		var eventType C.CGEventType
		if btn == Right {
			eventType = C.kCGEventRightMouseUp
		} else {
			eventType = C.kCGEventLeftMouseUp
		}
		up := C.CGEventCreateMouseEvent(nil, eventType, p, mouseButton)
		defer C.releaseCGEvent(up)
		C.CGEventPost(C.kCGHIDEventTap, up)
	}
}
