package workspace

import (
	"geomath"
	"testing"
)

var objA Object = Object{
	Name:   "objectA",
	Origin: geomath.Vector{0, 0, -3},
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

func TestScene1AddObj(t *testing.T) {
	s := Scene{Name: "SceneA"}
	s.Add(objA)
	if len(s.Objects) != 1 {
		t.Fail()
	}
}

func TestSaveSceneObjA(t *testing.T) {
	s := Scene{Name: "SceneA"}
	s.Add(objA)
	err := s.SaveAs("../../temp/")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadSceneObjA(t *testing.T) {
	s, err := Load("../../temp/SceneA.scn")
	if err != nil {
		t.Error(err)
	}
	if s.Objects[0].Name != "objectA" {
		t.Error(s.Objects[0])
	}
}
