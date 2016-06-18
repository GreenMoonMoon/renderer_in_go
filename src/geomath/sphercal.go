package geomath

//The norm regarding sphercal coordinate is that the up vector is the z axis.
//note: geomath library is based on the scratchapixel lessons, therefor the
//up axis used in cartesian coordinate system is the y axis.

type Sphercal struct {
    T, P float64
}

//Return a new sphercal coordinate from cartesian coordinate.
func NewSphercal(x, y, z float64) Sphercal {
    return Sphercal{
        0,
        0,
    }
}

//return a vector computed from the sphercal coordinate values.
//note: the function assume the sphercal is normalized.
func (s *Sphercal) ToVector() Vector {
    return Vector{
        0,
        0,
        0,
    }
}