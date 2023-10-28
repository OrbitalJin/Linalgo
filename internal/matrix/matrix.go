package matrix

import (
	"fmt"
	"strconv"
	"strings"

	t "github.com/OrbitalJin/Linalgo/types"
)

type Matrix struct {
  Rows int
  Cols int
  data [][]t.MatrixType
}

// Construct empty Matrix of size r, c
func New(r, c int) *Matrix {
  data := make([][]t.MatrixType, r)
  for i := range data {
    data[i] = make([]t.MatrixType, c)
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
  data := make([][]t.MatrixType, strings.Count(s, ";") + 1)
  rows := strings.Split(s, ";")
  for i, n := range rows {
    values := strings.Fields(n)
    row := make([]t.MatrixType, len(values))

    for j, value := range values {
      value, err := strconv.Atoi(value)
      if err != nil {
        panic(err)
      }
      row[j] = t.MatrixType(value)
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
func (m *Matrix) Set(p t.Pos, v t.MatrixType) error {
  if  !m.InBound(p) {
    return fmt.Errorf(
      "illegal access querry of matrix of size %d, %d: %d, %d",
      m.Rows, m.Cols,
      p.Row, p.Col,
    )
  }
  m.data[p.Row][p.Col] = v
  return nil
}

// Get a specific value of the Matrix
func (m *Matrix) Get(p t.Pos) (t.MatrixType, error) {
  if !m.InBound(p) {
    return 0, fmt.Errorf(
      "illegal access querry of matrix of size %d, %d: %d, %d",
      m.Rows, m.Cols,
      p.Row, p.Col,
    )
  }
  return m.data[p.Row][p.Col], nil
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
func (m *Matrix) SubMatrix(start, end t.Pos) (*Matrix, error) {
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
func (m *Matrix) Det() (t.MatrixType, error) {
  if !m.IsSquare() {
    return 0, fmt.Errorf(
      "incompatible shape (%d, %d), matrix must be square to compute it's determinat",
      m.Rows, m.Cols,
    )
  }
  switch m.Rows {
  case 2: return m.det2x2(), nil
  case 3: return m.det3x3(), nil
  }
  return 0, fmt.Errorf(
    "[NOT IMPLEMENTED YET] matrix must be 2x2 or 3x3 to compute it's determinant, got %dx%d",
    m.Rows, m.Cols,
  )
}

// Computer the determinant of a 2x2 matrix
func (m *Matrix) det2x2() t.MatrixType {
  a := m.data[0][0]
  b := m.data[0][1]
  c := m.data[1][0]
  d := m.data[1][1]
  return a * d - b * c
}

// Compute the determinant of a 3x3 matrix applying the Leibniz formula
func (m *Matrix) det3x3() t.MatrixType {
  aei := m.data[0][0] * m.data[1][1] * m.data[2][2] 
  bfg := m.data[0][1] * m.data[1][2] * m.data[2][0] 
  cdh := m.data[0][2] * m.data[1][0] * m.data[2][1] 
  ceg := m.data[0][2] * m.data[1][1] * m.data[2][0] 
  bdi := m.data[0][1] * m.data[1][0] * m.data[2][2] 
  afh := m.data[0][0] * m.data[1][2] * m.data[2][1] 
  return aei + bfg + cdh - ceg - bdi - afh
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

// Returns the shape of the matrix i.e. r, c
func (m *Matrix) Shape() t.Shape {
  return t.Shape{
    Rows: m.Rows,
    Cols: m.Cols,
  }
}

// Check whether a position is within range of the matrix
func (m *Matrix) InBound(p t.Pos) bool {
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