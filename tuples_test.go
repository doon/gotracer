package gotracer

import (
	"math"
	"testing"
)

func TestTup(t *testing.T) {

	p := point(4, -4, 3)
	v := vector(4, -4, 3)

	expectedP := tuple(4, -4, 3, 1)
	expectedV := tuple(4, -4, 3, 0)

	if p != expectedP {
		t.Fatalf("Point Doesn't create correct tup got %+v want %+v", p, expectedP)
	}
	if v != expectedV {
		t.Fatalf("Vector Doesn't create expected tup got %+v want %+v", v, expectedV)
	}
}

func TestAddTuple(t *testing.T) {
	t1 := tuple(3, -2, 5, 1)
	t2 := tuple(-2, 3, 1, 0)
	got := addTuple(t1, t2)
	want := tuple(1, 1, 6, 1)
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	type testStruct struct {
		t1       tup
		t2       tup
		expected tup
	}

	testData := []testStruct{
		{point(3, 2, 1), point(5, 6, 7), vector(-2, -4, -6)},
		{point(3, 2, 1), vector(5, 6, 7), point(-2, -4, -6)},
		{vector(3, 2, 1), vector(5, 6, 7), vector(-2, -4, -6)},
	}

	for _, s := range testData {
		v := subtractTuple(s.t1, s.t2)
		if v != s.expected {
			t.Error(
				"For", s.t1,
				"-", s.t2,
				"expected", s.expected,
				"got", v,
			)
		}
	}
}

func TestNegation(t *testing.T) {
	tup1 := tuple(1, -2, 3, -4)
	want := tuple(-1, 2, -3, 4)
	got := negateTuple(tup1)
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestMult(t *testing.T) {
	tup1 := tuple(1, -2, 3, -4)
	want := tuple(3.5, -7, 10.5, -14)
	got := mult(tup1, 3.5)
	if got != want {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestDiv(t *testing.T) {
	tup1 := tuple(1, -2, 3, -4)
	want := tuple(0.5, -1, 1.5, -2)
	got := div(tup1, 2)
	if got != want {
		t.Error(
			"For", tup1,
			"got", got,
			"expected", want,
		)
	}
}

func TestMagnitude(t *testing.T) {
	type testpair struct {
		vector tup
		mag    float64
	}

	tests := []testpair{
		{vector(1, 0, 0), 1},
		{vector(0, 1, 0), 1},
		{vector(0, 0, 1), 1},
		{vector(1, 2, 3), math.Sqrt(14)},
		{vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, pair := range tests {
		v := magnitude(pair.vector)
		if v != pair.mag {
			t.Error(
				"For", pair.vector,
				"expected", pair.mag,
				"got", v,
			)
		}
	}
}

func TestNormalization(t *testing.T) {
	type testpair struct {
		vector tup
		normal tup
	}

	tests := []testpair{
		{vector(4, 0, 0), vector(1, 0, 0)},
		{vector(1, 2, 3), vector(0.26726, 0.53452, 0.80178)},
	}

	for _, pair := range tests {
		v := normalize(pair.vector)
		if !equalTup(v, pair.normal) {
			t.Error(
				"For", pair.vector,
				"expected", pair.normal,
				"got", v,
			)
		}
	}
}

func TestDot(t *testing.T) {
	v1 := vector(1, 2, 3)
	v2 := vector(2, 3, 4)
	want := 20.0
	got := dot(v1, v2)
	if got != want {
		t.Error(
			"For", v1, v2,
			"expected", want,
			"got", got,
		)
	}
}

func TestCross(t *testing.T) {
	v1 := vector(1, 2, 3)
	v2 := vector(2, 3, 4)
	wantV1V2 := vector(-1, 2, -1)
	gotV1V2 := cross(v1, v2)
	wantV2V1 := vector(1, -2, 1)
	gotV2V1 := cross(v2, v1)

	if wantV1V2 != gotV1V2 {
		t.Error(
			"For", v1, v2,
			"expected", wantV1V2,
			"got", gotV1V2,
		)
	}

	if wantV2V1 != gotV2V1 {
		t.Error(
			"For", v2, v1,
			"expected", wantV2V1,
			"got", gotV2V1,
		)
	}
}
