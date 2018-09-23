package app

import (
	"fmt"
	u "github.com/kbinani/robot/unsafe"
	"github.com/kbinani/win"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

type PID uint32

type App struct {
	pid PID
}

func newApp(pid PID) *App {
	a := new(App)
	a.pid = pid
	return a
}

func (app *App) menu() *Menu {
	w := app.mainWindow()
	w.Activate()
	hMenu := win.GetMenu(w.hWnd)
	fmt.Printf("app.menu(); hMenu=%v w.hWnd=%v\n", hMenu, w.hWnd)
	if hMenu == win.HMENU(0) {
		return nil
	}
	return newMenu(hMenu, w.hWnd, win.HMENU(0))
}

func (app *App) mainWindow() *Window {
	var data tagEnumWindowsCallback
	data.Pid = uint32(app.pid)
	win.EnumWindows(syscall.NewCallback(enumWindowsCallback), (uintptr)(unsafe.Pointer(&data)))
	if data.Handle == win.HWND(0) {
		return nil
	}
	return newWindow(data.Handle)
}

type tagEnumWindowsCallback struct {
	Pid    uint32
	Handle win.HWND
}

func enumWindowsCallback(hWnd win.HWND, lParam uintptr) uintptr {
	var data *tagEnumWindowsCallback = (*tagEnumWindowsCallback)(unsafe.Pointer(lParam))
	var pid uint32
	win.GetWindowThreadProcessId(hWnd, &pid)
	if data.Pid != pid || !isMainWindow(hWnd) {
		return win.TRUE
	}
	data.Handle = hWnd
	return win.FALSE
}

func isMainWindow(hWnd win.HWND) bool {
	return win.GetWindow(hWnd, win.GW_OWNER) == win.HWND(0) && win.IsWindowVisible(hWnd)
}

func (app *App) path() string {
	return pathFromPID(app.pid)
}

func (app *App) name() string {
	path := pathFromPID(app.pid)
	p := unicode16FromString(path)

	size := win.GetFileVersionInfoSize(u.Ptr(p), nil)
	versionInfo := make([]byte, size)
	win.GetFileVersionInfo(u.Ptr(p), 0, size, u.Ptr(versionInfo))

	queryStr := "\\VarFileInfo\\Translation"
	query := unicode16FromString(queryStr)
	var trans *uint32
	var ulen uint32
	win.VerQueryValue(u.Ptr(versionInfo), u.Ptr(query), unsafe.Pointer(&trans), &ulen)

	format := "\\StringFileInfo\\%04x%04x\\FileDescription"
	var queryDescriptionStr string = fmt.Sprintf(format, win.HIWORD(*trans), win.LOWORD(*trans))
	var queryDescription []uint16 = unicode16FromString(queryDescriptionStr)

	var pValue *uint16
	retDesctiption := win.VerQueryValue(u.Ptr(versionInfo), u.Ptr(queryDescription), unsafe.Pointer(&pValue), &ulen)
	if retDesctiption != 0 {
		return win.UTF16PtrToString(pValue)
	}

	// When translation info is not stored, iterate all possible language/codepage values.

	var wCodePageID []uint16 = []uint16{
		0,    // 7-bit ASCII
		932,  // Japan (Shift – JIS X-0208)
		949,  // Korea (Shift – KSC 5601)
		950,  // Taiwan (Big5)
		1200, // Unicode
		1250, // Latin-2 (Eastern European)
		1251, // Cyrillic
		1252, // Multilingual
		1253, // Greek
		1254, // Turkish
		1255, // Hebrew
		1256, // Arabic
	}
	var wLanguageID []uint16 = []uint16{
		0x0401, // Arabic
		0x0402, // Bulgarian
		0x0403, // Catalan
		0x0404, // Traditional Chinese
		0x0405, // Czech
		0x0406, // Danish
		0x0407, // German
		0x0408, // Greek
		0x0409, // U.S. English
		0x040A, // Castilian Spanish
		0x040B, // Finnish
		0x040C, // French
		0x040D, // Hebrew
		0x040E, // Hungarian
		0x040F, // Icelandic
		0x0410, // Italian
		0x0411, // Japanese
		0x0412, // Korean
		0x0413, // Dutch
		0x0414, // Norwegian – Bokmal
		0x0810, // Swiss Italian
		0x0813, // Belgian Dutch
		0x0814, // Norwegian – Nynorsk
		0x0415, // Polish
		0x0416, // Portuguese (Brazil)
		0x0417, // Rhaeto-Romanic
		0x0418, // Romanian
		0x0419, // Russian
		0x041A, // Croato-Serbian (Latin)
		0x041B, // Slovak
		0x041C, // Albanian
		0x041D, // Swedish
		0x041E, // Thai
		0x041F, // Turkish
		0x0420, // Urdu
		0x0421, // Bahasa
		0x0804, // Simplified Chinese
		0x0807, // Swiss German
		0x0809, // U.K. English
		0x080A, // Spanish (Mexico)
		0x080C, // Belgian French
		0x0C0C, // Canadian French
		0x100C, // Swiss French
		0x0816, // Portuguese (Portugal)
		0x081A, // Serbo-Croatian (Cyrillic)
	}
	for _, page := range wCodePageID {
		for _, lang := range wLanguageID {
			queryDescriptionStr = fmt.Sprintf(format, lang, page)
			queryDescription = unicode16FromString(queryDescriptionStr)
			retDesctiption := win.VerQueryValue(u.Ptr(versionInfo), u.Ptr(queryDescription), unsafe.Pointer(&pValue), &ulen)
			if retDesctiption != 0 {
				return win.UTF16PtrToString(pValue)
			}
		}
	}
	return ""
}

func openProcess(pid PID) win.HANDLE {
	return win.OpenProcess(win.PROCESS_QUERY_INFORMATION|win.PROCESS_VM_READ, win.FALSE, uint32(pid))
}

func unicode16FromString(s string) []uint16 {
	r := make([]rune, 0)
	for _, c := range s {
		r = append(r, c)
	}
	b := utf16.Encode(r)
	return append(b, uint16(0))
}

func stringFromUnicode16(r []uint16) string {
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

func pathFromPID(pid PID) string {
	ph := openProcess(pid)
	defer win.CloseHandle(ph)

	var hMod win.HMODULE
	var cbNeeded uint32
	if ret := win.EnumProcessModules(ph, unsafe.Pointer(&hMod), uint32(unsafe.Sizeof(hMod)), &cbNeeded); ret == win.TRUE {
		name := make([]uint16, 1024)
		win.GetModuleFileNameEx(ph, hMod, u.Ptr(name),
			uint32(len(name))*uint32(unsafe.Sizeof(name[0])))
		return stringFromUnicode16(name)
	}
	return ""
}

func ps() []PID {
	buffer := make([]uint32, 512)
	pids := make([]PID, 0)
	for true {
		var bytes uint32
		cb := uint32(len(buffer)) * uint32(unsafe.Sizeof(buffer[0]))
		ret := win.EnumProcesses(u.Ptr(buffer),
			cb,
			&bytes)
		if ret == win.FALSE {
			return pids
		}
		if bytes == cb {
			buffer = make([]uint32, len(buffer)*2)
			continue
		}
		count := int(bytes) / int(unsafe.Sizeof(buffer[0]))
		for i := 0; i < count; i++ {
			pids = append(pids, PID(buffer[i]))
		}
		return pids
	}
	return pids
}
