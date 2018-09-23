package app

// #cgo LDFLAGS: -framework ApplicationServices
// #include <libproc.h>
// #include <ApplicationServices/ApplicationServices.h>
import "C"

import (
	"github.com/kbinani/robot/ax"
	"unsafe"
)

type PID C.pid_t

type App struct {
	pid   PID
	axApp *ax.UIElement
}

func newApp(pid PID) *App {
	a := new(App)
	a.pid = pid
	a.axApp = ax.CreateApplication(int(pid))
	return a
}

func (app *App) menu() *Menu {
	return newMenu(app.axApp.MenuBar())
}

func (app *App) path() string {
	return pathFromPID(C.pid_t(app.pid))
}

func (app *App) name() string {
	buffer := make([]C.char, 8192)
	size := C.proc_name(C.int(app.pid), unsafe.Pointer(&buffer[0]), C.uint32_t(len(buffer))*C.uint32_t(unsafe.Sizeof(buffer[0])))
	return C.GoStringN(&buffer[0], size)
}

func (app *App) windows() []*Window {
	ret := []*Window{}
	list := app.axApp.Windows()
	for _, w := range list {
		ret = append(ret, newWindow(w))
	}
	return ret
}

func pathFromPID(pid C.pid_t) string {
	buffer := make([]C.char, 1024)
	size := C.proc_pidpath(C.int(pid), unsafe.Pointer(&buffer[0]), C.uint32_t(len(buffer))*C.uint32_t(unsafe.Sizeof(buffer[0])))
	return C.GoStringN(&buffer[0], size)
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
