package geomath

import "math"

//The norm regarding spherical coordinate is that the up vector is the z axis.
//note: geomath library is based on the scratchapixel lessons, therefor the
//up axis used in cartesian coordinate system is the y axis.
//Note: The radius of a vector in spherical coordinate is the length.

type Spherical struct {
    T, P float64
}

//Return a new spherical coordinate from cartesian coordinate.
/*Lesson: The spherical coordinate can be computed with the help of cosine and 
arctangent functions. Node that the formula should be arccosine of z over radius.*/
func NewSpherical(x, y, z float64) Spherical {
    return Spherical{
        math.Acos(z),
        math.Atan2(y, x),
    }
}

//return a vector computed from the spherical coordinate values.
/*Note that the function assume the spherical is normalized. Also if the spherical
coordinate radius is not normalized, you have to multiply the result of each axis by
the radius.*/
func (s *Spherical) ToVector() Vector {
    return Vector{
        math.Cos(s.P) * math.Sin(s.T),
        math.Sin(s.P) * math.Sin(s.T),
        math.Cos(s.T),
    }
}