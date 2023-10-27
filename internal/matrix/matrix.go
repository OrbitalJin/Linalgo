package matrix

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OrbitalJin/Linalgo/types"
)

type Matrix struct {
  Rows int
  Cols int
  data [][]float32
}

// Construct empty Matrix of size r, c
func New(r, c int) *Matrix {
  data := make([][]float32, r)
  for i := range data {
    data[i] = make([]float32, c)
  }
  return &Matrix {
    Rows: r,
    Cols: c,
    data: data,
  }
}

// Construct Matrix from string
func NewFromString(s string) *Matrix {
  data := make([][]float32, strings.Count(s, ";") + 1)
  rows := strings.Split(s, ";")
  for i, n := range rows {
    values := strings.Fields(n)
    row := make([]float32, len(values))

    for j, value := range values {
      value, err := strconv.Atoi(value)
      if err != nil {
        panic(err)
      }
      row[j] = float32(value)
    }
    data[i] = row
  }
  return &Matrix{
    Rows: len(data),
    Cols: len(data[0]),
    data: data,
  }
}

// TODO: Construct from file

// Set a spefic value of the Matrix
func (m *Matrix) Set(r, c int, v float32) error {
  if  r >= m.Rows || c >= m.Cols {
    return fmt.Errorf(
      "illegal access querry of matrix of size %d, %d: %d, %d",
      m.Rows, m.Cols,
      r, c,
    )
  }
  m.data[r][c] = v
  return nil
}
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

// Negate a matrix
func (m *Matrix) Negate() {
  m.Transform(func (val float32) float32 {
    return -val
  })
}

// Scale matrix by value
func (m *Matrix) ScaleBy(s float32) {
  m.Transform(func (val float32) float32 {
    return val * s
  })
}

// Apply a specific function to every element of the Matrix
func (m *Matrix) Transform(f types.Transformer) {
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      m.data[r][c] = f(m.data[r][c])
    } 
  }
}

// Checks whether the two matrices are of the same dimensions
func (m *Matrix) OfSize(b *Matrix) bool {
  return m.Rows == b.Rows && m.Cols == b.Cols
}

// Check whetehr two matrices are equal i.e. a[i][j] == b[i][j]
func (m *Matrix) Equals(b *Matrix) bool {
  if !m.OfSize(b) {
    return false
  }
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      if m.data[r][c] != b.data[r][c] {
        return false
      }
    } 
  }
  return true
}

// Print the content of the matrix
func (m *Matrix) Print() {
  fmt.Print("* ")
  for c := 0; c < m.Cols; c++ {
    fmt.Printf("%d ", c + 1);
  }
  fmt.Println()
  for n, row := range m.data {
    fmt.Printf("%d ", n + 1)
    for _, val := range row {
      fmt.Printf("%f ", val)
    }
    fmt.Println();
  }

}