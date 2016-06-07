package geomath

import "testing"

/*Normaly I should separate all these test per function. But for the sake of
simplicity and readability I have included all these test inside the same
file*/

func TestVectorValue0(t *testing.T) {
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

func TestVectorValue101(t *testing.T) {
	var v Vector = Vector{1, 0, 1}
	if v.X != 1 {
		t.Fail()
	}
	if v.Y != 0 {
		t.Fail()
	}
	if v.Z != 1 {
		t.Fail()
	}
}

func TestVectorLengthF(t *testing.T) {
	v := Vector{1, 0.2, 1.3}
	if i := Length(v); i != 1.6522711641858308 {
		t.Error(i)
	}
}

func TestVectorNormal(t *testing.T) {
	v := Vector{323, 4, 67}
	if i := Length(Normal(v)); i != 1 {
		t.Error(i)
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

func TestVectorCrossProductB(t *testing.T) {
	v1 := Vector{1, 0, 0}
	v2 := Vector{0, 0, 1}
	v3 := Vector{0, -1, 0}
	if i := CrossProduct(v1, v2); i == v3 {
		return
	} else {
		t.Error(i)
	}
}
