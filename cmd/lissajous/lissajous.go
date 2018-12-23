package lissajous

import (
	"math"

	"github.com/hexhacks/pixop/global"
	pmath "github.com/hexhacks/pixop/math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

func Draw(win *pixelgl.Window) {
	fig := lissajous()

	fig.Draw(win)
}

func lissajous() *imdraw.IMDraw {
	const (
		twoPi  = pmath.TwoPi
		cycles = 1
		period = cycles * twoPi
		ext    = 2.0
	)

	var (
		tid    = global.Time
		bounds = global.Bounds
	)

	res := period / (200.0*pmath.SinPos(tid*twoPi*0.1) + 10)
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
