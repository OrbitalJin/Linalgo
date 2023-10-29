package matrix

import (
	"fmt"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Adjoint of a Matrix (Matrix of det)
func (m *Matrix) Adj() (*Matrix, error) {
  if !m.IsSquare() {
    return nil, fmt.Errorf(
      "incompatible shape (%d, %d), matrix must be square to compute it's determinat",
      m.Rows, m.Cols,
    )
  }

  newMat := New(m.Cols, m.Rows)
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      p := t.Pos{Row: r, Col: c}
      cofactor, err := m.Cofactor(p)
      if err != nil {
        return nil, err
      }
      newMat.Set(p, cofactor)
    }
  }
  return newMat.T(), nil
}



// Transpose Matrix
func (m *Matrix) T() *Matrix{
  newMat := New(m.Cols, m.Rows)
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      newMat.data[c][r] = m.data[r][c]
    }
  }
  m.data = newMat.data
  m.Rows = newMat.Rows
  m.Cols = newMat.Cols
  newMat = nil
  return m
}

// Negate a matrix
func (m *Matrix) Negate() *Matrix{
  m.Transform(func (val t.MatrixType) t.MatrixType {
    return -val
  })
  return m
}

// Scale matrix by value
func (m *Matrix) ScaleBy(s t.MatrixType) *Matrix{
  m.Transform(func (val t.MatrixType) t.MatrixType {
    return val * s
  })
  return m
}

// Apply a specific function to every element of the Matrix
func (m *Matrix) Transform(f t.Transformer) *Matrix{
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      m.data[r][c] = f(m.data[r][c])
    } 
  }
  return m
}