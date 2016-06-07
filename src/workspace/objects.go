package workspace

import "geomath"

func NewCube(position geomath.Vector, w, h, d float64) (cube *Object) {
    cube = &Object{
        Name: "cube1",
        Origin: position,
        Vertices: []geomath.Vector {
            geomath.Vector{-(w * 0.5), h * 0.5, d * 0.5},
            geomath.Vector{w * 0.5, h * 0.5, d * 0.5},
            geomath.Vector{w * 0.5, -(h * 0.5), d * 0.5},
            geomath.Vector{-(w * 0.5), -(h * 0.5), d * 0.5},
            geomath.Vector{-(w * 0.5), h * 0.5, -(d * 0.5)},
            geomath.Vector{w * 0.5, h * 0.5, -(d * 0.5)},
            geomath.Vector{w * 0.5, -(h * 0.5), -(d * 0.5)},
            geomath.Vector{-(w * 0.5), -(h * 0.5),-(d * 0.5)},
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