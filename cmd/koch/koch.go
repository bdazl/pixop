package koch

import (
	"math"

	"github.com/hexhacks/pixop/global"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

const (
	iterations = 4
)

type Koch struct {
	polygons []*imdraw.IMDraw // first polygon is the largest
}

func New() *Koch {
	return &Koch{make([]*imdraw.IMDraw, 0)}
}

type Edge struct {
	a, b pixel.Vec
}

type Triangle struct {
	a, b, c pixel.Vec
}

func (t Triangle) Edges() []Edge {
	return []Edge{
		{t.a, t.b},
		{t.b, t.c},
		{t.c, t.a},
	}
}

// Only the outermost edges
func (t Triangle) KochEdges() []Edge {
	return []Edge{
		{t.a, t.b},
		{t.b, t.c},
	}
}

func (t Triangle) Points() []pixel.Vec {
	return []pixel.Vec{t.a, t.b, t.c}
}

func (t Triangle) Polygon() *imdraw.IMDraw {
	imd := imdraw.New(nil)

	imd.Color = colornames.Green
	imd.EndShape = imdraw.RoundEndShape

	imd.Push(t.a, t.b, t.c)
	imd.Polygon(0)
	return imd
}

// the height of an isosceles triangle.
func isoHeight(side float64) float64 {
	hside := side / 2.0
	return math.Sqrt(side*side - hside*hside)
}

// Generate a koch sub-polygon from a line between a and b
func KochTriangle(a, b pixel.Vec) Triangle {

	bigV := b.Sub(a)                 // Full side
	bigC := a.Add(bigV.Scaled(0.5))  // Center, full side
	smallV := bigV.Scaled(1.0 / 3.0) // New side
	smallN := smallV.Normal()        // Normal
	smallN = smallN.Unit().Scaled(isoHeight(smallV.Len()))

	return Triangle{
		a.Add(smallV),
		bigC.Add(smallN),
		b.Add(smallV.Scaled(-1)),
	}
}

func (k *Koch) Setup() error {
	// The main triangle properties
	side := global.Bounds.H() / 3.0
	hSide := side / 2.0
	isoH := isoHeight(side)
	hIsoH := isoH / 2.0

	m := global.Bounds.Center()

	tri := Triangle{
		pixel.V(m.X, m.Y+hIsoH),
		pixel.V(m.X+hSide, m.Y-hIsoH),
		pixel.V(m.X-hSide, m.Y-hIsoH),
	}

	k.polygons = append(k.polygons, tri.Polygon())

	// Initialize edges to iterate
	it := iterations
	set := tri.Edges()

	for it > 0 {
		next := []Edge{}
		for len(set) > 0 {
			// Pop first element
			var edge Edge
			edge, set = set[0], set[1:]

			koch := KochTriangle(edge.a, edge.b)

			// Append the two smaller lines between new tri and
			// these edge points
			next = append(next,
				Edge{edge.a, koch.a},
				Edge{koch.c, edge.b},
			)

			// Append the outer triangle edges for this koch triangle
			next = append(next, koch.KochEdges()...)

			// TODO colors
			k.polygons = append(k.polygons, koch.Polygon())
		}

		it--
		set = next
	}

	return nil
}

func (k *Koch) Draw(win *pixelgl.Window) {
	for _, p := range k.polygons {
		p.Draw(win)
	}
}
