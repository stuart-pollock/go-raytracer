package rayprimitive

import "fmt"

// This interface describes the operations which may be applied
// toa Node in the Scene's object hierarchy.
type N interface {
	Transform(tx M) N
}

// A Node in the object hierarchy
// Each Node is located at a Point in space, and has its own transformation Matrix
// Each Node also has a list of child Nodes, and has a name.
type Node struct {
	Name  string
	Loc   Point
	Tx    Matrix
	Child []Node
}

// A primitive in the object hierarchy
// Each primitive is a Node, and which is a concretely visible visual object.
type primitive interface {
	N
	Intersect(tx M) primitive
}

// A Sphere, which has a location (of its centre) and a radius.
type Sphere struct {
	Node
	Radius float64
}

// A Polyhedron, specified as an array of Points and an array
// of faces defined by index references to these Points.
//
// Each face contains an array of (integer) indices into the
// array of Points, and each pair of adjacent Points defines a
// a sequence of edges defining the polygonal face.  The last Point
// of each face defines an edge with the first Point of the sequence.
//
// Note:  each polygon is assumed to be convex.
type Polyhedron struct {
	Pts   []Points
	Faces [][]int
}

// A Box, which is a special case of Polyhedron.
type Box struct {
	Polyhedron
}

// Return a Sphere of radius R centred at (0, 0, 0).
func GetSphere(name string, R float64) Sphere {
	return nil
}

// Return a Polyhedron with Points and faces defined by those Points.
func GetPolyhedron(name string, pts []Point, faces [][]int) Polyhedron {
	return nil
}

// Return a Box with opposing corners at (0, 0, 0) and (1, 1, 1).
func GetBox(name string) Box {
	return nil
}

// Private implementations
