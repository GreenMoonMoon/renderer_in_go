package tests

import (
	"geomath"
	"testing"
)

func TestNewVector(t *testing.T) {
	v := geomath.NewVec3(1, 2, 3)
	if v.Y() != 2 {
		t.Fail()
	}
	if l := v.Length(); l < 3.74165 || l > 3.74167 {
		t.Fail()
	}
	if n := v.Unit(); n.Length() != 1 {
		t.Fail()
	}
}

func TestEqualOppositeVector(t *testing.T) {
	a := geomath.NewVec3(1, 2, 3)
	b := geomath.NewVec3(1, 2, 3)
	c := geomath.Subtract(geomath.Vec3{}, a)
	if !geomath.Equal(a, b) {
		t.Fail()
	}
	if !geomath.Opposite(a, c) {
		t.Fail()
	}
}