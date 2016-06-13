package indev

import "testing"

//{0.7071067811865475 0.7071067811865475 0}

func TestVectorLength(t *testing.T) {
    v := Vector{0, 0, 1}
    l := v.Length()
    if l != 1 {
        t.Fail()
    }
}

func TestResizeVector(t *testing.T) {
    v := Vector{0, 0.5, 0.5}
    v.ResizeVector(2)
    if (v != Vector{0, 1, 1}) {
        t.Fail()
    }
}

func TestNormalizeVector(t *testing.T) {
    v := Vector{0.5, 0.5, 0}
    v.Normalize()
    if(v == Vector{0.5, 0.5, 0}) {
        t.Fail()
    }
}

func TestRotateVectorA(t *testing.T) {
    v := Vector{1, 0, 0}
    v.RotateVector(90, Vector{0, 0, 1})
    if (v != Vector{0, 1, 0}) {
        t.Error(v)
    }
}

/*func TestrotateVector(t *testing.T) {
    v := Vector{0.7071067811865475, 0.7071067811865475, 0}
    v.RotateVector(45, Vector{0, 0, 1})
    t.Error(v)
}*/

//test matrix transform