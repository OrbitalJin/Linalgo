package vector

import (
	"fmt"
	"strconv"
	"strings"

	t "github.com/OrbitalJin/Linalgo/types"
)

type Vector struct {
	data []t.MatrixType
	size int
}

func New(size int) *Vector{
	return &Vector{
		data: make([]t.MatrixType, size),
		size: size,
	}
}

func NewFromString(s string) *Vector {
	values := strings.Fields(s)
	data := make([]t.MatrixType, len(values))
	for i, v := range values {
		v, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		data[i] = t.MatrixType(v)
	}
	return &Vector{
		data: data,
		size: len(values),
	}
}

func NewFromSlice(data []t.MatrixType) *Vector {
	return &Vector{
		data: data,
		size: len(data),
	}
}

func (v *Vector) AsString() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v)), ","), "[]")
}

func (v *Vector) Print() {
	for e := range v.data {
		fmt.Printf("%d ", e);
	}
	fmt.Println()
}

// Checks whether two vectors are equal in size
func (v *Vector) OfSize(v2 *Vector) bool {
	return v.size == v2.size
}