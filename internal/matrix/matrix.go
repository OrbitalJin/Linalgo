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

// Construct new Identity Matrix
func NewIdentity(size int) (*Matrix, error) {
  if size < 2 {
    return nil, fmt.Errorf("identify matrix must be square >= 2, 2")
  }
  m := New(size, size)
  for i := 0; i < size; i++ {
    m.data[i][i] = 1
  }
  return m, nil
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

// Create a new SubMatrix
func (m *Matrix) SubMatrix(start, end types.Pos) (*Matrix, error) {
  if !m.InBound(start) {
    return nil, fmt.Errorf(
      "illegal access querry start of matrix of size %d, %d: %d, %d",
      m.Rows, m.Cols,
      start.Row, start.Col,
    )
  }
  if !m.InBound(end) {
    return nil, fmt.Errorf(
      "illegal access querry end of matrix of size %d, %d: %d, %d",
      m.Rows, m.Cols,
      end.Row, end.Col,
    )
  }
  r := end.Row - start.Row + 1
  c := end.Col - start.Col + 1
  sub := New(r, c)
  for i := 0; i < r; i++ {
    for j := 0; j < c; j++ {
      sub.data[i][j] = m.data[i + start.Row][j + start.Col]
    }
  }

  return sub, nil
}

// Compute the determinant of the matrix is applicable
func (m *Matrix) Det() (float32, error) {
  if !m.IsSquare() {
    return 0, fmt.Errorf(
      "incompatible shape (%d, %d), matrix must be square to compute it's determinat",
      m.Rows, m.Cols,
    )
  }
  if m.Rows == 2 && m.Cols == 2 {
    return (m.data[0][0] * m.data[1][1] - m.data[0][1] * m.data[1][0]), nil
  }
  return 0, fmt.Errorf(
    "[NOT IMPLEMENTED YET] matrix must be 2x2 to compute it's determinant, got %dx%d",
    m.Rows, m.Cols,
  )
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
      var val float32 = 0.0
      for n := 0; n < m.Cols; n++ {
        val += m.data[r][n] * b.data[n][c]
      }
      result.data[r][c] = val
    }
  }
  return result, nil
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
  m.Transform(func (val float32) float32 {
    return -val
  })
  return m
}

// Scale matrix by value
func (m *Matrix) ScaleBy(s float32) *Matrix{
  m.Transform(func (val float32) float32 {
    return val * s
  })
  return m
}

// Apply a specific function to every element of the Matrix
func (m *Matrix) Transform(f types.Transformer) *Matrix{
  for r := 0; r < m.Rows; r++ {
    for c := 0; c < m.Cols; c++ {
      m.data[r][c] = f(m.data[r][c])
    } 
  }
  return m
}

// Returns the shape of the matrix i.e. r, c
func (m *Matrix) Shape() types.Shape {
  return types.Shape{
    Rows: m.Rows,
    Cols: m.Cols,
  }
}

// Check whether a position is within range of the matrix
func (m *Matrix) InBound(p types.Pos) bool {
  return p.Row >= 0 && p.Row < m.Rows && p.Col >= 0 && p.Col < m.Cols
}

// Checks whether a matrix is square
func (m *Matrix) IsSquare() bool {
  return m.Rows == m.Cols
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