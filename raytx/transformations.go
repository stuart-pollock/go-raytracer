package raytx

import (
	"fmt"
	. "raymath"
)

// A Ray type:  it consists of a Point (base location) and a Vector (direction).
type Ray struct {
	P Point
	V Vector
}

// Scale transformation
func (*Matrix) Scale(factor float64) Matrix {
	return nil
}

// Translate transformation
// Translate in x, y, and z by a Vector (a, b, c)
func (*Matrix) Translate(op Vector) Matrix {
	return nil
}

// Rotate transformation
// Rotate around an `axis` (0 = 'x'; 1 = 'y'; 2 = 'z') by `rads` radians
func (*Matrix) Rotate(axis int, rads float64) Matrix {
	return nil
}
