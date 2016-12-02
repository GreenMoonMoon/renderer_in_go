package geomath

import (
    "testing"
    "math"
)

func TestRotateVectorA(t *testing.T) {
    v := Vector{1, 0, 0}
    m := RotationMatrix(90, ZAxis)
    r := MultVectMatrix33(&v, m)
    if (r != Vector{
        math.Cos(math.Pi / 2),
        math.Sin(math.Pi / 2),
        0,
    }) {
        t.Error(m)
    }
}

func TestRotateVectorB(t *testing.T) {
    v := Vector{0, 1, 0}
    m := RotationMatrix(90, ZAxis)
    r := MultVectMatrix33(&v, m)
    if (r != Vector{
        -math.Sin(math.Pi / 2),
        math.Cos(math.Pi / 2),
        0,
    }) {
        t.Error(m)
    }
}

func TestInvertMatrix(t *testing.T) {
    m := [4][4]float64{
        [4]float64{1, 2, 3, 4},
        [4]float64{5, 6, 7, 8},
        [4]float64{9, 10, 11, 12},
        [4]float64{13, 14, 15, 16},
    }
    //mb := Invert(m)
    t.Error(m)
}