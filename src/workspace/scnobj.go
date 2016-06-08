//Scene objects
package workspace

import "geomath"

type Scene struct {
	Name    string
	Objects []Object
}

type Object struct {
	Name     string
	Origin   geomath.Vector
	Vertices []geomath.Vector
	Triangles []int
}

func (s *Scene) Add(o Object) {
	s.Objects = append(s.Objects, o)
}
