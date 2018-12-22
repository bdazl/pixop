package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	startTime = time.Now()
	tid       = 0.0
)

func lissajous(bounds pixel.Rect) *imdraw.IMDraw {
	const (
		cycles = 5
		res    = 0.1
		ext    = 2.0
	)

	hw := (bounds.Max.X - bounds.Min.X) / 2
	hh := (bounds.Max.Y - bounds.Min.Y) / 2
	imd := imdraw.New(nil)
	imd.Color = colornames.Green
	imd.EndShape = imdraw.RoundEndShape

	freq := 4.0
	phase := tid

	var prev *pixel.Vec
	for t := 0.0; t < cycles*math.Pi*2.0; t += res {
		x := math.Sin(freq*t + phase)
		y := math.Sin(freq*t + phase + ext)

		sx := bounds.Min.X + hw + (hw * x)
		sy := bounds.Min.Y + hh + (hh * y)

		if prev == nil {
			prev = new(pixel.Vec)
			*prev = pixel.V(sx, sy)
		} else {
			curr := pixel.V(sx, sy)
			imd.Push(*prev, curr)
			*prev = curr
		}
	}

	imd.Line(30)
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

	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		tid = time.Now().Sub(startTime).Seconds()
		liss := lissajous(win.Bounds())
		liss.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
