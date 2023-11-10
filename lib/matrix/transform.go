package matrix

import (
	"fmt"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Inverse of a Matrix
// Returns  pointer to a new matrix
func (m *Matrix) Inverse() (*Matrix, error) {
	if !m.IsSquare() {
		return nil, squareError(m.Shape(), "Inverse")
	}

	det, err := m.Det()
	if err != nil {
		return nil, err
	}
	if det == 0 {
		return nil, fmt.Errorf("Matrix is singular (non-invertible), determinant is zero")
	}
	adj, err := m.Adj()
	if err != nil {
		return nil, err
	}
	return adj.ScaleBy(t.MatrixType(1 / det)), nil
}

// Augment a matrix
// Returns a pointer to a new matrix
func (m *Matrix) Augment(b *Matrix) (*Matrix, error) {
	if m.Rows != b.Rows {
		return nil, fmt.Errorf("augmentation: row size mismatch, r1=%d, r2=%d", m.Rows, b.Rows)
	}
	augmented := New(m.Rows, m.Cols+b.Cols)
	for i := 0; i < m.Rows; i++ {
		augmented.data[i] = append(m.data[i], b.data[i]...)
	}
	return augmented, nil
}

// Adjugate of a Matrix (Matrix of det)
// Returns a pointer to a new matrix
func (m *Matrix) Adj() (*Matrix, error) {
	if !m.IsSquare() {
		return nil, squareError(m.Shape(), "Adjugate")
	}
	// Base Case 2x2
	if m.Rows == 2 {
		m.data[0][0], m.data[1][1] = m.data[1][1], m.data[0][0]
		m.data[0][1], m.data[1][0] = -m.data[0][1], -m.data[1][0]
		return m, nil
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

// Reduced Row Echelon Form: Gaussian Elimination
// Performs the operation in place
func (m *Matrix) GaussRREF() (*Matrix, error) {
	lead := 0
	for r := 0; r < m.Rows; r++ {
		if lead >= m.Cols {
			return m, nil
		}
		i := r
		for m.data[i][lead] == 0 {
			i++
			if i == m.Rows {
				i = r
				lead++
				if lead == m.Cols {
					return m, nil
				}
			}
		}
		m.SwapRows(i, r)
		f := 1 / m.data[r][lead]
		m.ScaleRowBy(r, f)
		for i = 0; i < m.Rows; i++ {
			if i != r {
				f = m.data[i][lead]
				for j, e := range m.data[r] {
					m.data[i][j] -= e * f
				}
			}
		}
		lead++
	}
	return m, nil
}

// Transpose Matrix
// Returns a pointer to a new matrix
func (m *Matrix) T() *Matrix {
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

// Increment all elements by value
// Performs the operation in place
func (m *Matrix) IncrementBy(v t.MatrixType) *Matrix {
	m.Transform(func(elem t.MatrixType) t.MatrixType {
		return elem + v
	})
	return m
}

// Negate a matrix
func (m *Matrix) Negate() *Matrix {
	m.Transform(func(val t.MatrixType) t.MatrixType {
		return -val
	})
	return m
}

// Scale matrix by value
// Performs the operation in place
func (m *Matrix) ScaleBy(s t.MatrixType) *Matrix {
	m.Transform(func(val t.MatrixType) t.MatrixType {
		return val * s
	})
	return m
}

// Apply a specific function to every element of the Matrix
// Performs the operation in place
func (m *Matrix) Transform(f t.Transformer) *Matrix {
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			m.data[r][c] = f(m.data[r][c])
		}
	}
	return m
}

// Clip matrix values to upper and lower values given
// Performs the operation in place
func (m *Matrix) Clip(low, up t.MatrixType) *Matrix {
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			if m.data[r][c] > up {
				m.data[r][c] = up
			} else if m.data[r][c] < low {
				m.data[r][c] = low
			}
		}
	}
	return m
}

// Swap rows of a matrix
// Performs the operation in place
func (m *Matrix) SwapRows(r1, r2 int) (*Matrix, error) {
	if r1 < 0 || r1 >= m.Rows || r2 < 0 || r2 >= m.Rows {
		return nil, fmt.Errorf("row Swapping: index out of bounds")
	}
	m.data[r1], m.data[r2] = m.data[r2], m.data[r1]
	return m, nil
}

// Transform Row
// Performs the operation in place
func (m *Matrix) TransformRow(r int, f t.Transformer) (*Matrix, error) {
	if r < 0 || r >= m.Rows {
		return nil, fmt.Errorf("row Transformation: index out of bounds")
	}
	for c := 0; c < m.Cols; c++ {
		m.data[r][c] = f(m.data[r][c])
	}
	return m, nil
}

// Scale Row
// Performs the operation in place
func (m *Matrix) ScaleRowBy(r int, s t.MatrixType) (*Matrix, error) {
	_, err := m.TransformRow(r, func(val t.MatrixType) t.MatrixType {
		return val * s
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}