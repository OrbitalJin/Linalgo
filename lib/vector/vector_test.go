package vector

import (
	"testing"

	mt "github.com/OrbitalJin/Linalgo/types"
)

func TestConstruction(t *testing.T) {
	v := NewFromSlice([]mt.MatrixType{1, 2, 3})
	v.Print()
}

func TestAsString(t *testing.T) {
	v := NewFromString("1 2 3")
	v.Print()
}

func TestAdd(t *testing.T) {
	v := NewFromString("1 2 3")
	v2 := NewFromString("1 2 3")
	ans := NewFromString("2 4 6")
	v.Add(v2)
	if !v.Equals(ans) {
		t.Errorf("Vector addition failed")
		v.Print()
	}
}

func TestTransform(t *testing.T) {
	v := NewFromString("1 2 3")
	ans := NewFromString("2 4 6")
	v.Transform(func(x mt.MatrixType) mt.MatrixType { return x * 2 })
	if !v.Equals(ans) {
		t.Errorf("Vector transformation failed")
		v.Print()
	}
}

func TestMultiply(t *testing.T) {
	v := NewFromString("1 2 3")
	v2 := NewFromString("1 2 3")
	ans := NewFromString("1 4 9")
	v.Multiply(v2)
	if !v.Equals(ans) {
		t.Errorf("Vector multiplication failed")
		v.Print()
	}
}

func TestDot(t *testing.T) {
	v := NewFromString("1 2 3")
	v2 := NewFromString("1 2 3")
	ans := mt.MatrixType(14)
	dot, _ := v.Dot(v2)
	if dot != ans {
		t.Errorf("Vector dot product failed")
	}
}