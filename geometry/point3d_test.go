package geometry_test

import (
	"math"
	"testing"

	"github.com/gossie/aoc-utils/geometry"
)

func TestAddPoint3d(t *testing.T) {
	p := geometry.NewPoint3d(7, 17, -3)
	sum := p.Add(geometry.NewPoint3d(2, -7, 1))
	if sum.X() != 9 || sum.Y() != 10 || sum.Z() != -2 {
		t.Fatalf("%v != 9 or %v != 10 or %v != -2", sum.X(), sum.Y(), sum.Z())
	}
}

func TestSubtractPoint3d(t *testing.T) {
	p := geometry.NewPoint3d(7, 17, -3)
	sum := p.Subtract(geometry.NewPoint3d(2, -7, 1))
	if sum.X() != 5 || sum.Y() != 24 || sum.Z() != -4 {
		t.Fatalf("%v != 5 or %v != 24 or %v != -4", sum.X(), sum.Y(), sum.Z())
	}
}

func TestMultiplyPoint3d(t *testing.T) {
	p := geometry.NewPoint3d(7, 17, -4)
	product := p.Multiply(-2)
	if product.X() != -14 || product.Y() != -34 || product.Z() != 8 {
		t.Fatalf("%v != -14 or %v != -34 or %v != 8", product.X(), product.Y(), product.Z())
	}
}

func TestAirDistanceToPoint3d_1(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(4, 8, 6))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_2(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(0, 8, 6))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_3(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(4, 0, 6))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_4(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(4, 8, 0))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_5(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(0, 0, 6))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_6(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(0, 8, 0))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_7(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(4, 0, 0))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestAirDistanceToPoint3d_8(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.AirDistanceTo(geometry.NewPoint3d(0, 0, 0))
	if distance != math.Sqrt(29) {
		t.Fatalf("%v != %v", distance, math.Sqrt(29))
	}
}

func TestManhattenDistanceToPoint3d_1(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(4, 8, 6))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_2(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(0, 8, 6))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_3(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(4, 0, 6))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_4(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(4, 8, 0))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_5(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(0, 0, 6))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_6(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(0, 8, 0))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_7(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(4, 0, 0))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}

func TestManhattenDistanceToPoint3d_8(t *testing.T) {
	p := geometry.NewPoint3d(2, 4, 3)
	distance := p.ManhattenDistanceTo(geometry.NewPoint3d(0, 0, 0))
	if distance != 9 {
		t.Fatalf("%v != %v", distance, 9)
	}
}
