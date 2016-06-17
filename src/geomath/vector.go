package geomath

import "math"

var XAxis, YAxis, ZAxis = Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}

type Vector struct {
	X, Y, Z float64
}
type Point Vector
type Normal Vector

func (v *Vector) Length() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v *Vector) Scale(s float64) {
    v.X *= s
    v.Y *= s
    v.Z *= s
}

func (v *Vector) Add(vector Vector) {
    v.X += vector.X
    v.Y += vector.Y
    v.Z += vector.Z
}

func (v *Vector) subtract(vector Vector) {
    v.X -= vector.X
    v.Y -= vector.Y
    v.Z -= vector.Z
}

func (v *Vector) Normalize() {
    l := 1 / v.Length()
    v.X = v.X * l
    v.Y = v.Y * l
    v.Z = v.Z * l
}

func DotProduct(a, b Vector) float64 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
}

func CrossProduct(a, b Vector) Vector {
	return Vector{
		a.Y*b.Z - b.Y*a.Z,
		a.Z*b.X - b.Z*a.X,
		a.X*b.Y - b.X*a.Y,
	}
}

func Sum(a, b Vector) Vector {
	return Vector{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}
