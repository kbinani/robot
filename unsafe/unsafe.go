package unsafe

import (
	"fmt"
	"reflect"
	u "unsafe"
)

func Ptr(arr interface{}) u.Pointer {
	switch t := arr.(type) {
	case []uint8:
		if len(t) == 0 {
			return u.Pointer(uintptr(0))
		} else {
			return u.Pointer(&t[0])
		}
	case []uint16:
		if len(t) == 0 {
			return u.Pointer(uintptr(0))
		} else {
			return u.Pointer(&t[0])
		}
	case []uint32:
		if len(t) == 0 {
			return u.Pointer(uintptr(0))
		} else {
			return u.Pointer(&t[0])
		}
	}
	panic(fmt.Errorf("unsupported type: %v", reflect.TypeOf(arr)))
}
