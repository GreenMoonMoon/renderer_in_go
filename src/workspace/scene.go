package workspace

import "geomath"

type Scene struct {
    Name string
    Objects []Object
}

type Object struct {
	Name     string
	Center   geomath.Vector
	Vertices []geomath.Vector
	triangle []int
}

func (s *Scene) Add(o Object) {
    s.Objects = append(s.Objects, o)
}