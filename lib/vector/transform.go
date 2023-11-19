package vector

import "github.com/OrbitalJin/Linalgo/types"

// Apply a specific function to every element of the vector
// Performs the operation in-place
func (v *Vector) Transform(f types.Transformer) *Vector {
	for i := 0; i < v.size; i++ {
		v.data[i] = f(v.data[i])
	}
	return v
} 

// Scale the vector by a constant
// Performs the operation in-place
func (v *Vector) ScaleBy(c types.MatrixType) *Vector {
	for i := 0; i < v.size; i++ {
		v.data[i] *= c
	}
	return v
}

// Increment the vector by a constant
// Performs the operation in-place
func (v *Vector) IncrementBy(c types.MatrixType) *Vector {
	for i := 0; i < v.size; i++ {
		v.data[i] += c
	}
	return v
}