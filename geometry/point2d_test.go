package geometry_test

import (
	"math"
	"testing"

	"github.com/gossie/aoc-utils/geometry"
)

func TestAddPoint2d(t *testing.T) {
	p := geometry.CreatePoint2d(7, 17)
	sum := p.Add(geometry.CreatePoint2d(2, -7))
	if sum.X() != 9 || sum.Y() != 10 {
		t.Fatalf("%v != 9 or %v != 10", sum.X(), sum.Y())
	}
}

func TestMultiplyPoint2d(t *testing.T) {
	p := geometry.CreatePoint2d(7, 17)
	product := p.Multiply(4)
	if product.X() != 28 || product.Y() != 68 {
		t.Fatalf("%v != 28 or %v != 10", product.X(), product.Y())
	}
}

func TestAirDistanceToPoint2d_1(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.AirDistanceTo(geometry.CreatePoint2d(4, 8))
	if distance != math.Sqrt(20) {
		t.Fatalf("%v != %v", distance, math.Sqrt(20))
	}
}

func TestAirDistanceToPoint2d_2(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.AirDistanceTo(geometry.CreatePoint2d(0, 8))
	if distance != math.Sqrt(20) {
		t.Fatalf("%v != %v", distance, math.Sqrt(20))
	}
}

func TestAirDistanceToPoint2d_3(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.AirDistanceTo(geometry.CreatePoint2d(0, 0))
	if distance != math.Sqrt(20) {
		t.Fatalf("%v != %v", distance, math.Sqrt(20))
	}
}

func TestAirDistanceToPoint2d_4(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.AirDistanceTo(geometry.CreatePoint2d(4, 0))
	if distance != math.Sqrt(20) {
		t.Fatalf("%v != %v", distance, math.Sqrt(20))
	}
}

func TestManhattenDistanceToPoint2d_1(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.ManhattenDistanceTo(geometry.CreatePoint2d(4, 8))
	if distance != 6 {
		t.Fatalf("%v != %v", distance, 6)
	}
}

func TestManhattenDistanceToPoint2d_2(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.ManhattenDistanceTo(geometry.CreatePoint2d(0, 8))
	if distance != 6 {
		t.Fatalf("%v != %v", distance, 6)
	}
}

func TestManhattenDistanceToPoint2d_3(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.ManhattenDistanceTo(geometry.CreatePoint2d(0, 0))
	if distance != 6 {
		t.Fatalf("%v != %v", distance, 6)
	}
}

func TestManhattenDistanceToPoint2d_4(t *testing.T) {
	p := geometry.CreatePoint2d(2, 4)
	distance := p.ManhattenDistanceTo(geometry.CreatePoint2d(4, 0))
	if distance != 6 {
		t.Fatalf("%v != %v", distance, 6)
	}
}
