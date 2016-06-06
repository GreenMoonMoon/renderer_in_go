package workspace

import "geomath"

type Scene struct {
    Name string
    Elements []interface{}
}

type Object struct {
	Name     string
	Center   geomath.Vector
	Vertices []geomath.Vector
	triangle []int
}

func (s *Scene) Add(i interface{}) {
    s.Elements = append(s.Elements, i)
}