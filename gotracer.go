package gotracer

import (
	"math"
)

const epsilon float64 = 0.00001

func equals(f1 float64, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}

//Tup The main tuple represeting a point/vector/color
type tup struct {
	x, y, z, w float64
}

func (t tup) equals(t2 tup) bool {
	return equals(t.x, t2.x) &&
		equals(t.y, t2.y) &&
		equals(t.z, t2.z) &&
		equals(t.w, t2.w)
}

func tuple(x float64, y float64, z float64, w float64) tup {
	return tup{x, y, z, w}
}
func point(x float64, y float64, z float64) tup {
	return tup{x, y, z, 1.0}
}

func vector(x float64, y float64, z float64) tup {
	return tup{x, y, z, 0.0}
}

func addTuple(t1 tup, t2 tup) tup {
	return tup{
		t1.x + t2.x,
		t1.y + t2.y,
		t1.z + t2.z,
		t1.w + t2.w,
	}
}
func subtractTuple(t1 tup, t2 tup) tup {
	return tup{
		t1.x - t2.x,
		t1.y - t2.y,
		t1.z - t2.z,
		t1.w - t2.w,
	}
}

func negateTuple(t tup) tup {
	return tup{-t.x, -t.y, -t.z, -t.w}
}

func mult(t tup, f float64) tup {
	return tup{
		t.x * f,
		t.y * f,
		t.z * f,
		t.w * f,
	}
}

func div(t tup, f float64) tup {
	return tup{
		t.x / f,
		t.y / f,
		t.z / f,
		t.w / f,
	}
}

func (t tup) magnitude() float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2))
}

func (t tup) normalize() tup {
	mag := t.magnitude()
	return tup{
		t.x / mag,
		t.y / mag,
		t.z / mag,
		t.w / mag,
	}
}

func (t tup) dot(t2 tup) float64 {
	return t.x*t2.x + t.y*t2.y + t.z*t2.z + t.w*t2.w
}

func (t tup) cross(t2 tup) tup {
	return vector(
		t.y*t2.z-t.z*t2.y,
		t.z*t2.x-t.x*t2.z,
		t.x*t2.y-t.y*t2.x,
	)
}

type rgb struct {
	R, G, B float64
}

func color(red float64, green float64, blue float64) rgb {
	return rgb{red, green, blue}
}

func (c rgb) add(c2 rgb) rgb {
	return color(
		c.R+c2.R,
		c.G+c2.G,
		c.B+c2.B,
	)
}

func (c rgb) sub(c2 rgb) rgb {
	return color(
		c.R-c2.R,
		c.G-c2.G,
		c.B-c2.B,
	)
}

func (c rgb) scale(f float64) rgb {
	return color(
		c.R*f,
		c.G*f,
		c.B*f,
	)
}

func (c rgb) equals(c2 rgb) bool {
	return equals(c.R, c2.R) &&
		equals(c.G, c2.G) &&
		equals(c.B, c2.B)
}

func (c rgb) mult(c2 rgb) rgb {
	return color(
		c.R*c2.R,
		c.G*c2.G,
		c.B*c2.B,
	)
}

type Canvas [][]rgb

func canvas(XSize int, YSize int, defaultColor rgb) Canvas {
	// Allocate the top-level slice
	c := make([][]rgb, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]rgb, XSize*YSize)
	//loop over all pixels and set to default color
	for i := range pixels {
		pixels[i] = defaultColor
	}
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range c {
		c[i], pixels = pixels[:XSize], pixels[XSize:]
	}
	return c
}

func (c Canvas) write(x int, y int, color rgb) {
	c[y][x] = color
}

func (c Canvas) at(x int, y int) rgb {
	return c[y][x]
}
