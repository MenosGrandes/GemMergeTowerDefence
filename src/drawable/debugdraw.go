package drawable

import (
	"example/hello/src/aabb"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var whitePixel *ebiten.Image

func init() {
	whitePixel = ebiten.NewImage(1, 1)
	whitePixel.Fill(color.White)
}

func DrawSolidArrow(screen *ebiten.Image, x, y, dx, dy float64, length float64, col uint32) {
	l := math.Hypot(dx, dy)
	if l == 0 {
		return
	}
	ndx, ndy := dx/l, dy/l

	// Tip of the arrow
	tipX := x + ndx*length
	tipY := y + ndy*length

	// Arrowhead width
	headWidth := length * 0.4
	angle := math.Pi / 2 // 90° perpendicular

	// Rotate ±90° for base
	leftX := ndx*math.Cos(angle) - ndy*math.Sin(angle)
	leftY := ndx*math.Sin(angle) + ndy*math.Cos(angle)

	rightX := ndx*math.Cos(-angle) - ndy*math.Sin(-angle)
	rightY := ndx*math.Sin(-angle) + ndy*math.Cos(-angle)

	baseLeftX := x + leftX*headWidth*0.5
	baseLeftY := y + leftY*headWidth*0.5

	baseRightX := x + rightX*headWidth*0.5
	baseRightY := y + rightY*headWidth*0.5

	verts := []ebiten.Vertex{
		{DstX: float32(tipX), DstY: float32(tipY), SrcX: 0, SrcY: 0},
		{DstX: float32(baseLeftX), DstY: float32(baseLeftY), SrcX: 0, SrcY: 0},
		{DstX: float32(baseRightX), DstY: float32(baseRightY), SrcX: 0, SrcY: 0},
	}

	r, g, b, a := argbToColor(col)
	for i := range verts {
		verts[i].ColorR = r
		verts[i].ColorG = g
		verts[i].ColorB = b
		verts[i].ColorA = a
	}

	indices := []uint16{0, 1, 2}
	screen.DrawTriangles(verts, indices, whitePixel, nil)
}

func argbToColor(c uint32) (r, g, b, a float32) {
	a = float32((c>>24)&0xff) / 255
	r = float32((c>>16)&0xff) / 255
	g = float32((c>>8)&0xff) / 255
	b = float32(c&0xff) / 255
	return
}

func DrawCircle(dst *ebiten.Image, x, y, r float64, clr color.Color) {
	for angle := 0.0; angle < 2*math.Pi; angle += 0.01 {
		x1 := x + r*math.Cos(angle)
		y1 := y + r*math.Sin(angle)
		dst.Set(int(x1), int(y1), clr)
	}
}
func DrawRectPolygon(dst *ebiten.Image, x, y, w, h float64, clr color.Color) {
	// Define rectangle corners
	vertices := []ebiten.Vertex{
		{DstX: float32(x), DstY: float32(y), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(x + w), DstY: float32(y), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(x), DstY: float32(y + h), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: float32(x + w), DstY: float32(y + h), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	}

	indices := []uint16{
		0, 1, 2,
		2, 1, 3,
	}

	// Create an image filled with the color
	img := ebiten.NewImage(1, 1)
	defer img.Dispose()
	img.Fill(clr)

	opts := &ebiten.DrawTrianglesOptions{}
	dst.DrawTriangles(vertices, indices, img, opts)
}

func DrawCirclePolygon(dst *ebiten.Image, cx, cy, r float64, clr color.Color) {
	segments := 30
	vertices := make([]ebiten.Vertex, segments+1)
	indices := make([]uint16, segments*3)

	// Center vertex
	vertices[0] = ebiten.Vertex{
		DstX:   float32(cx),
		DstY:   float32(cy),
		ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
	}

	// Create circle vertices
	for i := 0; i < segments; i++ {
		angle := 2 * math.Pi * float64(i) / float64(segments)
		x := cx + r*math.Cos(angle)
		y := cy + r*math.Sin(angle)
		vertices[i+1] = ebiten.Vertex{
			DstX:   float32(x),
			DstY:   float32(y),
			ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
		}
	}

	// Create triangle fan indices
	for i := 0; i < segments; i++ {
		indices[i*3] = 0
		indices[i*3+1] = uint16(i + 1)
		if i == segments-1 {
			indices[i*3+2] = 1
		} else {
			indices[i*3+2] = uint16(i + 2)
		}
	}

	img := ebiten.NewImage(1, 1)
	defer img.Dispose()
	img.Fill(clr)

	opts := &ebiten.DrawTrianglesOptions{}
	dst.DrawTriangles(vertices, indices, img, opts)
}
func DrawRect2(dst *ebiten.Image, x, y, w, h float64, clr color.Color) {
	for dx := 0.0; dx < w; dx++ {
		dst.Set(int(x+dx), int(y), clr)
	}
	// Bottom edge
	for dx := 0.0; dx < w; dx++ {
		dst.Set(int(x+dx), int(y+h-1), clr)
	}
	// Left edge
	for dy := 0.0; dy < h; dy++ {
		dst.Set(int(x), int(y+dy), clr)
	}
	// Right edge
	for dy := 0.0; dy < h; dy++ {
		dst.Set(int(x+w-1), int(y+dy), clr)
	}
}

var DebugDrawUtil = struct {
	DrawRect  func(dst *ebiten.Image, x, y, w, h float64, clr color.Color)
	DrawRect2 func(dst *ebiten.Image, x, y, w, h float64, clr color.Color)
	DrawAABB  func(dst *ebiten.Image, aabb aabb.AABB, clr color.Color)

	DrawCircle        func(dst *ebiten.Image, x, y, r float64, clr color.Color)
	DrawRectPolygon   func(dst *ebiten.Image, x, y, w, h float64, clr color.Color)
	DrawCirclePolygon func(dst *ebiten.Image, cx, cy, r float64, clr color.Color)
}{
	DrawRect: func(dst *ebiten.Image, x, y, w, h float64, clr color.Color) {
		for dx := 0.0; dx < w; dx++ {
			for dy := 0.0; dy < h; dy++ {
				dst.Set(int(x+dx), int(y+dy), clr)
			}
		}
	},
	DrawRect2: DrawRect2,
	DrawAABB: func(dst *ebiten.Image, aabb aabb.AABB, clr color.Color) {
		DrawRect2(dst,
			aabb.Center.X-aabb.HalfWidth,
			aabb.Center.Y-aabb.HalfHeight,
			aabb.HalfWidth*2,
			aabb.HalfHeight*2,
			clr)
	},
	DrawCircle:        DrawCircle,
	DrawRectPolygon:   DrawRectPolygon,
	DrawCirclePolygon: DrawCirclePolygon,
}
