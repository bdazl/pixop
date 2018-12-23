package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	twoPi = math.Pi * 2.0
)

var (
	startTime = time.Now()
	tid       = 0.0
	bounds    pixel.Rect
)

func lissajous() *imdraw.IMDraw {
	const (
		cycles = 1
		period = cycles * twoPi
		ext    = 2.0
	)

	res := period / (200.0*sinPos(tid*twoPi*0.1) + 10)
	hw := (bounds.Max.X - bounds.Min.X) / 2
	hh := (bounds.Max.Y - bounds.Min.Y) / 2

	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.EndShape = imdraw.RoundEndShape

	freq := 9.0
	phase := tid

	for t := 0.0; t <= period; t += res {
		x := math.Sin(t + phase)
		y := math.Sin(freq*t + phase + ext)

		sx := bounds.Min.X + hw + (hw * x)
		sy := bounds.Min.Y + hh + (hh * y)

		imd.Push(pixel.V(sx, sy))
	}

	imd.Line(3)
	return imd
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

	bounds = win.Bounds()

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		tid = time.Now().Sub(startTime).Seconds()
		liss := lissajous()
		liss.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func sinPos(x float64) float64 {
	return math.Sin(x)*0.5 + 0.5
}
