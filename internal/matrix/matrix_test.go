package matrix

import (
	"testing"
)

func TestSum(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("2 4 6 ; 8 10 12 ; 14 16 18")
	mat1.Add(mat2)
	if !mat1.Equals(res) {
		t.Errorf("Matrix addition failed")
	}
}

func TestSub(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("0 0 0 ; 0 0 0 ; 0 0 0")
	mat1.Sub(mat2)
	if !mat1.Equals(res) {
		t.Errorf("Matrix Substraction failed")
	}
}

func TestScaleBy(t *testing.T) {
	var scalar float32 = -2.0
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("-2 -4 -6 ; -8 -10 -12 ; -14 -16 -18")
	mat.ScaleBy(scalar)
	if !mat.Equals(res) {
		t.Errorf("Matrix Scaling failed")
	}
}

func TestSet(t *testing.T) {
	var value float32 = -69.0
	r := 1
	c := 1
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("1 2 3 ; 4 -69 6 ; 7 8 9")
	err := mat.Set(r, c, value)
	if err != nil || !mat.Equals(res) {
		mat.Print()
		t.Errorf("Matrix Setting Failed: %s", err)
	}
}

func TestNegate(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("-1 -2 -3 ; -4 -5 -6 ; -7 -8 -9")
	mat.Negate()
	if !mat.Equals(res) {
		t.Errorf("Matrix Negation Failed")
	}
}

func TestTranspose(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res := NewFromString("1 4 7 ; 2 5 8 ; 3 6 9")
	mat.T()
	if !mat.Equals(res) || mat.Rows != res.Rows || mat.Cols != res.Cols {
		t.Errorf("Matrix Transposition Failed")
	}
}

func TestDot(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6")
	mat2 := NewFromString("-1 -2 ; -3 -4 ; -5 -6")
	res := NewFromString("-22 -28 ; -49 -64")
	mat3, err := mat1.Dot(mat2)
	if !mat3.Equals(res) || err != nil {
		t.Errorf("Matrix Transposition (1) Failed: %s", err)
	}
	mat1 = NewFromString("2 1 ; 4 5")
	mat2 = NewFromString("2 6 1 ; 1 2 3")
	res = NewFromString("5 14 5 ; 13 34 19")
	mat3, err = mat1.Dot(mat2)
	if !mat3.Equals(res) || err != nil {
		t.Errorf("Matrix Transposition (2) Failed: %s", err)
	}

}

func TestIdentity(t *testing.T) {
	// Sample Tests
	I3, err := NewIdentity(3);
	res := NewFromString("1 0 0 ; 0 1 0 ; 0 0 1")
	if !I3.Equals(res) || err != nil {
		t.Errorf("Matrix Identity (1) Failed: %s", err)
	}
	I2, err := NewIdentity(2);
	res = NewFromString("1 0 ; 0 1")
	if !I2.Equals(res) || err != nil {
		t.Errorf("Matrix Identity (2) Failed: %s", err)
	}
	// Neutral Property
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res, err = mat.Dot(I3);
	if !mat.Equals(res) || err != nil {
		t.Errorf("Matrix Identity (3) Failed: %s", err)
	}
}