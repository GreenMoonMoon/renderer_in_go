package geomath

import "math"

type Vector struct {
    X, Y, Z float64
}

func Length(v Vector) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func Normal(v Vector) Vector {
    l := 1 / Length(v)
    return Vector{v.X * l, v.Y * l, v.Z * l}
}

func DotProduct(a, b Vector) float64 {
    return (a.X * b.X + a.Y * b.Y + a.Z * b.Z)
}

func CrossProduct(a, b Vector) Vector {
    return Vector{
        a.Y * b.Z - b.Y * a.Z,
        a.Z * b.X - b.Z * a.X,
        a.X * b.Y - b.X * a.Y,
    }
}