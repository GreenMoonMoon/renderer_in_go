package rendering

import "geomath"

func ProjectOnCanvas(v geomath.Vector) [2]float64 {
    return [2]float64{v.X / -(v.Z), v.Y / -(v.Z)}
}