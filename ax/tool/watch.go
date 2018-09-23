package main

// #cgo LDFLAGS: -framework ApplicationServices
// #include <ApplicationServices/ApplicationServices.h>
// #include <libproc.h>
// // https://github.com/NUIKit/GraphicsServices/blob/master/GraphicsServices/GSCoreFoundationBridge.h
// typedef struct __CFRuntimeClass {	// Version 0 struct
// 	CFIndex version;
// 	const char *className;
// 	void (*init)(CFTypeRef cf);
// 	CFTypeRef (*copy)(CFAllocatorRef allocator, CFTypeRef cf);
// #if MAC_OS_X_VERSION_10_2 <= MAC_OS_X_VERSION_MAX_ALLOWED
// 	void (*finalize)(CFTypeRef cf);
// #else
// 	void (*dealloc)(CFTypeRef cf);
// #endif
// 	Boolean (*equal)(CFTypeRef cf1, CFTypeRef cf2);
// 	CFHashCode (*hash)(CFTypeRef cf);
// 	CFStringRef (*copyFormattingDesc)(CFTypeRef cf, CFDictionaryRef formatOptions);	// str with retain
// 	CFStringRef (*copyDebugDesc)(CFTypeRef cf);	// str with retain
// #if MAC_OS_X_VERSION_10_5 <= MAC_OS_X_VERSION_MAX_ALLOWED
// #define CF_RECLAIM_AVAILABLE 1
// 	void (*reclaim)(CFTypeRef cf);
// #endif
// } CFRuntimeClass;
// extern const CFRuntimeClass * _CFRuntimeGetClassWithTypeID(CFTypeID typeID);
// char const* GetClassName(CFTypeID id)
// {
//     CFRuntimeClass const* cls = _CFRuntimeGetClassWithTypeID(id);
//     return cls->className;
// }
import "C"

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"unicode/utf16"
	"unsafe"
)

func main() {
	f, err := os.Open("ax_darwin.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewScanner(bufio.NewReader(f))
	// reg := regexp.MustCompile(`^\s*([A-Za-z0-9]*)Attribute\s*=\s*"(.*)"$`)
	reg := regexp.MustCompile(`^\s*([A-Za-z0-9]*)Attribute\s*=\s*"(.*)".*$`)

	attrs := map[string]string{}

	for r.Scan() {
		line := r.Text()
		result := reg.FindSubmatch([]byte(line))
		if len(result) == 0 {
			continue
		}
		name := fmt.Sprintf("%sAttribute", string(result[1]))
		value := string(result[2])
		attrs[name] = value
	}

	fmt.Printf("attrs=%v\n", attrs)
	result := map[string]string{}
	for true {
		for _, pid := range ps() {
			app := C.AXUIElementCreateApplication(pid)
			defer C.CFRelease((C.CFTypeRef)(app))
			traverse(app, attrs, &result)
		}
		runtime.GC()
	}
}

func traverse(ref C.AXUIElementRef, attrs map[string]string, result *map[string]string) {
	// names := copyAttributeNames(ref)
	for name, value := range attrs {
		if value == "AXChildren" {
			continue
		}
		_, ok := (*result)[name]
		if ok {
			continue
		}
		var v C.CFTypeRef
		a := cfstr(value)
		defer C.CFRelease((C.CFTypeRef)(a))
		C.AXUIElementCopyAttributeValue(ref, a, &v)
		if v == nil {
			continue
		}
		defer C.CFRelease((C.CFTypeRef)(v))

		desc := typename(v)
		if len(desc) > 0 {
			(*result)[name] = desc
			fmt.Printf("%s = \"%s\" // %s\n", name, value, desc)
		}
	}

	// valueDescriptionAttr := cfstr("AXValueDescription")
	// defer C.CFRelease((C.CFTypeRef)(valueDescriptionAttr))
	// var valueDescription C.CFStringRef
	// C.AXUIElementCopyAttributeValue(ref, valueDescriptionAttr, (*C.CFTypeRef)(unsafe.Pointer(&valueDescription)))
	// if valueDescription != nil {
	// 	description := stringFromCFString(valueDescription)
	// 	if len(description) > 0 {
	// 		fmt.Printf("valueDescription=%s\n", stringFromCFString(valueDescription))
	// 	}
	// 	C.CFRelease(C.CFTypeRef(valueDescription))
	// }

	var children C.CFArrayRef
	childrenAttr := cfstr("AXChildren")
	defer C.CFRelease((C.CFTypeRef)(childrenAttr))
	C.AXUIElementCopyAttributeValue(ref, childrenAttr, (*C.CFTypeRef)(unsafe.Pointer(&children)))
	if children == nil {
		return
	}
	defer C.CFRelease((C.CFTypeRef)(children))
	num := int(C.CFArrayGetCount(children))
	for i := 0; i < num; i++ {
		child := (C.AXUIElementRef)(C.CFArrayGetValueAtIndex(children, C.CFIndex(i)))
		traverse(child, attrs, result)
	}
}

func typename(v C.CFTypeRef) string {
	typeID := C.CFGetTypeID(v)
	descPtr := C.GetClassName(typeID)
	desc := C.GoString(descPtr)

	if desc == "CFArray" {
		num := C.CFArrayGetCount(v)
		if num > 0 {
			o := (C.CFTypeRef)(C.CFArrayGetValueAtIndex(v, C.CFIndex(0)))
			subDesc := typename(o)
			return fmt.Sprintf("CFArray<%s>", subDesc)
		} else {
			return ""
		}
	} else if desc == "CFNumber" {
		t := C.CFNumberGetType((C.CFNumberRef)(v))
		m := map[C.CFNumberType]string{
			1:  "SInt8",
			2:  "SInt16",
			3:  "SInt32",
			4:  "SInt64",
			5:  "Float32",
			6:  "Float64",
			7:  "Char",
			8:  "Short",
			9:  "Int",
			10: "Long",
			11: "LongLong",
			12: "Float",
			13: "Double",
			14: "CFIndex",
			15: "NSInteger",
			16: "CGFloat",
		}
		return fmt.Sprintf("CFNumber<%s>", m[t])
	} else if desc == "AXValue" {
		t := C.AXValueGetType((C.AXValueRef)(v))
		m := map[C.AXValueType]string{
			1: "CGPoint",
			2: "CGSize",
			3: "CGRect",
			4: "CFRange",
			5: "AXError",
			// 0: "Illegal",
		}
		if t > 0 {
			return fmt.Sprintf("AXValue<%s>", m[t])
		} else {
			return ""
		}
	}

	return desc
}

func copyAttributeNames(ref C.AXUIElementRef) []string {
	names := []string{}
	var arr C.CFArrayRef
	C.AXUIElementCopyAttributeNames(ref, &arr)
	if arr == nil {
		return names
	}
	defer C.CFRelease((C.CFTypeRef)(arr))
	num := int(C.CFArrayGetCount(arr))
	for i := 0; i < num; i++ {
		item := (C.CFStringRef)(C.CFArrayGetValueAtIndex(arr, C.CFIndex(i)))
		names = append(names, stringFromCFString(item))
	}
	return names
}

func ps() []C.pid_t {
	num := C.proc_listallpids(nil, 0)
	pids := make([]C.int, num*2)
	num = C.proc_listallpids(unsafe.Pointer(&pids[0]), C.int(len(pids))*C.int(unsafe.Sizeof(pids[0])))
	result := make([]C.pid_t, 0)
	for i := 0; i < int(num); i++ {
		result = append(result, C.pid_t(pids[i]))
	}
	return result
}

func stringFromCFString(s C.CFStringRef) string {
	ptr := C.CFStringGetCStringPtr(s, C.kCFStringEncodingUTF8)
	if ptr != nil {
		return C.GoString(ptr)
	}
	length := uint32(C.CFStringGetLength(s))
	uniPtr := C.CFStringGetCharactersPtr(s)
	if uniPtr == nil || length == 0 {
		return ""
	}
	return stringFromUnicode16Ptr((*uint16)(uniPtr), length)
}

func stringFromUnicode16Ptr(p *uint16, length uint32) string {
	r := []uint16{}
	ptr := uintptr(unsafe.Pointer(p))
	for i := uint32(0); i < length; i++ {
		c := *(*uint16)(unsafe.Pointer(ptr))
		r = append(r, c)
		if c == 0 {
			break
		}
		ptr = ptr + unsafe.Sizeof(c)
	}
	r = append(r, uint16(0))
	decoded := utf16.Decode(r)
	n := 0
	for i, r := range decoded {
		if r == rune(0) {
			n = i
			break
		}
	}
	return string(decoded[:n])
}

func cfstr(s string) C.CFStringRef {
	var ptr *C.char = C.CString(s)
	ret := C.CFStringCreateWithCString(nil, ptr, C.kCFStringEncodingUTF8)
	C.free(unsafe.Pointer(ptr))
	return ret
}
