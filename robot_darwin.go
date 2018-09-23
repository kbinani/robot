package robot

// #include <CoreFoundation/CoreFoundation.h>
import "C"

func cfstr(s string) C.CFStringRef {
	return C.CFStringCreateWithCString(nil, C.CString(s), C.kCFStringEncodingUTF8)
}
