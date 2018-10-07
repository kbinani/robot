package robot

/*
#include <Carbon/Carbon.h>

static void releaseCGEvent(CGEventRef o) {
    CFRelease(o);
}

static float get_keyboard_backlight_brightness() {
    io_service_t service_object = IOServiceGetMatchingService(kIOMasterPortDefault,
                                                              IOServiceMatching("AppleLMUController"));
    kern_return_t kr = KERN_FAILURE;
    io_connect_t connect;
    if (service_object) {
        kr = IOServiceOpen(service_object, mach_task_self(), 0, &connect);
    }
    IOObjectRelease(service_object);
    if (kr != KERN_SUCCESS) {
        return -1;
    }
    uint32_t input_count = 1;
    uint32_t output_count = 1;
    uint64_t unknown = 0;
    uint64_t brightness = 0;
    int kGetLEDBrightnessID = 1;
    kr = IOConnectCallMethod(connect,
                             kGetLEDBrightnessID,
                             &unknown,
                             input_count,
                             nil,
                             0,
                             &brightness,
                             &output_count,
                             nil,
                             0);
    IOServiceClose(connect);
    if (kr != KERN_SUCCESS) {
        return -1;
    }
    return brightness / 4091.0;
}
*/
import "C"

import (
	"github.com/kbinani/robot/key"
)

func setKeyboardStatus(nativeKeyCode int, down bool) {
	event := C.CGEventCreateKeyboardEvent(0, (C.CGKeyCode)(nativeKeyCode), C.bool(down))
	if event == 0 {
		return
	}
	defer C.releaseCGEvent(event)
	C.CGEventPost(C.kCGHIDEventTap, event)
}

func isKeyboardDown(code int) bool {
	if code > 0xff {
		flags := C.CGEventSourceFlagsState(C.kCGEventSourceStateHIDSystemState)
		return (flags & (C.CGEventFlags)(code)) == (C.CGEventFlags)(code)
	} else {
		ret := C.CGEventSourceKeyState(C.kCGEventSourceStateHIDSystemState, (C.CGKeyCode)(code))
		return ret == C.bool(true)
	}
}

func nativeKeyCode(code key.Code) int {
	switch code {
	case key.A:
		return C.kVK_ANSI_A
	case key.B:
		return C.kVK_ANSI_B
	case key.C:
		return C.kVK_ANSI_C
	case key.D:
		return C.kVK_ANSI_D
	case key.E:
		return C.kVK_ANSI_E
	case key.F:
		return C.kVK_ANSI_F
	case key.G:
		return C.kVK_ANSI_G
	case key.H:
		return C.kVK_ANSI_H
	case key.I:
		return C.kVK_ANSI_I
	case key.J:
		return C.kVK_ANSI_J
	case key.K:
		return C.kVK_ANSI_K
	case key.L:
		return C.kVK_ANSI_L
	case key.M:
		return C.kVK_ANSI_M
	case key.N:
		return C.kVK_ANSI_N
	case key.O:
		return C.kVK_ANSI_O
	case key.P:
		return C.kVK_ANSI_P
	case key.Q:
		return C.kVK_ANSI_Q
	case key.R:
		return C.kVK_ANSI_R
	case key.S:
		return C.kVK_ANSI_S
	case key.T:
		return C.kVK_ANSI_T
	case key.U:
		return C.kVK_ANSI_U
	case key.V:
		return C.kVK_ANSI_V
	case key.W:
		return C.kVK_ANSI_W
	case key.X:
		return C.kVK_ANSI_X
	case key.Y:
		return C.kVK_ANSI_Y
	case key.Z:
		return C.kVK_ANSI_Z
	case key.Alt:
		return C.kVK_Option
	case key.Ctrl:
		return C.kVK_Control
	case key.RCtrl:
		return C.kVK_RightControl
	case key.Esc:
		return C.kVK_Escape
	case key.Tab:
		return C.kVK_Tab
	case key.Return:
		return C.kVK_Return
	case key.Shift:
		return C.kVK_Shift
	case key.RShift:
		return C.kVK_RightShift
	case key.Capital:
		return C.kVK_CapsLock
	case key.Space:
		return C.kVK_Space
	case key.Prior:
		return C.kVK_PageUp
	case key.Next:
		return C.kVK_PageDown
	case key.End:
		return C.kVK_End
	case key.Home:
		return C.kVK_Home
	case key.Left:
		return C.kVK_LeftArrow
	case key.Up:
		return C.kVK_UpArrow
	case key.Right:
		return C.kVK_RightArrow
	case key.Down:
		return C.kVK_DownArrow
	case key.Delete:
		return C.kVK_Delete
	case key.Help:
		return C.kVK_Help
	case key.Separator:
		return C.kVK_ANSI_Comma
	case key.Subtract:
		return C.kVK_ANSI_Minus
	case key.Decimal:
		return C.kVK_ANSI_Period
	case key.Divide:
		return C.kVK_ANSI_Slash
	case key.F1:
		return C.kVK_F1
	case key.F2:
		return C.kVK_F2
	case key.F3:
		return C.kVK_F3
	case key.F4:
		return C.kVK_F4
	case key.F5:
		return C.kVK_F5
	case key.F6:
		return C.kVK_F6
	case key.F7:
		return C.kVK_F7
	case key.F8:
		return C.kVK_F8
	case key.F9:
		return C.kVK_F9
	case key.F10:
		return C.kVK_F10
	case key.F11:
		return C.kVK_F11
	case key.F12:
		return C.kVK_F12
	case key.VolumeMute:
		return C.kVK_Mute
	case key.VolumeDown:
		return C.kVK_VolumeDown
	case key.VolumeUp:
		return C.kVK_VolumeUp
	case key.OemPlus:
		return C.kVK_ANSI_KeypadPlus
	case key.OemMinus:
		return C.kVK_ANSI_KeypadMinus
	case key.OemPeriod:
		return C.kVK_ANSI_KeypadDecimal
	case key.Command:
		return C.kCGEventFlagMaskCommand
	}
	return -1
}

func backlightBrightness() float32 {
	return (float32)(C.get_keyboard_backlight_brightness())
}
