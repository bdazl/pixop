package koch

import (
	"math"

	"github.com/hexhacks/pixop/global"

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

func kochSplit(dst *imdraw.IMDraw, a, b pixel.Vec, it int) {
	if it == 0 {
		return
	}

	bigV := b.Sub(a)                 // Full side
	bigC := a.Add(bigV.Scaled(0.5))  // Center, full side
	smallV := bigV.Scaled(1.0 / 3.0) // New side
	smallN := smallV.Normal()        // Normal
	smallN = smallN.Unit().Scaled(isoHeight(smallV.Len()))

	newPtA := a.Add(smallV)
	newPtB := bigC.Add(smallN)
	newPtC := b.Add(smallV.Scaled(-1))

	it--

	kochSplit(dst, a, newPtA, it)
	dst.Push(newPtA)

	kochSplit(dst, newPtA, newPtB, it)
	dst.Push(newPtB)

	kochSplit(dst, newPtB, newPtC, it)
	dst.Push(newPtC)

	kochSplit(dst, newPtC, b, it)
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

	a := pixel.V(m.X, m.Y+hIsoH)
	b := pixel.V(m.X+hSide, m.Y-hIsoH)
	c := pixel.V(m.X-hSide, m.Y-hIsoH)

	it := 4
	imd.Push(a)
	kochSplit(imd, a, b, it)

	imd.Push(b)
	kochSplit(imd, b, c, it)

	imd.Push(c)
	kochSplit(imd, c, a, it)

	imd.Polygon(1)

	return nil
}

func (k *Koch) Draw(win *pixelgl.Window) {
	k.imd.Draw(win)
}
