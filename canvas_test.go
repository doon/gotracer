package gotracer

import "testing"

func TestColor(t *testing.T) {
	c := color(-0.5, 0.4, 1.7)
	if !equals(c.R, -0.5) {
		t.Error("Red not correct")
	}
	if !equals(c.B, 1.7) {
		t.Error("Blue not correct")
	}
	if !equals(c.G, 0.4) {
		t.Error("Green not correct")
	}
}

func TestAddingColors(t *testing.T) {
	c1 := color(0.9, 0.6, 0.75)
	c2 := color(0.7, 0.1, 0.25)
	want := color(1.6, 0.7, 1.0)
	got := c1.add(c2)
	if !want.equals(got) {
		t.Error(
			"Expected", want,
			"Got", got,
		)
	}
}
func TestSubtractingColors(t *testing.T) {
	c1 := color(0.9, 0.6, 0.75)
	c2 := color(0.7, 0.1, 0.25)
	want := color(0.2, 0.5, 0.5)
	got := c1.sub(c2)
	if !want.equals(got) {
		t.Error(
			"Expected", want,
			"Got", got,
		)
	}
}
func TestMutiplyingByScalar(t *testing.T) {
	c := color(0.2, 0.3, 0.4)
	want := color(0.4, 0.6, 0.8)
	got := c.scale(2)
	if !want.equals(got) {
		t.Error(
			"Expected", want,
			"Got", got,
		)
	}
}

func TestMultiplyColors(t *testing.T) {
	c1 := color(1, 0.2, 0.4)
	c2 := color(0.9, 1, 0.1)
	want := color(0.9, 0.2, 0.04)
	got := c1.mult(c2)
	if !want.equals(got) {
		t.Error(
			"Expected", want,
			"Got", got,
		)
	}
}
