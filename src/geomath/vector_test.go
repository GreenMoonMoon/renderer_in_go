package geomath

import "testing"

/*Normaly I should separate all these test per function. But for the sake of
simplicity and readability I have included all these test inside the same
file*/

func TestVectorValueA(t *testing.T) {
	var v Vector
	if v.X != 0 {
		t.Fail()
	}
	if v.Y != 0 {
		t.Fail()
	}
	if v.Z != 0 {
		t.Fail()
	}
}

func TestVectorLength(t *testing.T) {
	v := Vector{1, 0.2, 1.3}
	if i := v.Length(); i != 1.6522711641858308 {
		t.Fail()
	}
}

func TestMultByScalar(t *testing.T) {
    v := Vector{0, 0.5, 0.5}
    v.Scale(2)
    if (v != Vector{0, 1, 1}) {
        t.Fail()
    }
}

func TestVectorNormalizing(t *testing.T) {
	v := Vector{323, 4, 67}
	v.Normalize()
	if i := v.Length(); i != 1 {
		t.Fail()
	}
}

func TestVectorDotProduct(t *testing.T) {
	v1 := Vector{0, 1, 0}
	v2 := Vector{2, 1, 0}
	if i := DotProduct(v1, v2); i != 1 {
		t.Error(i)
	}
}

func TestVectorCrossProductA(t *testing.T) {
	v1 := Vector{3, 1, 7}
	v2 := Vector{2, 1, 0}
	v3 := Vector{-7, 14, 1}
	if i := CrossProduct(v1, v2); i == v3 {
		return
	} else {
		t.Error(i)
	}
}