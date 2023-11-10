package matrix

import (
	"fmt"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Add a matrix b to the current matrix
func (m *Matrix) Add(b *Matrix) error {
	if !m.OfSize(b) {
		return fmt.Errorf(
			"size mismatch, cannot add matrix %d, %d to matrix %d, %d",
			b.Rows, b.Cols,
			m.Rows, m.Cols,
		)
	}
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			m.data[r][c] += b.data[r][c]
		}
	}
	return nil
}

// Substract a matrix from the current matrix
func (m *Matrix) Sub(b *Matrix) error {
	if !m.OfSize(b) {
		return fmt.Errorf(
			"size mismatch, cannot substract matrix %d, %d from matrix %d, %d",
			b.Rows, b.Cols,
			m.Rows, m.Cols,
		)
	}
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			m.data[r][c] -= b.data[r][c]
		}
	}
	return nil
}

// Dot Product
func (m *Matrix) Dot(b *Matrix) (*Matrix, error) {
	if m.Cols != b.Rows {
		return nil, fmt.Errorf(
			"size mismatch, cannot apply dot product between Matrices of dimension %d, %d and %d, %d",
			m.Rows, m.Cols,
			b.Rows, b.Cols,
		)
	}
	result := New(m.Rows, b.Cols)
	for r := 0; r < result.Rows; r++ {
		for c := 0; c < result.Cols; c++ {
			var val t.MatrixType = 0.0
			for n := 0; n < m.Cols; n++ {
				val += m.data[r][n] * b.data[n][c]
			}
			result.data[r][c] = val
		}
	}
	return result, nil
}

// Element-wise product
func (m *Matrix) Mul(b *Matrix) error {
	if !m.OfSize(b) {
		return fmt.Errorf(
			"size mismatch, cannot apply element-wise multiplication between matrix %d, %d and matrix %d, %d",
			b.Rows, b.Cols,
			m.Rows, m.Cols,
		)
	}
	for r := 0; r < m.Rows; r++ {
		for c:= 0; c < m.Cols; c++ {
			m.data[r][c] *= b.data[r][c]
		}
	}
	return nil
}

// Matrix `Division` A/B => A.dot(B.Inv())
func (m *Matrix) Div(b *Matrix) (*Matrix, error) {
	// B must be square, that case is handled inside inv
	invB, err := b.Inverse()
	if err != nil {
		return nil, err
	}
	res, err := m.Dot(invB)
	if err != nil {
		return nil, err
	}
	return res, nil
}
