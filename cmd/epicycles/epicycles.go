package epicycles

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/hexhacks/pixop/global"

	"image/color"
)

const (
	unit = 100.0
)

type Epicycles struct {
	circle *imdraw.IMDraw // unit circle
	win    *pixelgl.Window
}

func New() *Epicycles {
	circle := imdraw.New(nil)
	circle.Color = colorful.HappyColor()
	circle.Push(pixel.V(0, 0))
	circle.Circle(unit, 3)

	return &Epicycles{circle: circle}
}

func (e *Epicycles) Setup() error {
	return nil
}

func (e *Epicycles) Draw(win *pixelgl.Window) {
	var (
		bounds = global.Bounds
		center = bounds.Center()
	)
	e.win = win

	e.Circ(unit, center)
	e.Circ(unit*2, pixel.V(100, 100))
}

func (e *Epicycles) Circ(r float64, v pixel.Vec) {
	scal := r / unit
	mat := pixel.IM
	mat = mat.ScaledXY(pixel.ZV, pixel.V(scal, scal))
	mat = mat.Moved(v)

	e.win.SetMatrix(mat)
	e.circle.Draw(e.win)
}

func GenerateCurve(
	count int,
	thickness float64,
	closed bool,
	curve func(int) (*pixel.Vec, color.Color)) *imdraw.IMDraw {

	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape

	for i := 0; i < count; i++ {
		PushPoint(i, imd, curve)
	}

	if closed {
		PushPoint(0, imd, curve)
	}

	imd.Line(thickness)
	return imd
}

func PushPoint(i int, imd *imdraw.IMDraw, curve func(int) (*pixel.Vec, color.Color)) {
	vec, col := curve(i)

	imd.Color = col
	imd.Push(*vec)
}
