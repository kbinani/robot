package app

import (
	"path/filepath"
	"regexp"
)

// Find searches existing (running) app by its executable file path.
func Find(exePath string) []*App {
	r, _ := regexp.Compile(exePath)
	pids := ps()
	apps := make([]*App, 0)

	for _, pid := range pids {
		path := pathFromPID(pid)
		name := filepath.Base(path)
		if r == nil {
			if name == exePath {
				app := newApp(PID(pid))
				apps = append(apps, app)
			}
		} else {
			if r.MatchString(name) {
				app := newApp(PID(pid))
				apps = append(apps, app)
			}
		}
	}
	return apps
}

// Path returns absolute file path of the process's executable file.
func (app *App) Path() string {
	return app.path()
}

// Name returns the name or short description of the app.
func (app *App) Name() string {
	return app.name()
}

// PID returns the process id of the app.
func (app *App) PID() PID {
	return app.pid
}

// Menu returns the top-level menu of the app.
func (app *App) Menu() *Menu {
	return app.menu()
}

// Windows returns list of Window of the app.
func (app *App) Windows() []*Window {
	return app.windows()
}
