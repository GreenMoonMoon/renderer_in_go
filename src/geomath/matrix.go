package geomath

import "math"

/*The func are first designed for 4 by 4 matrices.
But I believe it could be convinient to make specialized version for 3 by 3
and 2 by 2 matrices. It should help performance by avoiding certain calculation
when they are not needed.
I don't know which would be faster creating a type for each matrices or
building functions for slices.
In the slices version, the array could be unidimentionnal and we can alwas adapt
the size dynamicly.
*/

var degradC float64 = math.Pi / 180

var Identity = Mat44{
	[4]float64{1, 0, 0, 0},
	[4]float64{0, 1, 0, 0},
	[4]float64{0, 0, 1, 0},
	[4]float64{0, 0, 0, 1},
}

type Mat44 [4][4]float64

//Rotate return a rotated matrix based on Mat44 m.
func (m *Mat44) Rotate(angle float64, axis Vec3) Mat44 {
	//Convert the degree angle in radiant.
	r := angle * degradC
	  //Compute the sine and cosine of radiant r
	  //before setting the rotation matrix.
	  c := math.Cos(r)
	  s := math.Sin(r)
	  //Perform the rotation on specified axis.s
	  switch axis {
	      case XAxis:
	          result = Mat44{
	              [3]float64{1, 0, 0},
	              [3]float64{0, c, s},
	              [3]float64{0, -s, c},
	          }
	      case YAxis:
	          result = Mat44{
	              [3]float64{c, 0, s},
	              [3]float64{0, 1, 0},
	              [3]float64{ -s, 0, c},
	          }
	      case ZAxis:
	          result = Mat44{
	              [3]float64{c, s, 0},
	              [3]float64{-s, c, 0},
	              [3]float64{0, 0, 1},
	          }
	  }

	return result
}

//Mult return a vector 
func Mult(v *Vec3, m Mat44) Vec3 {
	/*To be able to multiply the 1X3 matrix by a 4X4 matrix. We add a
	  homogeneous coordinate to the point by form of Point{X, Y, Z, 1}.
	  The fourth row is used for various transformation such as projections.
	  Normaly set to (0, 0, 0, 1), if different we divide all the coordinate
	  by the result W to set it back at one before returning Point{x, y, z}.*/
	x := v[0]*m[0][0] + v[1]*m[1][0] + v[2]*m[2][0] + m[3][0]
	y := v[0]*m[0][1] + v[1]*m[1][1] + v[2]*m[2][1] + m[3][1]
	z := v[0]*m[0][2] + v[1]*m[1][2] + v[2]*m[2][2] + m[3][2]
	w := v[0]*m[0][3] + v[1]*m[1][3] + v[2]*m[2][3] + m[3][3]

	if w != 1 && w != 0 {
		return Vec3{x / w, y / w, z / w}
	}

	return Vec3{x, y, z}
}

//Function for simple transform without W'.
func MultPointMatrix43(p *Point, m [4][3]float64) Point {
	return Point{
		p.X*m[0][0] + p.Y*m[1][0] + p.Z*m[2][0] + m[3][0],
		p.X*m[0][1] + p.Y*m[1][1] + p.Z*m[2][1] + m[3][1],
		p.X*m[0][2] + p.Y*m[1][2] + p.Z*m[2][2] + m[3][2],
	}
}

//Transform by Matrix 3X3
func MultVectMatrix33(v *Vector, m [3][3]float64) Vector {
	return Vector{
		v.X*m[0][0] + v.Y*m[1][0] + v.Z*m[2][0],
		v.X*m[0][1] + v.Y*m[1][1] + v.Z*m[2][1],
		v.X*m[0][2] + v.Y*m[1][2] + v.Z*m[2][2],
	}
}

//Multiply two matrices.
//note: Similar to the dot product of two vector.
//Maybe it should be renamed MatrixDotProduct
func MultMatrices33(a, b [3][3]float64) (m [3][3]float64) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j] + a[i][2]*b[2][j]
		}
	}
	return
}

/*Matrix<T, N> & invert()
{
Matrix<T, N> mat;
for (unsigned column = 0; column < N; ++column) {
// Swap row in case our pivot point is not working
if (m[column][column] == 0) {
unsigned big = column;
for (unsigned row = 0; row < N; ++row)
if (fabs(m[row][column]) > fabs(m[big][column])) big = row;
// Print this is a singular matrix, return identity ?
if (big == column) fprintf(stderr, "Singular matrix\n");
// Swap rows
else for (unsigned j = 0; j < N; ++j) {
std::swap(m[column][j], m[big][j]);
std::swap(mat.m[column][j], mat.m[big][j]);
}
}
// Set each row in the column to 0
for (unsigned row = 0; row < N; ++row) {
if (row != column) {
T coeff = m[row][column] / m[column][column];
if (coeff != 0) {
for (unsigned j = 0; j < N; ++j) {
m[row][j] -= coeff * m[column][j];
mat.m[row][j] -= coeff * mat.m[column][j];
}
// Set the element to 0 for safety
m[row][column] = 0;
}
}
}
}
// Set each element of the diagonal to 1
for (unsigned row = 0; row < N; ++row) {
for (unsigned column = 0; column < N; ++column) {
mat.m[row][column] /= m[row][row];
}
}
*this = mat;
return *this;
}*/

//Invert a 4 by 4 matrix.
//Row elementary operation.
/*func Invert(m Matrix) Matrix {
	mi := Identity
	for column := 0; column < 4; column++ {
		if m[column][column] < math.SmallestNonzeroFloat64 {
			b := column
			for row := 0; row < 4; row++ {
				if math.Abs(m[row][column]) > math.Abs(m[b][column]) {
					b = row
				}
			}
			if b == column {
				return Identity
			} else {
				for j := 0; j < 4; j++ {
					t := m[column][j]
					m[column][j] = m[b][j]
					m[b][j] = t
					t = mi[column][j]
					mi[column][j] = mi[b][j]
					mi[b][j] = t
					//fmt.Println(mi)
				}
			}
		}
		for row := 0; row < 4; row++ {
			if row != column {
				coeff := m[row][column] / m[column][column]
				if coeff < math.SmallestNonzeroFloat64 {
					for j := 0; j < 4; j++ {
						m[row][j] -= coeff * m[column][j]
						mi[row][j] -= coeff * mi[column][j]
					}
					m[row][column] = 0
				}
			}
		}
	}
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			mi[row][column] = mi[row][column] / mi[row][row]
		}
	}
	return mi
}*/