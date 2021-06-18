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

// This enmerates the axes along which some transformations may be applied.
type Axis int

const (
	X Axis = iota
	Y
	Z
)

// Scale transformation
// Scale along `axis` by a factor of`factor`.
func (*Matrix) Scale(axis Axis, factor float64) error {
	return nil
}

// Translate transformation
// Translate in x, y, and z by a Vector (a, b, c)
func (*Matrix) Translate(op Vector) error {
	return nil
}

// Rotate transformation
// Rotate around an `axis` by `rads` radians
func (*Matrix) Rotate(axis Axis, rads float64) error {
	return nil
}
