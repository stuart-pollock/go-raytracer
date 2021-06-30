package raymath

import (
	"math"
)

// import "fmt"

// Following the course notes for University of Waterloo's CS488
// (https://student.cs.uwaterloo.ca/~cs488/Winter2021/notes.pdf)

// Numerical threshold
// Because we are not working algebraically, we expect rounding to happen.
// We need to have a non-zero threshold which bounds zero, but which effectively means zero
// (any value in [-EPSILON, EPSILON] is treated as 0).
const EPSILON = 0.00000000001

// Point-Vector struct
type IPV interface {
	Add(op IV) *IPV
	MatMult(op IM) *IPV
}

// P defines operations on Points
type IP interface {
	IPV
	Subtract(op IP) IV
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
	Cross(v Vector) IN

	// Calculate a standard dot product of this vector with another Vector,
	// and return the value as a float64.
	Dot(op Vector) float64

	// Return the length (magnitude) of this Vector.
	Len() float64

	// Scale this Vector by a factor. The return value is this Vector.
	Mult(factor float64) IV
}

// Vector represents a direction of a certain magnitude in three-dimensional space
type Vector struct {
	PV
}

type IN interface {
	IV
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
	val := Point{}

	val.XYZ = make([]float64, 4)
	val.XYZ[0] = x
	val.XYZ[1] = y
	val.XYZ[2] = z
	val.XYZ[3] = 1

	return val
}

// Adds a Vector to this PV. Its return value is the PV which
// is `Add`'s receiver.
func (pv *PV) Add(op Vector) *PV {
	for i := 0; i < len(pv.XYZ)-1; i++ {
		pv.XYZ[i] += op.XYZ[i]
	}
	return pv
}

// Subtracts a Point from this Point. The return value is a Vector with
// a manitude equal to the distance between the Points.
func (p *Point) Subtract(op Point) Vector {
	v := GetVector(0, 0, 0)

	for i := 0; i < len(v.XYZ)-1; i++ {
		v.XYZ[i] = p.XYZ[i] - op.XYZ[i]
	}

	return v
}

// Premultiply the current Point or Vector by a transformation matrix
func (pv *PV) MatMult(op Matrix) *PV {
	scratch := GetVector(0, 0, 0)

	for row := 0; row < len(op.ROW); row++ {
		for col := 0; col < len(op.ROW[row]); col++ {
			scratch.XYZ[row] += op.ROW[row][col] * pv.XYZ[col]
		}
	}

	for i := 0; i < len(pv.XYZ)-1; i++ {
		pv.XYZ[i] = scratch.XYZ[i]
	}

	return pv
}

// Return a Vector which is prepopulated with a direction and magnitude
func GetVector(x, y, z float64) Vector {
	vec := Vector{}

	vec.XYZ = make([]float64, 4)
	vec.XYZ[0] = x
	vec.XYZ[1] = y
	vec.XYZ[2] = z
	vec.XYZ[3] = 0

	return vec
}

// Calculate the Cross product between the two Vectors.
func (v *Vector) Cross(op Vector) Normal {
	n := GetNormal(0, 0, 0)

	myLen := len(v.XYZ) - 1
	for i := 0; i < myLen; i++ {
		n.XYZ[i] =
			v.XYZ[(i+1)%myLen]*op.XYZ[(i+2)%myLen] -
				op.XYZ[(i+1)%myLen]*v.XYZ[(i+2)%myLen]
	}

	return n
}

// Calculate the dot product between the two Vectors. In the interest of speed,
// the function does not normalize either of the Vectors.
//
// result := (a, b, c).(d, e, f) == a*d + b*e + c*f
func (v *Vector) Dot(op Vector) float64 {
	var sum float64

	for i := 0; i < len(v.XYZ)-1; i++ {
		sum += v.XYZ[i] * op.XYZ[i]
	}

	return sum
}

// Calculate the length (magnitude) of this Vector.
// result := sqrt((a, b, c).(a, b, c)) == sqrt(a*a + b*b + c*c)
func (v *Vector) Len() float64 {
	return math.Sqrt(v.Dot(*v))
}

// Multiply the current vector by a float64 factor. The return value
// is the magnified Vector.
func (v *Vector) Mult(factor float64) *Vector {
	for i := 0; i < len(v.XYZ)-1; i++ {
		v.XYZ[i] *= factor
	}

	return v
}

// Return a Normal which is prepopulated with a direction and magnitude
func GetNormal(x, y, z float64) Normal {
	norm := Normal{}

	norm.XYZ = make([]float64, 4)
	norm.XYZ[0] = x
	norm.XYZ[1] = y
	norm.XYZ[2] = z
	norm.XYZ[3] = 0

	return norm
}

// Postmultiply the current Normal by a transformation Matrix.
func (n *Normal) MatMult(op Matrix) *Normal {
	scratch := GetNormal(0, 0, 0)

	for col := 0; col < len(op.ROW); col++ {
		for row := 0; row < len(op.ROW[col]); row++ {
			scratch.XYZ[row] += op.ROW[row][col] * n.XYZ[col]
		}
	}

	for i := 0; i < len(n.XYZ)-1; i++ {
		n.XYZ[i] = scratch.XYZ[i]
	}

	return n
}

// Return an identity Matrix
//
// (1 0 0 0)
// (0 1 0 0)
// (0 0 1 0)
// (0 0 0 1)
func GetIdentityMatrix() Matrix {
	mat := Matrix{}

	mat.ROW = make([][]float64, 4)
	for row := 0; row < 4; row++ {
		mat.ROW[row] = make([]float64, 4)
		mat.ROW[row][row] = 1
	}

	return mat
}

// Return a prepopulated Matrix
//
// (a b c p.XYZ[0])
// (d e f p.XYZ[1])
// (g h i p.XYZ[2])
// (0 0 0 1)
func GetMatrix(a, b, c, d, e, f, g, h, i float64, p Point) Matrix {
	mat := Matrix{}

	mat.ROW = make([][]float64, 4)
	for row := 0; row < 4; row++ {
		mat.ROW[row] = make([]float64, 4)
	}

	mat.ROW[0][0] = a
	mat.ROW[0][1] = b
	mat.ROW[0][2] = c
	mat.ROW[0][3] = p.XYZ[0]
	mat.ROW[1][0] = d
	mat.ROW[1][1] = e
	mat.ROW[1][2] = f
	mat.ROW[1][3] = p.XYZ[1]
	mat.ROW[2][0] = g
	mat.ROW[2][1] = h
	mat.ROW[2][2] = i
	mat.ROW[2][3] = p.XYZ[2]
	mat.ROW[3][3] = 1

	return mat
}

// Premultiply the current Matrix byand another Matrix, and return the current Matrix.
func (m *Matrix) MatMult(op Matrix) *Matrix {
	scratch := GetIdentityMatrix()
	// zero-out the main diagonal
	for row := 0; row < len(op.ROW); row++ {
		scratch.ROW[row][row] = 0
	}

	for row := 0; row < len(op.ROW); row++ {
		for col := 0; col < len(op.ROW[row]); col++ {
			for i := 0; i < len(op.ROW[row]); i++ {
				scratch.ROW[row][col] += op.ROW[row][i] * m.ROW[i][col]
			}
		}
	}

	for row := 0; row < len(m.ROW); row++ {
		for col := 0; col < len(m.ROW[row]); col++ {
			m.ROW[row][col] = scratch.ROW[row][col]
		}
	}

	return m
}

// Private implementations
