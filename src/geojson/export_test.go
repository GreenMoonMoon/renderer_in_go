package geojson

import (
    "testing"
    "geomath"
)

type object struct {
    vertices []geomath.Vector
    triangles []int
}

var square object = object{
    vertices: []geomath.Vector{
        geomath.Vector{1, 1, 2},
        geomath.Vector{1, -1, 2},
        geomath.Vector{-1, -1, 2},
        geomath.Vector{-1, 1, 2},
    },
    triangles: []int{0, 1, 2, 2, 3, 0},
}


func TestExportNull(t *testing.T) {
    a, _ := Export(nil)
    if a !=  ""{
        t.Fail()
    }
}