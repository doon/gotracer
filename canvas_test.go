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

func TestCreateCanvas(t *testing.T) {
	black := color(0, 0, 0)
	c := canvas(10, 20, black)

	if len(c) != 20 {
		t.Error("Expected y == 20 got", len(c))
	}
	if len(c[0]) != 10 {
		t.Error("Expected X == 10 got", len(c[0]))
	}

	for y := range c {
		for x := range c[y] {
			if !c[y][x].equals(black) {
				t.Error("Expected", x, ",", y, "to be black, got", c[y][x])
			}
		}
	}
}

func TestWritePixelToCanvas(t *testing.T) {
	black := color(0, 0, 0)
	red := color(1, 0, 0)
	c := canvas(10, 20, black)
	c.write(2, 3, red)
	if !c.at(2, 3).equals(red) {
		t.Error("expected pixel at 2,3 to be", red, "got", c.at(2, 3))
	}
	if !c.at(1, 3).equals(black) {
		t.Error("expected pixel at 1,3 to be", black, "got", c.at(1, 3))
	}
}
