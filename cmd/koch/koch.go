package koch

import (
	"github.com/hexhacks/pixop/global"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

type Koch struct {
	side float64 // side length of the largest isosceles triangle
	imd  *imdraw.IMDraw
}

func New() *Koch {
	return &Koch{}
}

// the height of an isosceles triangle.
func isoHeight(side float64) float64 {
	hside := side / 2.0
	return math.Sqrt(side*side - hside*hside)
}

func (k *Koch) Setup() error {
	k.side = global.Bounds.H() / 3.0
	hSide := k.side / 2.0
	isoH := isoHeight(k.side)
	hIsoH := isoH / 2.0

	m := global.Bounds.Center()

	k.imd = imdraw.New(nil)
	imd := k.imd

	imd.Color = colornames.Green
	imd.EndShape = imdraw.RoundEndShape

	imd.Push(pixel.V(m.X, m.Y+hIsoH))
	imd.Push(pixel.V(m.X-hSide, m.Y-hIsoH))
	imd.Push(pixel.V(m.X+hSide, m.Y-hIsoH))

	imd.Polygon(0)

	return nil
}

func (k *Koch) Draw(win *pixelgl.Window) {
	k.imd.Draw(win)
}
