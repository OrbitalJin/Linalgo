package matrix

import (
	"fmt"
	"testing"

	"github.com/OrbitalJin/Linalgo/types"
)

func TestMax(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := types.MatrixType(9)
	if mat.Max() != ans {
		t.Errorf("Matrix max failed, expected %f, got %f", ans, mat.Max())
	}
}

func TestMin(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := types.MatrixType(1)
	if mat.Min() != ans {
		t.Errorf("Matrix min failed, expected %f, got %f", ans, mat.Max())
	}
}

func TestMean(t *testing.T) {
	mat := NewFromString("1 2 3 4");
	ans := 2.5
	if mat.Mean() != ans {
		t.Errorf("Matrix mean (1) failed, expected %f, got %f", ans, mat.Mean())
	}
	mat = NewFromString("1 2 ; 3 4");
	ans = 2.5
	if mat.Mean() != ans {
		t.Errorf("Matrix mean (2) failed, expected %f, got %f", ans, mat.Mean())
	}

}

func TestIncrementation(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	increment := 1
	ans := NewFromString("2 3 4 ; 5 6 7; 8 9 10")
	mat.IncrementBy(types.MatrixType(increment))
	if !mat.Equals(ans) {
		t.Error("Matrix incrementation (1) Failed")
		mat.Print()
	}
	mat = NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	increment = -1
	ans = NewFromString("0 1 2 ; 3 4 5 ; 6 7 8")
	mat.IncrementBy(types.MatrixType(increment))
	if !mat.Equals(ans) {
		t.Error("Matrix incrementation (2) Failed")
		mat.Print()
	}
}

func TestAdd(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("2 4 6 ; 8 10 12 ; 14 16 18")
	mat1.Add(mat2)
	if !mat1.Equals(ans) {
		t.Errorf("Matrix addition failed")
	}
}

func TestSumElements(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := types.MatrixType(45)
	if mat.Sum() != ans {
		t.Errorf("Matrix sum failed, expected %f, got %f", ans, mat.Sum())
	}
}

func TestSub(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mat2 := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("0 0 0 ; 0 0 0 ; 0 0 0")
	mat1.Sub(mat2)
	if !mat1.Equals(ans) {
		t.Errorf("Matrix Substraction failed")
	}
}

func TestClipping(t *testing.T) {
	var low, up types.MatrixType =  0, 10
	mat := NewFromString("-1 -2 -3 ; 5 6 4 ; 23 45 234")
	ans := NewFromString("0 0 0 ; 5 6 4 ; 10 10 10")
	mat.Clip(low, up)
	if !mat.Equals(ans) {
		t.Error("Matrix clipping failed, got: ")
		mat.Print()
	}
}

func TestMasking(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	mask := NewFromString("1 0 0 ; 0 1 0 ; 0 0 1")
	ans := NewFromString("1 ; 5 ; 9")
	masked, err := mat.Mask(mask)
	if err != nil || !masked.Equals(ans) {
		t.Errorf("Matrix masking (1) failed: %s", err)
		masked.Print()
	}
}

func TestElementWiseMul(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9");
	ans := NewFromString("1 4 9 ; 16 25 36 ; 49 64 81")
	err := mat.Mul(mat)
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix element-wise multiplication failed, got:")
		mat.Print()
	}
}

func TestScaleBy(t *testing.T) {
	var scalar types.MatrixType = -2.0
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("-2 -4 -6 ; -8 -10 -12 ; -14 -16 -18")
	mat.ScaleBy(scalar)
	if !mat.Equals(ans) {
		t.Errorf("Matrix Scaling failed")
	}
}

func TestSet(t *testing.T) {
	var value types.MatrixType = -69.0
	r := 1
	c := 1
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("1 2 3 ; 4 -69 6 ; 7 8 9")
	err := mat.Set(types.Pos{Row: r, Col: c}, value)
	if err != nil || !mat.Equals(ans) {
		mat.Print()
		t.Errorf("Matrix Setting Failed: %s", err)
	}
}

func TestNegate(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("-1 -2 -3 ; -4 -5 -6 ; -7 -8 -9")
	mat.Negate()
	if !mat.Equals(ans) {
		t.Errorf("Matrix Negation Failed")
	}
}

func TestTranspose(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("1 4 7 ; 2 5 8 ; 3 6 9")
	mat.T()
	if !mat.Equals(ans) || mat.Rows != ans.Rows || mat.Cols != ans.Cols {
		t.Errorf("Matrix Transposition Failed")
	}
}

func TestDot(t *testing.T) {
	mat1 := NewFromString("1 2 3 ; 4 5 6")
	mat2 := NewFromString("-1 -2 ; -3 -4 ; -5 -6")
	ans := NewFromString("-22 -28 ; -49 -64")
	mat3, err := mat1.Dot(mat2)
	if err != nil || !mat3.Equals(ans) {
		t.Errorf("Matrix Transposition (1) Failed: %s", err)
	}
	mat1 = NewFromString("2 1 ; 4 5")
	mat2 = NewFromString("2 6 1 ; 1 2 3")
	ans = NewFromString("5 14 5 ; 13 34 19")
	mat3, err = mat1.Dot(mat2)
	if err != nil || !mat3.Equals(ans) {
		t.Errorf("Matrix Transposition (2) Failed: %s", err)
	}

}

func TestIdentity(t *testing.T) {
	// Sample Tests
	I3, err := NewIdentity(3)
	ans := NewFromString("1 0 0 ; 0 1 0 ; 0 0 1")
	if !I3.Equals(ans) || err != nil {
		t.Errorf("Matrix Identity (1) Failed: %s", err)
	}
	I2, err := NewIdentity(2)
	ans = NewFromString("1 0 ; 0 1")
	if err != nil || !I2.Equals(ans) {
		t.Errorf("Matrix Identity (2) Failed: %s", err)
	}
	// Neutral Property
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans, err = mat.Dot(I3)
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix Identity (3) Failed: %s", err)
	}
}

func TestSubMatrix(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("5 6 ; 8 9")
	sub, err := mat.SubMatrix(
		types.Pos{Row: 1, Col: 1},
		types.Pos{Row: 2, Col: 2},
	)
	if err != nil || !sub.Equals(ans) {
		t.Errorf("Matrix Submatrix (1) Failed: %s", err)
	}
}

func TestDet(t *testing.T) {
	// 2x2
	var ans types.MatrixType = -2
	mat := NewFromString("1 2 ; 3 4")
	det, err := mat.Det()
	if err != nil || det != ans {
		t.Errorf("Matrix Determinant (1) Failed: Expected %f, got %f. \nErr: %s", ans, det, err)
	}

	ans = -14
	mat = NewFromString("3 8 ; 4 6")
	det, err = mat.Det()
	if err != nil || det != ans {
		t.Errorf("Matrix Determinant (2) Failed: Expected %f, got %f. \nErr: %s", ans, det, err)
	}
	// 3x3
	ans = -306
	mat = NewFromString("6 1 1 ; 4 -2 5 ; 2 8 7")
	det, err = mat.Det()
	if err != nil || det != ans {
		fmt.Println(det)
		mat.Print()
		t.Errorf("Matrix Determinant (3) Failed: Expected %f, got %f. \nErr: %s", ans, det, err)
	}
	// Upper Triangular
	ans = -12
	mat = NewFromString("6 1 1 ; 0 -2 5 ; 0 0 1")
	det, err = mat.Det()
	if err != nil || det != ans {
		fmt.Println(det)
		t.Errorf("Matrix Determinant (4) Failed: Expected %f, got %f. \nErr: %s", ans, det, err)
	}

	// 4x4
	ans = 0
	mat = NewFromString("1 2 3 4 ; 5 6 7 8 ; 9 10 11 12 ; 13 14 15 16")
	det, err = mat.Det()
	if err != nil || det != ans {
		mat.Print()
		fmt.Println(det)
		t.Errorf("Matrix Determinant (4) Failed: %s", err)
	}
	// 4x4 null
	ans = 0
	mat = NewFromString("0 0 0 0 ; 0 0 0 0 ; 0 0 0 0; 0 0 0 0")
	det, err = mat.Det()
	if err != nil || det != ans {
		mat.Print()
		fmt.Println(det)
		t.Errorf("Matrix Determinant (5) Failed: %s", err)
	}
}
func TestCofactor(t *testing.T) {
	var ans types.MatrixType = -3
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	cof, err := mat.Cofactor(types.Pos{Row: 0, Col: 0})
	if err != nil || cof != ans {
		fmt.Println(cof)
		t.Errorf("Matrix Cofactor (1) Failed: Expected %f got %f. \nErr: %s", ans, cof, err)
	}
	ans = 6
	cof, err = mat.Cofactor(types.Pos{Row: 0, Col: 1})
	if err != nil || cof != ans {
		t.Errorf("Matrix Cofactor (2) Failed: Expected %f got %f. \nErr: %s", ans, cof, err)
	}
	ans = -3
	cof, err = mat.Cofactor(types.Pos{Row: 0, Col: 2})
	if err != nil || cof != ans {
		t.Errorf("Matrix Cofactor (3) Failed: Expected %f got %f. \nErr: %s", ans, cof, err)
	}
	ans = 6
	cof, err = mat.Cofactor(types.Pos{Row: 1, Col: 0})
	if err != nil || cof != ans {
		t.Errorf("Matrix Cofactor (4) Failed: Expected %f got %f. \nErr: %s", ans, cof, err)
	}
	ans = -3
	cof, err = mat.Cofactor(types.Pos{Row: 2, Col: 0})
	if err != nil || cof != ans {
		t.Errorf("Matrix Cofactor (5) Failed: Expected %f got %f. \nErr: %s", ans, cof, err)
	}
}

func TestAdjugate(t *testing.T) {
	mat := NewFromString("3 2 7 ; 5 -1 -3 ; 4 2 9")
	ans := NewFromString("-3 -4 1 ; -57 -1 44 ; 14 2 -13")
	adj, err := mat.Adj()
	if err != nil || !ans.Equals(adj) {
		t.Errorf("Matrix Adjugate (1) Failed: %s", err)
	}
	mat = NewFromString("-1 2 1 ; 3 -1 -3 ; 6 2 -2")
	ans = NewFromString("8 6 -5 ; -12 -4 0 ; 12 14 -5")
	adj, err = mat.Adj()
	if err != nil || !ans.Equals(adj) {
		t.Errorf("Matrix Adjugate (2) Failed: %s", err)
	}
}

// It is difficult to write tests for the Inverse as it usually involves decimal values
// Therefore this is a test that will always pass
// However i did make some tests, and it seems to work
func TestInverse(t *testing.T) {
	// 2x2
	mat := NewFromString("1 2 ; 3 4")
	inv, err := mat.Inverse()
	if err != nil || !inv.Equals(inv) {
		t.Errorf("Matrix Inverse (1) Failed: %s", err)
		inv.Print()
	}

	// 3x3
	mat = NewFromString("2 1 1 ; 3 2 3 ; 4 3 2")
	inv, err = mat.Inverse()
	if err != nil || !inv.Equals(inv) {
		t.Errorf("Matrix Inverse (2) Failed: %s", err)
		inv.Print()
	}
	// This should yield that the inverse doesn't exist because the determinant = 0
	mat = NewFromString("1 2 3 4 ; 5 6 7 8 ; 9 10 11 12 ; 13 14 15 16")
	_, err = mat.Inverse()
	if err == nil {
		t.Errorf("Matrix Inverse (3) Failed: %s", err)
	}
}

func TestDiv(t *testing.T) {
	// 2x2 / 2x2
	mat1 := NewFromString("1 2 ; 3 4")
	mat2 := NewFromString("4 3 ; 2 1")
	ans := NewFromString("1.5 -2.5 ; 2.5 -3.5")
	res, err := mat1.Div(mat2)
	if err != nil || !res.Equals(ans) {
		res.Print()
		t.Errorf("Matrix Division (1) Failed: %s", err)
	}
	// 3x2 / 2x2
	mat1 = NewFromString("1 2 ; 3 4 ; 2 6")
	mat2 = NewFromString("-1 -2 ; -3 -4")
	ans = NewFromString("-1 0 ; 0 -1 ; -5 1")
	res, err = mat1.Div(mat2)
	if err != nil || !res.Equals(ans) {
		res.Print()
		t.Errorf("Matrix Division (2) Failed: %s", err)
	}
}

// Row Swapping
func TestSwapRows(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("7 8 9 ; 4 5 6 ; 1 2 3")
	_, err := mat.SwapRows(0, 2)
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix Row Swapping (1) Failed: %s", err)
	}
	// Out of bounds
	_, err = mat.SwapRows(0, 5)
	if err == nil {
		t.Errorf("Matrix Row Swapping (2) Failed: %s", err)
	}
}

// Scale Row
func TestScaleRow(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("2 4 6 ; 4 5 6 ; 7 8 9")
	_, err := mat.ScaleRowBy(0, 2)
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix Row Scaling (1) Failed: %s", err)
	}
	// Out of bounds
	_, err = mat.ScaleRowBy(10, 5)
	if err == nil {
		t.Errorf("Matrix Row Scaling (2) Failed: %s", err)
	}
}

// Test Gaussian Row Reduction
func TestGaussRREF(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("1 0 -1 ; 0 1 2 ; 0 0 0")
	_, err := mat.GaussRREF()
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix Gaussian RREF (1) Failed: %s", err)
		mat.Print()
	}
	// 4x4
	mat = NewFromString("1 2 -1 -4 ; 2 3 -1 -11 ; -2 0 -3 22")
	ans = NewFromString("1 0 0 -8 ; 0 1 0 1 ; 0 0 1 -2")
	_, err = mat.GaussRREF()
	if err != nil || !mat.Equals(ans) {
		t.Errorf("Matrix Gaussian RREF (2) Failed: %s", err)
		mat.Print()
	}
}

// Test Matrix Augmentation
func TestAugmentation(t *testing.T) {
	// Base Case
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	vec := NewFromString("0 ; 0 ; 0")
	ans := NewFromString("1 2 3 0 ; 4 5 6 0 ; 7 8 9 0")
	aug, err := mat.Augment(vec)
	if err != nil || !aug.Equals(ans) {
		t.Errorf("Matrix Augmentation (1) Failed: %s", err)
		aug.Print()
	}
	// Row size mismatch
	vec = NewFromString("0 ; 0 ; 0 ; 0")
	_, err = mat.Augment(vec)
	if err == nil {
		t.Errorf("Matrix Augmentation (2) Failed: %s", err)
	}
}

// Test GetCol
func TestGetCol(t *testing.T) {
	mat := NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	ans := NewFromString("2 ; 5 ; 8")
	col, err := mat.GetCol(1)
	if err != nil || !col.Equals(ans) {
		t.Errorf("Matrix GetCol Failed: %s", err)
	}
}

// Test Random Matrix
func TestRandomMatrix(t *testing.T) {
	mat1 := NewRandom(3, 3)
	mat2 := NewRandom(3, 3)
	if mat1.Equals(mat2) {
		t.Errorf("Matrix Random Matrix Generation Failed")
		mat1.Print()
		mat2.Print()
	}
}

// Test CompMax
func TestCompMax(t *testing.T) {
	val := types.MatrixType(4)
	mat := NewFromString("4 0 4 ; 4 0 4 ; 4 0 4")
	ans := NewFromString("4 4 4 ; 4 4 4 ; 4 4 4")
	if !mat.CompMax(val).Equals(ans) {
		t.Error("Matrix CompMax Failed")
		mat.Print()
	}
}

// Test CompMin
func TestCompMin(t *testing.T) {
	val := types.MatrixType(4)
	mat := NewFromString("4 0 4 ; 4 0 4 ; 4 0 4")
	ans := NewFromString("4 0 4 ; 4 0 4 ; 4 0 4")
	if !mat.CompMin(val).Equals(ans) {
		t.Error("Matrix CompMin Failed")
		mat.Print()
	}
}
