package geomath

import (
    "testing"
    "math"
)

func TestRotateVectorA(t *testing.T) {
    v := Vector{1, 0, 0}
    m := RotationMatrix(90, ZAxis)
    r := TransformMatrix33(&v, m)
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
    r := TransformMatrix33(&v, m)
    if (r != Vector{
        -math.Sin(math.Pi / 2),
        math.Cos(math.Pi / 2),
        0,
    }) {
        t.Error(m)
    }
}