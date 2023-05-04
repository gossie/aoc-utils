package geometry_test

import (
	"testing"

	"github.com/gossie/aoc-utils/geometry"
)

func TestContainsPoint_true(t *testing.T) {
	line := geometry.NewLine2d(geometry.NewPoint2d(1, 2), geometry.NewPoint2d(3, 3))
	point := geometry.NewPoint2d(5, 4)
	if !line.Contains(point) {
		t.Fatalf("point %v should be on line %v", point, line)
	}
}

func TestContainsPoint_false(t *testing.T) {
	line := geometry.NewLine2d(geometry.NewPoint2d(1, 2), geometry.NewPoint2d(3, 3))
	point := geometry.NewPoint2d(16, 5)
	if line.Contains(point) {
		t.Fatalf("point %v should not be on line %v", point, line)
	}
}

func TestIntersects_hasIntersection(t *testing.T) {
	line1 := geometry.NewLine2d(geometry.NewPoint2d(1, 2), geometry.NewPoint2d(3, 3))
	line2 := geometry.NewLine2d(geometry.NewPoint2d(0, 0), geometry.NewPoint2d(1, 1))
	intersection, _ := line1.IntersectsAt(line2)
	if intersection != geometry.NewPoint2d(3, 3) {
		t.Fatalf("intersection should be at %v", geometry.NewPoint2d(3, 3))
	}
}

func TestIntersects_noIntersection(t *testing.T) {
	line1 := geometry.NewLine2d(geometry.NewPoint2d(1, 2), geometry.NewPoint2d(3, 3))
	line2 := geometry.NewLine2d(geometry.NewPoint2d(0, 1), geometry.NewPoint2d(2, 2))
	_, err := line1.IntersectsAt(line2)
	if err == nil {
		t.Fatalf("lines should not intersect")
	}
}
