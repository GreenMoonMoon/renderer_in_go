package workspace

import (
	"geomath"
	"testing"
)

func TestCube(t *testing.T) {
	if c := NewCube(geomath.Vector{0, 0, 0}, 4, 5, 6); c.Name != "cube1" {
		t.Fail()
	}

}

/*func TestGrid(t *testing.T) {
    if g := NewGrid(geomath.Vector{3, 0, 2}, 10, 10, 10); g.Name != "grid1" {
        t.Fail()
    }
}*/
