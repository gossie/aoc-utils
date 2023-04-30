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
}

func TestIntersects_noIntersection(t *testing.T) {

}
