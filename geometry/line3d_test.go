package geometry_test

import (
	"testing"

	"github.com/gossie/aoc-utils/geometry"
)

func Test3dLineContainsPoint_true(t *testing.T) {
	line := geometry.NewLine3d(geometry.NewPoint3d(1, 2, 3), geometry.NewPoint3d(3, 3, 3))
	point := geometry.NewPoint3d(5, 4, 3)
	if !line.Contains(point) {
		t.Fatalf("point %v should be on line %v", point, line)
	}
}

func Test3dLineContainsPoint_false(t *testing.T) {
	line := geometry.NewLine3d(geometry.NewPoint3d(1, 2, 3), geometry.NewPoint3d(3, 3, 3))
	point := geometry.NewPoint3d(5, 4, 4)
	if !line.Contains(point) {
		t.Fatalf("point %v should be on line %v", point, line)
	}
}

// func Test3dLinesIntersect(t *testing.T) {
// 	line1 := geometry.NewLine3d(geometry.NewPoint3d(1, 2, 3), geometry.NewPoint3d(3, 3, 3))
// 	line2 := geometry.NewLine3d(geometry.NewPoint3d(0, 0, 0), geometry.NewPoint3d(1, 1, 1))
// 	intersection, _ := line1.IntersectsAt(line2)
// 	if intersection != geometry.NewPoint3d(3, 3) {
// 		t.Fatalf("intersection should be at %v", geometry.NewPoint3d(3, 3))
// 	}
// }

// func TestIntersects_noIntersection(t *testing.T) {
// 	line1 := geometry.NewLine3d(geometry.NewPoint3d(1, 2), geometry.NewPoint3d(3, 3))
// 	line2 := geometry.NewLine3d(geometry.NewPoint3d(0, 1), geometry.NewPoint3d(2, 2))
// 	_, err := line1.IntersectsAt(line2)
// 	if err == nil || err.Error() != "no intersection" {
// 		t.Fatalf("lines should not intersect")
// 	}
// }

// func TestIntersects_sameLine(t *testing.T) {
// 	line1 := geometry.NewLine3d(geometry.NewPoint3d(1, 2), geometry.NewPoint3d(3, 3))
// 	line2 := geometry.NewLine3d(geometry.NewPoint3d(5, 4), geometry.NewPoint3d(7, 5))
// 	_, err := line1.IntersectsAt(line2)
// 	if err == nil || err.Error() != "same line" {
// 		t.Fatalf("lines should not intersect")
// 	}
// }
