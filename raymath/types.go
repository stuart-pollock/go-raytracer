package raymath

import "fmt"

// Following the course notes for University of Waterloo's CS488
// (https://student.cs.uwaterloo.ca/~cs488/Winter2021/notes.pdf)

// Point-Vector struct
type PV interface {
	Add(op V) PV
	PreMult(op M) PV
}

// P defines operations on Points
type P interface {
	PV
	Subtract(op P) V
}

// Point represents a location in three-dimensional space
type Point struct {
	XYZ []int
}

// V defines operations on Vectors
type V interface {
	PV
	Cross(v V) N
	Dot(op V) float64
	Len() float64
	Mult(factor float64) V
}

// Vector represents a direction of a certain magnitude in three-dimensional space
type Vector struct {
	XYZ []int
}

// N defines operations on Normal 'Vectors'
type N interface {
	V
	PostMult(op M) N
}

// Normal represents a direction perpendicular to a surface at some
// point on the surface
type Normal struct {
	Vector
}

// M defines operations on Matrix structs
type M interface {
	PreMult(op M) M
}

// Matrix represents a compsition of transformations which may be
// applied to either Points or Vectors
type Matrix struct {
	T [][]int
}

// Return to the caller a Point at the origin (0, 0, 0)
func GetPoint() Point {
	return nil
}

// Return a Point which is prepopulated with co-ordinates
func GetPoint(x, y, z float64) Point {
	return nil
}

// Adds a Vector to this PV. Its return value is the PV which
// is `Add`'s receiver.
func (*PV) Add(op Vector) PV {
	return nil
}

// Subtracts a Point from this Point. The return value is a Vector with
// a manitude equal to the distance between the Points.
func (Point) Subtract(op Point) Vector {
	return nil
}

// Premultiply the current point by a transformation matrix
func (*Point) PreMult(op Matrix) Point {
	return nil
}

// Return a zero-Vector (0, 0, 0) to the caller
func GetVector() Vector {
	return nil
}

// Return a Vector which is prepopulated with a direction and magnitude
func GetVector(x, y, z float64) Vector {
	return nil
}

// Calculate the Cross product between the two Vectors.
func (Vector) Cross(op Vector) Normal {
	return nil
}

// Calculate the dot product between the two Vectors. In the interest of speed,
// the function does not normalize either of the Vectors.
func (Vector) Dot(op Vector) float64 {
	return 0
}

// Calculate the length (magnitude) of this Vector.
func (Vector) Len() float64 {
	return 0
}

// Multiply the current vector by a float64 factor. The return value
// is the magnified Vector.
func (*Vector) Mult(factor float64) Vector {
	return nil
}

// Premultiply the current point by a transformation matrix
func (*Vector) PreMult(op Matrix) Vector {
	return nil
}

// Postmultiply the current vector by a transformation Matrix.
func (*Vector) PostMult(op Matrix) Vector {
	return nil
}

// Return an identity Matrix to the caller
//
// (1 0 0 0)
// (0 1 0 0)
// (0 0 1 0)
// (0 0 0 1)
func GetMatrix() Matrix {
	return nil
}

// Premultiply the current Matrix byand another Matrix, and return the current Matrix.
func (*Matrix) PreMult(op Matrix) Matrix {
	return nil
}

// Private implementations
