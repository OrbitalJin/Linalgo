package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Matrix struct
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
	return &Matrix{
		Rows: r,
		Cols: c,
		data: data,
	}
}

// New Matrix filled with random values
func NewRandom(r, c int) *Matrix {
	data := make([][]t.MatrixType, r)
	for i := range data {
		data[i] = make([]t.MatrixType, c)
		for j := 0; j < c; j++ {
			data[i][j] = t.MatrixType(rand.Float64())
		}
	}
	return &Matrix{
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
	data := make([][]t.MatrixType, strings.Count(s, ";")+1)
	rows := strings.Split(s, ";")
	for i, n := range rows {
		values := strings.Fields(n)
		row := make([]t.MatrixType, len(values))

		for j, value := range values {
			value, err := strconv.ParseFloat(value, 64)
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
	if !m.InBound(p) {
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
func (m *Matrix) Get(p t.Pos) t.MatrixType {
	if !m.InBound(p) {
		panic(fmt.Errorf(
			"illegal access querry of matrix of size %d, %d: %d, %d",
			m.Rows, m.Cols,
			p.Row, p.Col,
		))
	}
	return m.data[p.Row][p.Col]
}

// Get a specific column
func (m *Matrix) GetCol(c int) (*Matrix, error) {
	if c >= m.Cols {
		return nil, fmt.Errorf(
			"illegal column access querry of matrix of size %d, %d: %d",
			m.Rows, m.Cols, c,
		)
	}
	// To get last row
	if c == -1 {
		c = m.Cols - 1
	}
	col := New(m.Rows, 1)
	for i := 0; i < m.Rows; i++ {
		col.data[i][0] = m.data[i][c]
	}
	return col, nil
}

// Get Maximum value
func (m *Matrix) Max() t.MatrixType {
	max := m.data[0][0]
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			if m.data[r][c] > max {
				max = m.data[r][c]
			}
		}
	}
	return max
}

// Get Minimum value
func (m *Matrix) Min() t.MatrixType {
	min := m.data[0][0]
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			if m.data[r][c] < min {
				min = m.data[r][c]
			}
		}
	}
	return min
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
			sub.data[i][j] = m.data[i+start.Row][j+start.Col]
		}
	}
	return sub, nil
}

// Compare a matrix to a scalar element-wise and returns the max
func (m *Matrix) CompMax(v t.MatrixType) *Matrix {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if v > m.data[i][j] {
				m.data[i][j] = v
			}
		}
	}
	return m
}

// Compare a matrix to a scalar element-wise and returns the min
func (m *Matrix) CompMin(v t.MatrixType) *Matrix {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if v < m.data[i][j] {
				m.data[i][j] = v
			}
		}
	}
	return m
}

// Get the sum of all of the elemtents of the matrix
func (m *Matrix) Sum() t.MatrixType {
	var total t.MatrixType
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			total += m.data[r][c]
		}
	}
	return total
}

// Compute the determinant of the matrix if applicable
func (m *Matrix) Det() (t.MatrixType, error) {
	if !m.IsSquare() {
		return 0, squareError(m.Shape(), "determinant")
	}
	return det(m), nil
}

// Returns the cofactor at the given position
func (m *Matrix) Cofactor(p t.Pos) (t.MatrixType, error) {
	if !m.IsSquare() {
		return 0, squareError(m.Shape(), "cofactors")
	}

	sign := t.MatrixType(math.Pow(-1, float64(p.Row+p.Col)))
	minor, _ := m.Minor(p)
	return sign * minor, nil
}

// Returns the Minor at the given position
func (m *Matrix) Minor(p t.Pos) (t.MatrixType, error) {
	if !m.IsSquare() {
		return 0, squareError(m.Shape(), "minors")
	}

	subMatrix := New(m.Rows-1, m.Cols-1)
	subRow := 0
	for r := 0; r < m.Rows; r++ {
		if r == p.Row {
			continue
		}
		subCol := 0
		for c := 0; c < m.Cols; c++ {
			if c == p.Col {
				continue
			}
			subMatrix.Set(t.Pos{Row: subRow, Col: subCol}, m.data[r][c])
			subCol++
		}
		subRow++
	}
	return det(subMatrix), nil
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
		fmt.Printf("%d ", c+1)
	}
	fmt.Println()
	for n, row := range m.data {
		fmt.Printf("%d ", n+1)
		for _, val := range row {
			fmt.Printf("%f ", val)
		}
		fmt.Println()
	}

}
