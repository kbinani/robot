package main

import (
	"fmt"
	"github.com/kbinani/robot"
	"github.com/kbinani/robot/app"
	"github.com/kbinani/robot/key"
	"image"
	"math"
	"time"
)

func main() {
	itermApps := app.Find("^iTerm*")
	for _, a := range itermApps {
		fmt.Printf("Name: \"%s\", Path: \"%s\", PID: %d\n", a.Name(), a.Path(), a.PID())
	}
	notepadApps := app.Find("^notepad*")
	for _, a := range notepadApps {
		fmt.Printf("Name: \"%s\", Path: \"%s\", PID: %d\n", a.Name(), a.Path(), a.PID())
	}

	robot.Pw(robot.MonitorOff)
	time.Sleep(10 * time.Second)
	robot.Pw(robot.MonitorOn)
	time.Sleep(10 * time.Second)

	// Get current mouse position
	pos, _ := robot.Mpos()
	fmt.Printf("%v\n", pos)

	// Move mouse, with circular trajectory.
	div := 90
	ox := 500
	oy := 500
	radius := 200.0
	for i := 0; i <= div; i++ {
		radian := float64(i) / float64(div) * 2.0 * math.Pi
		sin, cos := math.Sincos(radian)
		x := ox + int(radius*sin)
		y := oy + int(radius*cos)
		robot.Mmv(image.Pt(x, y))
		time.Sleep(10 * time.Millisecond)
	}

	// Click (500, 500)
	robot.Btn(robot.Left, robot.Click, image.Pt(500, 500))

	// Type keyboard "Hello"
	robot.Kbd(key.Shift, robot.Down)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.H, robot.Down)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.H, robot.Up)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.Shift, robot.Up)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.E, robot.Click)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.L, robot.Click)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.L, robot.Click)
	time.Sleep(50 * time.Millisecond)
	robot.Kbd(key.O, robot.Click)
	time.Sleep(50 * time.Millisecond)
}
