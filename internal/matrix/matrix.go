package matrix

import (
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