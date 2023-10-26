package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
  Rows int
  Cols int
  data [][]int
}

// Construct empty Matrix of size r, c
func New(r, c int) *Matrix {
  data := make([][]int, r)
  for i := range data {
    data[i] = make([]int, c)
  }
  return &Matrix {
    Rows: r,
    Cols: c,
    data: data,
  }
}

// Construct Matrix from string
func NewFromString(s string) *Matrix {
  data := make([][]int, strings.Count(s, ";") + 1)
  rows := strings.Split(s, ";")
  for i, n := range rows {
    values := strings.Fields(n)
    row := make([]int, len(values))

    for j, value := range values {
      value, err := strconv.Atoi(value)
      if err != nil {
        panic(err)
      }
      row[j] = value
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

// Add a matrix b to the current matrix
func (m *Matrix) Add(b *Matrix) error {
  if m.Rows != b.Rows || m.Cols != b.Cols {
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
  if m.Rows != b.Rows || m.Cols != b.Cols {
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
      fmt.Printf("%d ", val)
    }
    fmt.Println();
  }

}