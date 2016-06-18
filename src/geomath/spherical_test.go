package geomath

import (
    "testing"
    //"math"
)

func TestNewSpherical(t *testing.T) {
    s := Spherical{}
    if s.T != 0 && s.P != 0 {
        t.Fail()
    }
}

func TestNewSpheXYZ(t *testing.T) {
    s := NewSpherical(1, 0, 0)
    if (s != Spherical{1.5707963267948966, 0}) {
        t.Fail()
    }
}

func TestVectorFromSpheCoor(t *testing.T) {
    s := NewSpherical(1, 0, 0)
    v := s.ToVector()
    if (v.X != 1) {
        t.Fail()
    }
}