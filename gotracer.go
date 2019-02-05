package gotracer

import "math"

const epsilon float64 = 0.00001

func equal(f1 float64, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}

//Tup The main tuple represeting a point/vector/color
type tup struct {
	x float64
	y float64
	z float64
	w float64
}

func equalTup(t1 tup, t2 tup) bool {
	return equal(t1.x, t2.x) &&
		equal(t1.y, t2.y) &&
		equal(t1.z, t2.z) &&
		equal(t1.w, t2.w)
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

func magnitude(t tup) float64 {
	return math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2))
}

func normalize(t tup) tup {
	mag := magnitude(t)
	return tup{
		t.x / mag,
		t.y / mag,
		t.z / mag,
		t.w / mag,
	}
}

func dot(t1 tup, t2 tup) float64 {
	return t1.x*t2.x + t1.y*t2.y + t1.z*t2.z + t1.w*t2.w
}

func cross(t1 tup, t2 tup) tup {
	return vector(
		t1.y*t2.z-t1.z*t2.y,
		t1.z*t2.x-t1.x*t2.z,
		t1.x*t2.y-t1.y*t2.x,
	)
}
