package matrix

import (
	"math"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Recursive function to compute the determinant of a matrix based on the Laplace expansion
func det(m *Matrix) t.MatrixType {
	// Base Case
  if m.Rows == 2 {
    return m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
  }
  var determinant t.MatrixType
  // Iterate over the matrix
  for i := 0; i < m.Rows; i++ {
    for j := 0; j < m.Cols; j++ {
      // Set the pivot being m.data[i][j] and the sign (+/-)
      pivot := t.MatrixType(math.Pow(-1, float64(i + j))) * m.data[i][j]
      // Make a submatrix of dimensions n-1
      subMatrix := New(m.Rows - 1, m.Cols - 1)
      // Fill up the submatrix excluding the pivot's column and row I manually increment the submatrix's
      // Row and column because otherwise it will offset the submatrix when skipping the pivot's row and column
      subRow := 0
      for r := 0; r < m.Rows; r++ {
          subCol := 0
          for c := 0; c < m.Cols; c++ {
            if r != i && c != j {
              subMatrix.Set(t.Pos{Row: subRow, Col: subCol}, m.data[r][c])
              subCol++
            }
          }
          subRow++
      }
      subDet := det(subMatrix)
      // Multiply by the determinant of that submatrix resursivley
      pivot *= subDet
      // Add the result to the determinant
      determinant += pivot
    }
  }
  return determinant
}