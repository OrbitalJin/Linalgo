package vector

import (
	"fmt"

	t "github.com/OrbitalJin/Linalgo/types"
)

// Add two vectors together
func (v *Vector) Add(v2 *Vector) error {
	if !v.OfSize(v2) {
		return fmt.Errorf("Vectors are not of the same size")
	}
	for i := 0; i < v.size; i++ {
		v.data[i] += v2.data[i]
	}
	return nil
}

// Checks whether two vectors are equal
func (v *Vector) Equals(v2 *Vector) bool {
	if !v.OfSize(v2) {
		return false
	}
	for i := 0; i < v.size; i++ {
		if v.data[i] != v2.data[i] {
			return false
		}
	}
	return true
}

// Subtract two vectors
func (v *Vector) Subtract(v2 *Vector) error {
	if !v.OfSize(v2) {
		return fmt.Errorf("Vectors are not of the same size")
	}
	for i := 0; i < v.size; i++ {
		v.data[i] -= v2.data[i]
	}
	return nil
}

// Multiply two vectors
func (v *Vector) Multiply(v2 *Vector) error {
	if !v.OfSize(v2) {
		return fmt.Errorf("Vectors are not of the same size")
	}
	for i := 0; i < v.size; i++ {
		v.data[i] *= v2.data[i]
	}
	return nil
}

// Dot product of two vectors
func (v *Vector) Dot(v2 *Vector) (t.MatrixType, error) {
	if !v.OfSize(v2) {
		return 0, fmt.Errorf("Vectors are not of the same size")
	}
	var sum t.MatrixType
	for i := 0; i < v.size; i++ {
		sum += v.data[i] * v2.data[i]
	}
	return sum, nil
}

// Get the maximum value of the vector
func (v *Vector) Max() t.MatrixType {
	var max t.MatrixType
	for _, e := range v.data {
		if e > max {
			max = e
		}
	}
	return max
}

// Get the cumulative sum of the vector
func (v *Vector) Sum() *Vector {
	for i := 1; i < v.size; i++ {
		v.data[i] += v.data[i-1]
	}
	return v
}