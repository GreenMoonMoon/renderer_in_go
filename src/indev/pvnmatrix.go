package indev

import "math"

var AngRadC float64 = 180 / math.Pi
//var XAxis, YAxis, ZAxis Vector 

type Point struct {
    X, Y, Z float64
}

type Vector struct {
    X, Y, Z float64
}

type Normal struct {
    X, Y, Z float64
}

func (p *Point) Matrix44Transform(m [4][4]float64) Point {
    /*To be able to multiply the 1X3 matrix by a 4X4 matrix. We add a 
    homogeneous coordinate to the point by form of Point{X, Y, Z, 1}.
    The fourth row is used for various transformation such as projections.
    Normaly set to (0, 0, 0, 1), if different we divide all the coordinate
    by the result W to set it back at one before returning Point{x, y, z}.*/
    x := p.X * m[0][0] + p.Y * m[1][0] + p.Z * m[2][0] + m[3][0]
    y := p.X * m[0][1] + p.Y * m[1][1] + p.Z * m[2][1] + m[3][1]
    z := p.X * m[0][2] + p.Y * m[1][2] + p.Z * m[2][2] + m[3][2]
    w := p.X * m[0][3] + p.Y * m[1][3] + p.Z * m[2][3] + m[3][3]
        
    if w != 1 && w != 0 {
        return Point{x / w, y / w, z / w}
    }
    
    return Point{x, y, z}
}

//Function for simple transform without W'.
func (p *Point) Matrix43Transform(m [4][3]float64) Point {
    return Point{
        p.X * m[0][0] + p.Y * m[1][0] + p.Z * m[2][0] + m[3][0],
        p.X * m[0][1] + p.Y * m[1][1] + p.Z * m[2][1] + m[3][1],
        p.X * m[0][2] + p.Y * m[1][2] + p.Z * m[2][2] + m[3][2],
    }
}

func (v *Vector) Matrix33Transform(m [3][3]float64) Vector {
    return Vector{
        v.X * m[0][0] + v.Y * m[1][0] + v.Z * m[2][0],
        v.X * m[0][1] + v.Y * m[1][1] + v.Z * m[2][1],
        v.X * m[0][2] + v.Y * m[1][2] + v.Z * m[2][2],
    }
}

func RotMatByAngle(m *[3][3]float64, angle float64, axis Vector) {
    switch axis {
        case Vector{1, 0, 0}:
            m[1][1] = math.Cos(angle)
            m[1][2] = math.Sin(angle)
            m[2][1] = -math.Sin(angle)
            m[2][2] = math.Cos(angle)
        case Vector{0, 1, 0}:
            m[0][0] = math.Cos(angle)
            m[0][2] = -math.Sin(angle)
            m[2][0] = math.Sin(angle)
            m[2][2] = math.Cos(angle)
        case Vector{0, 0, 1}:
            m[0][0] = math.Cos(angle)
            m[0][1] = math.Sin(angle)
            m[1][0] = -math.Sin(angle)
            m[1][1] = math.Cos(angle)
    }
}

//==============================================================================

func (v *Vector) Length() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v *Vector) Normalize() {
    l := v.Length()
    v.X = v.X / l
    v.Y = v.Y / l
    v.Z = v.Z / l
}

func (v *Vector) ResizeVector(s float64) {
    v.X *= s
    v.Y *= s
    v.Z *= s
}

//Create small function to rotate, translate and resize the vector and points in order
//to understand how to regroup them in a single matrix function.
func (v *Vector) RotateVector(angle float64, axis Vector) {
    //Unit circle extend the definition of normal square triangle trigonometrie
    //to use with a circle
    //rad := angle * AngRadC
    switch axis {
        case Vector{1, 0, 0}:

        case Vector{0, 1, 0}:

        case Vector{0, 0, 1}:
            
            v.X = math.Cos(math.Pi / 2)
            v.Y = math.Sin(math.Pi / 2)
    }
    return
}