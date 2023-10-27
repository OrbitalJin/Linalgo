package matrix

import (
	"testing"
)

func TestSum(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	resMat := NewFromString("2 4 6 ; 8 10 12 ; 14 16 18")
	mat1.Add(mat2)
	if !mat1.Equals(resMat) {
		t.Errorf("Matrix addition failed")
	}
}

func TestSub(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	resMat := NewFromString("0 0 0 ; 0 0 0 ; 0 0 0")
	mat1.Sub(mat2)
	if !mat1.Equals(resMat) {
		t.Errorf("Matrix Substraction failed")
	}
}

func TestScaleBy(t *testing.T) {
	var scalar float32 = -2.0
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	resMat := NewFromString("-2 -4 -6 ; -8 -10 -12 ; -14 -16 -18")
	mat.ScaleBy(scalar)
	if !mat.Equals(resMat) {
		t.Errorf("Matrix Scaling failed")
	}
}

func TestSet(t *testing.T) {
	var value float32 = -69.0
	r := 1
	c := 1
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	resMat := NewFromString("1 2 3 ; 4 -69 6 ; 7 8 9")
	err := mat.Set(r, c, value)
	if err != nil || !mat.Equals(resMat) {
		mat.Print()
		t.Errorf("Matrix Setting Failed: %s", err)
	}
}

func TestNegate(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	resMat := NewFromString("-1 -2 -3 ; -4 -5 -6 ; -7 -8 -9")
	mat.Negate()
	if !mat.Equals(resMat) {
		t.Errorf("Matrix Negation Failed")
	}
}