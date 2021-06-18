package raymath

// Following the course notes for University of Waterloo's CS488
// (https://student.cs.uwaterloo.ca/~cs488/Winter2021/notes.pdf)

// Numerical threshold
// Because we are not working algebraically, we expect rounding to happen.
// We need to have a non-zero threshold which bounds zero, but which effectively means zero
// (any value in [-EPSILON, EPSILON] is treated as 0).
const EPSILON = 0.00000000001

// Point-Vector struct
type IPV interface {
	Add(op Vector) *PV
	MatMult(op Matrix) *PV
}

// P defines operations on Points
type IP interface {
	IPV
	Subtract(op Point) Vector
	Set(x, y, z int) *PV
}

// PointVec represents a supertructure for POints, Vectorsa, and surface Normals.
type PV struct {
	XYZ []float64
}

// Point represents a location in three-dimensional space
type Point struct {
	PV
}

// V defines operations on Vectors
type IV interface {
	IPV

	// Calculate the cross product of this Vector with another Vector,
	// and return the result as a Normal.
	Cross(v Vector) Normal

	// Calculate a standard dot product of this vector with another Vector,
	// and return the value as a float64.
	Dot(op Vector) float64

	// Return the length (magnitude) of this Vector.
	Len() float64

	// Scale this Vector by a factor. The return value is this Vector.
	Mult(factor float64) Vector
}

// Vector represents a direction of a certain magnitude in three-dimensional space
type Vector struct {
	PV
}

// Normal represents a direction perpendicular to a surface at some
// point on the surface
type Normal struct {
	Vector
}

// Premltiply a Matrix by another Matrix
type IM interface {
	MatMult(op Matrix) *Matrix
}

// Matrix represents a compsition of transformations which may be
// applied to either Points or Vectors
type Matrix struct {
	// A row of columns
	ROW [][]float64
}

// Return a Point which is prepopulated with co-ordinates
func GetPoint(x, y, z float64) Point {
	return Point{}
}

// Adds a Vector to this PV. Its return value is the PV which
// is `Add`'s receiver.
func (pv *PV) Add(op Vector) *PV {
	return nil
}

// Subtracts a Point from this Point. The return value is a Vector with
// a manitude equal to the distance between the Points.
func (*Point) Subtract(op Point) Vector {
	return Vector{}
}

// Premultiply the current Point or Vector by a transformation matrix
func (pv *PV) MatMult(op Matrix) *PV {
	return nil
}

// Return a Vector which is prepopulated with a direction and magnitude
func GetVector(x, y, z float64) Vector {
	return Vector{}
}

// Calculate the Cross product between the two Vectors.
func (*Vector) Cross(op Vector) Normal {
	return Normal{}
}

// Calculate the dot product between the two Vectors. In the interest of speed,
// the function does not normalize either of the Vectors.
//
// result := (a, b, c).(d, e, f) == a*d + b*e + c*f
func (*Vector) Dot(op Vector) float64 {
	return 0
}

// Calculate the length (magnitude) of this Vector.
// result := sqrt((a, b, c).(a, b, c)) == sqrt(a*a + b*b + c*c)
func (*Vector) Len() float64 {
	return 0
}

// Multiply the current vector by a float64 factor. The return value
// is the magnified Vector.
func (v *Vector) Mult(factor float64) *Vector {
	return nil
}

// Return a Normal which is prepopulated with a direction and magnitude
func GetNormal(x, y, z float64) Normal {
	return Normal{}
}

// Postmultiply the current Normal by a transformation Matrix.
func (n *Normal) MatMult(op Matrix) *Normal {
	return nil
}

// Return an identity Matrix
//
// (1 0 0 0)
// (0 1 0 0)
// (0 0 1 0)
// (0 0 0 1)
func GetIdentityMatrix() Matrix {
	return Matrix{}
}

// Return a prepopulated Matrix
//
// (a b c p.XYZ[0])
// (d e f p.XYZ[1])
// (g h i p.XYZ[2])
// (0 0 0 1)
func GetMatrix(a, b, c, d, e, f, g, h, i float64, p Point) Matrix {
	return Matrix{}
}

// Premultiply the current Matrix byand another Matrix, and return the current Matrix.
func (m *Matrix) MatMult(op Matrix) *Matrix {
	return nil
}

// Private implementations
