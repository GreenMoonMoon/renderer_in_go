package workspace

import "testing"

var objA := Object{
	Name:   "objectA",
	Center: geomath.Vector{0, 0, -3},
	Vertices: []geomath.Vector{
		geomath.Vector{-1, 1, 1},
		geomath.Vector{1, 1, 1},
		geomath.Vector{1, -1, 1},
		geomath.Vector{-1, -1, 1},
		geomath.Vector{-1, 1, -1},
		geomath.Vector{1, 1, -1},
		geomath.Vector{1, -1, -1},
		geomath.Vector{-1, -1, -1},
	},
}

func TestEmptyScene(t *testing.T) {
    s := Scene{Name: "Empty"}
    if s.Name != "Empty" {
        t.Fail()
    }
}

func TestScene1EmptyObj(t *testing.T) {
    s := Scene{Name: "SceneA"}
    s.Add(Object{})
    if len(s.Elements) != 1 {
        t.Fail()
    }
}

func TestSceneAddObjA(t *testing.T) {
    s := Scene{Name: "SceneB"}
    s.Add(objA)
}