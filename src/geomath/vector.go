package geomath

import "math"

//Basic variables for axis
var XAxis, YAxis, ZAxis = Vec3{1, 0, 0}, Vec3{0, 1, 0}, Vec3{0, 0, 1}
var Zero = Vec3{0, 0, 0}

//Vec3 type represent a 3 component vector in euclidean space.
type Vec3 [3]float64

func (v *Vec3) X() float64 { return v[0] } //X return the x element of Vec3 v.
func (v *Vec3) Y() float64 { return v[1] } //Y return the y element of Vec3 v.
func (v *Vec3) Z() float64 { return v[2] } //Z return the z element of Vec3 v.

//NewVec3 returns a new Vec3 build with x, y, z arguments.
func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

//Length returns the length of Vec3 v.
func Length(v Vec3) float64 {
	return math.Sqrt(Dot(*v, *v))
}

//Unit returns the normalized vector of Vec3 v.
func Unit(v Vec3) Vec3 {
	l := 1 / v.Length()
	return Vec3{
		v[0] * l,
		v[1] * l,
		v[2] * l,
	}
}

//Add return the sum of Vec3 a and b.
func Add(a, b Vec3) Vec3 {
	return Vec3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

//Subtract returns the difference between Vec3 a and b.
func Subtract(a, b Vec3) Vec3 {
	return Vec3{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func Rotate(vector Vec3, axis Vec3, angle float64) Vec3 {
	r := angle * degradC
	return Add(Scale(vector, math.Sin(r)), Add(Scale(Cross(axis, vector), math.Sin(r)), Scale(Scale(axis, Dot(axis, vector)), (1 - math.Cos(r)))))
}

//Theta return the theta angle of Vec3 v in a spherical coordinate system.
//Note: We assume that the vector is normalized.
func Theta(v Vec3) float64 {
	//The vector should be normalized. We don't normalize it in this
	//function because the current function will affect the vector itself.
	return math.Acos(v[2])
}

//Phi retrurn the phi angle of vec3 v in a spherical coordinate system.
func Phi(v Vec3) float64 {
	p := math.Atan2(v[1], v[0])
	if p < 0 {
		return p + 2*math.Pi
	} else {
		return p
	}
}

//Scale return a scaled vector Vec3 v by scalar s.
func Scale(v Vec3, s float64) Vec3 {
	return Vec3{
		v[0] * s,
		v[1] * s,
		v[2] * s,
	}
}

//Dot return the dot proctuct of Vec3 a and b.
func Dot(a, b Vec3) float64 {
	return (a[0]*b[0] + a[1]*b[1] + a[2]*b[2])
}

//Cross return the cross product of Vec3 a and b.
func Cross(a, b Vec3) Vec3 {
	return Vec3{
		a[1]*b[2] - b[1]*a[2],
		a[2]*b[0] - b[2]*a[0],
		a[0]*b[1] - b[0]*a[1],
	}
}

//Box return the box product (Scalar triple product) of Vec3 a, b and c.
func Box(a, b, c Vec3) float64 {
	//The box product is also the determinant of a 3 by 3 matrix composed
	//of vectors.
	return Dot(a, Cross(b, c))
}

//Equal return true if Vec3 a is equal to Vec3 b. False otherwise.
func Equal(a, b Vec3) bool {
	if a[0] == b[0] && a[1] == b[1] && a[2] == b[2] {
		return true
	} else {
		return false
	}
}

//Opposite return true if Vec3 a is the opposite of Vec3 b. False otherwise.
func Opposite(a, b Vec3) bool {
	if a[0] == -b[0] && a[1] == -b[1] && a[2] == -b[2] {
		return true
	} else {
		return false
	}
}
