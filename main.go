package main

import (
	"math"
	"time"

	"github.com/hexhacks/pixop/cmd/lissajous"
	"github.com/hexhacks/pixop/global"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	twoPi = math.Pi * 2.0
)

var (
	drawFunc func(*pixelgl.Window)
)

func main() {
	drawFunc = lissajous.Draw

	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Jacob <3 Ulrika",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	global.Bounds = win.Bounds()

	for !win.Closed() {
		calcTime()

		win.Clear(colornames.Skyblue)

		drawFunc(win)

		win.Update()
	}
}

func calcTime() {
	start := global.StartTime
	global.Time = time.Now().Sub(start).Seconds()
}
