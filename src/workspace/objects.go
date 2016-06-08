//Objects prefabs
package workspace

import "geomath"

func NewCube(position geomath.Vector, w, h, d float64) (cube Object) {
	cube = Object{
		Name:   "cube1",
		Origin: position,
		Vertices: []geomath.Vector{
			geomath.Vector{-(w * 0.5), h * 0.5, d * 0.5},
			geomath.Vector{w * 0.5, h * 0.5, d * 0.5},
			geomath.Vector{w * 0.5, -(h * 0.5), d * 0.5},
			geomath.Vector{-(w * 0.5), -(h * 0.5), d * 0.5},
			geomath.Vector{-(w * 0.5), h * 0.5, -(d * 0.5)},
			geomath.Vector{w * 0.5, h * 0.5, -(d * 0.5)},
			geomath.Vector{w * 0.5, -(h * 0.5), -(d * 0.5)},
			geomath.Vector{-(w * 0.5), -(h * 0.5), -(d * 0.5)},
		},
		Triangles: []int{
			0, 1, 3, 1, 2, 3,
			3, 2, 7, 2, 6, 7,
			7, 6, 4, 6, 5, 4,
			4, 5, 0, 5, 1, 0,
			1, 5, 2, 5, 6, 2,
			4, 0, 3, 0, 7, 3,
		},
	}
	return
}

/*func NewGrid(position geomath.Vector, w, d float64, subdivision int)
(grid *Object) {
    s := subdivision * subdivision
    vertices := make([]geomath.Vector, s)
    for i := 0; i < s; i++ {
        for j := 0; j < subdivision; j++ {

        }
    }

    grid = &Object{
        Name: "grid1"
        Origin: position,
        Vertices: vertices,
    }

    return
}*/
