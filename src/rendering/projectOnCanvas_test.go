package rendering

import (
    "geomath"
    "testing"
)

var corners []geomath.Vector = []geomath.Vector{
    geomath.Vector{1, -1, -5},
    geomath.Vector{1, -1, -3},
    geomath.Vector{1, 1, -5},
    geomath.Vector{1, 1, -3},
    geomath.Vector{-1, -1, -5},
    geomath.Vector{-1, -1, -3},
    geomath.Vector{-1, 1, -5},
    geomath.Vector{-1, 1, -3},
}

func Test1Point(t *testing.T) {
    pixel := ProjectOnCanvas(corners[0])
    if pixel != [2]float64{0.2, -0.2} {
        t.Error(pixel)
    }
}

func Test2Point(t *testing.T) {
    p1 := ProjectOnCanvas(corners[1])
    p2 := ProjectOnCanvas(corners[2])
    if p1 != [2]float64{(1 / 3), (-1 / 3)} && p2 != [2]float64{0.2, 0.2} {
        t.Error(p1, p2)
    }
}