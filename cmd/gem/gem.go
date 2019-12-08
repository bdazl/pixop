package gem

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	unit = 100.0

	particleCount = 10
)

type Gem struct {
	circle *imdraw.IMDraw // unit circle
	win    *pixelgl.Window
	particles
}

// many particles
type particles struct {
	pos  []pixel.Vec
	vel  []pixel.Vec
	mass []float64
}

func (p *particles) Pos(f func(int, *pixel.Vec)) {
	for i, ps := range p.pos {
		f(i, &ps)
		p.pos[i] = ps
	}
}

func (p *particles) Vel(f func(int, *pixel.Vec)) {
	for i, v := range p.vel {
		f(i, &v)
	}
}

func (p *particles) Mass(f func(int, *float64)) {
	for i, m := range p.mass {
		f(i, &m)
	}
}

func (p *particles) Draw(g *Gem) {
	p.Pos(func(i int, p *pixel.Vec) {
		g.Circ(5, p)
	})
}

func New() *Gem {
	const (
		// estimate screen size
		w = 1024
		h = 800
	)
	circle := imdraw.New(nil)
	circle.Color = colorful.HappyColor()
	circle.Push(pixel.V(0, 0))
	circle.Circle(unit, 0)

	gem := &Gem{
		circle: circle,
		particles: particles{
			pos:  make([]pixel.Vec, particleCount),
			vel:  make([]pixel.Vec, particleCount),
			mass: make([]float64, particleCount),
		},
	}

	parts := &gem.particles

	parts.Pos(func(i int, v *pixel.Vec) {
		base := ((i + 100) * i * i * i)
		v.X = float64(base % w)
		v.Y = float64(base % h)
	})

	return gem
}

func (g *Gem) Setup() error {
	return nil
}

func (g *Gem) Draw(win *pixelgl.Window) {
	/*var (
		bounds = global.Bounds
		center = bounds.Center()
	)*/
	g.win = win

	g.particles.Draw(g)
}

func (g *Gem) Circ(r float64, v *pixel.Vec) {
	scal := r / unit
	mat := pixel.IM
	mat = mat.ScaledXY(pixel.ZV, pixel.V(scal, scal))
	mat = mat.Moved(*v)

	g.win.SetMatrix(mat)
	g.circle.Draw(g.win)
}
