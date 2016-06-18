package geomath

import "math"

//Degree to radiant ration. 
var degradC float64 = math.Pi / 180

var Identity [3][3]float64 = [3][3]float64{
    [3]float64{1, 0, 0},
    [3]float64{0, 1, 0},
    [3]float64{0, 0, 1},
}

//Return a rotation matrix.
func RotationMatrix(angle float64, axis Vector) (m [3][3]float64) {
    //Convert the degree angle in radiant.
    r := angle * degradC
    //Compute the sine and cosine of radiant r
    //before setting the rotation matrix.
    c := math.Cos(r)
    s := math.Sin(r)
    //Perform the rotation on specified axis.
    switch axis {
        case XAxis:
            m = [3][3]float64{
                [3]float64{1, 0, 0},
                [3]float64{0, c, s},
                [3]float64{0, -s, c},
            }
        case YAxis:
            m = [3][3]float64{
                [3]float64{c, 0, s},
                [3]float64{0, 1, 0},
                [3]float64{ -s, 0, c},
            }
        case ZAxis:
            m = [3][3]float64{
                [3]float64{c, s, 0},
                [3]float64{-s, c, 0},
                [3]float64{0, 0, 1},
            }
    }
    
    return
}

func TransformMatrix44(p *Point, m [4][4]float64) Point {
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
func TransformMatrix43(p *Point, m [4][3]float64) Point {
    return Point{
        p.X * m[0][0] + p.Y * m[1][0] + p.Z * m[2][0] + m[3][0],
        p.X * m[0][1] + p.Y * m[1][1] + p.Z * m[2][1] + m[3][1],
        p.X * m[0][2] + p.Y * m[1][2] + p.Z * m[2][2] + m[3][2],
    }
}

//Transform by Matrix 3X3
func TransformMatrix33(v *Vector, m [3][3]float64) Vector {
    return Vector{
        v.X * m[0][0] + v.Y * m[1][0] + v.Z * m[2][0],
        v.X * m[0][1] + v.Y * m[1][1] + v.Z * m[2][1],
        v.X * m[0][2] + v.Y * m[1][2] + v.Z * m[2][2],
    }
}

//Multiply two matrices.
//note: Similar to the dot product of two vector.
//Maybe it should be renamed MatrixDotProduct
func MultMatrices33(a, b [3][3]float64) (m [3][3]float64) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            m[i][j] = a[i][0] * b[0][j] + a[i][1] * b[1][j] + a[i][2] * b[2][j]
        }
    }
    return
}