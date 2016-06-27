package geomath

import "math"

//Basic variables for axis
var XAxis, YAxis, ZAxis = Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}
var Zero = Vector{0, 0, 0}

//The structure represente a 3 dimension cartesian coordinate. For simplicity,
//it is called Vector, a sphercal coordinate, on the other hand, is called 
//simply Sphercal.
type Vector struct {
	X, Y, Z float64
}

//Simple distinction between the same structure to allow a more readable code.
//todo: change to an interface to allow reusing the same function on structure
type Point Vector
type Normal Vector

//return the length of a vector
func (v *Vector) Length() float64 {
    //return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
    return math.Sqrt(Dot(*v, *v))
}

//Normalize the vector
//todo: add a function that does not affect the vector and return a new vector.
func (v *Vector) Normalized() Vector {
    l := 1 / v.Length()
    return Vector{
        v.X * l,
        v.Y * l,
        v.Z * l,
    }
}

//Add a vector to vector v.
func Add(a, b Vector) Vector {
    return Vector{
    a.X + b.X,
    a.Y + b.Y,
    a.Z + b.Z,
    }
}

//Subtract a vector from vector v.
func subtract(a, b Vector) Vector{
    return Vector {
        a.X - b.X,
        a.Y - b.Y,
        a.Z - b.Z,
    }
}

//Scalar multiplication of a vector.
func (v *Vector) Scale(scalar float64) {
    v.X *= scalar
    v.Y *= scalar
    v.Z *= scalar
}

//Return the theta of the vector in a spherical coordinate system.
//Note: We assume that the vector is normalized.
func (v *Vector) Theta() float64 {
    //The vector should be normalized. We don't normalize it in this
    //function because the current function will affect the vector itself.
    return math.Acos(v.Z)
}

//Retrurn the phi of vector v in a spherical coordinate system.
func (v *Vector) Phi() float64 {
    p := math.Atan2(v.Y, v.X)
    if p < 0 {
        return p + 2 * math.Pi
    } else {
        return p
    }
}

//Retrun the dot product of two vector.
func Dot(a, b Vector) float64 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
}

//Return the cross product of two vector.
func Cross(a, b Vector) Vector {
	return Vector{
		a.Y*b.Z - b.Y*a.Z,
		a.Z*b.X - b.Z*a.X,
		a.X*b.Y - b.X*a.Y,
	}
}

//Box Product (Scalar triple product) of three vectors
//The box product is also the determinant of a 3 by 3 matrix composed
//of vectors. 
func Box(a, b, c Vector) float64 {
    return Dot(a, Cross(b, c))
}

func Equal(a, b Vector) bool {
    if a.X == b.X && a.Y == b.Y && a.Z == b.Z {
        return true
    } else {
        return false
    }
}

func Opposite(a, b Vector) bool {
    if a.X == -b.X && a.Y == -b.Y && a.Z == -b.Z {
        return true
    } else {
        return false
    }
}