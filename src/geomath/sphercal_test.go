package geomath

import (
    "testing"
    //"math"
)

func TestNewSphercal(t *testing.T) {
    s := Sphercal{}
    if s.T != 0 && s.P != 0 {
        t.Fail()
    }
}

func TestNewSpheXYZ(t *testing.T) {
    s := NewSphercal(1, 0, 0)
    if (s != Sphercal{0, 0}) {
        t.Fail()
    }
}

func TestVectorFromSpheCoor(t *testing.T) {
    s := NewSphercal(1, 0, 0)
    v := s.ToVector()
    if (v != Vector{1, 0, 0}) {
        t.Fail()
    }
}