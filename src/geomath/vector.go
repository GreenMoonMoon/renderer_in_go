package geomath

import "math"

//Basic variables for axis
var XAxis, YAxis, ZAxis = Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}

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
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

//Scale a vector by a scalar
func (v *Vector) Scale(scalar float64) {
    v.X *= scalar
    v.Y *= scalar
    v.Z *= scalar
}

//Add a vector to vector v.
func (v *Vector) Add(vector Vector) {
    v.X += vector.X
    v.Y += vector.Y
    v.Z += vector.Z
}

//Subtract a vector from vector v.
func (v *Vector) subtract(vector Vector) {
    v.X -= vector.X
    v.Y -= vector.Y
    v.Z -= vector.Z
}

//Normalize the vector
//todo: add a function that does not affect the vector and return a new vector.
func (v *Vector) Normalize() {
    l := 1 / v.Length()
    v.X = v.X * l
    v.Y = v.Y * l
    v.Z = v.Z * l
}

//Return the theta of the vector in a spherical coordinate system.
//Note: We assume that the vector is normalized.
func (v *Vector) SphericalTheta() float64 {
    //The vector should be normalized. We don't normalize it in this
    //function because the current function will affect the vector itself.
    return math.Acos(v.Z)
}

//Retrurn the phi of vector v in a spherical coordinate system.
func (v *Vector) SphericalPhi() float64 {
    p := math.Atan2(v.Y, v.X)
    if p < 0 {
        return p + 2 * math.Pi
    } else {
        return p
    }
}

//Retrun the dot product of two vector.
func DotProduct(a, b Vector) float64 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
}

//Return the cross product of two vector.
func CrossProduct(a, b Vector) Vector {
	return Vector{
		a.Y*b.Z - b.Y*a.Z,
		a.Z*b.X - b.Z*a.X,
		a.X*b.Y - b.X*a.Y,
	}
}