package system

import (
	"testing"

	"github.com/OrbitalJin/Linalgo/internal/matrix"
)

// Test Gaussian Row Reduction
func TestGaussJordanRREF(t *testing.T) {
	// Test 1
	mat := matrix.NewFromString("1 2 3 ; 4 5 6 ; 2 1 4")
	res := matrix.NewFromString("-1 ; 2 ; -1")
	// Expected solution vector
	ans := matrix.NewFromString("1.75 ; 0.5 ; -1.25") 
	// System Init
	sys, err := New(*mat, *res)
	if err != nil {
		t.Errorf("System Initialization (1) Failed: %s", err)
	}
	// Solve
	sol, err := sys.Solve()
	if err != nil || !sol.Equals(ans) {
		t.Errorf("System Solving (1) Failed: %s", err)
		sol.Print()
	}

	// Test 2
	mat = matrix.NewFromString("1 2 1 ; 2 3 3 ; 1 2 2")
	res = matrix.NewFromString("3 ; 6 ; 4")
	// Expected solution vector
	ans = matrix.NewFromString("0 ; 1 ; 1")
	// System Init
	sys, err = New(*mat, *res)
	if err != nil {
		t.Errorf("System Initialization (2) Failed: %s", err)
	}
	// Solve
	sol, err = sys.Solve()
	if err != nil || !sol.Equals(ans) {
		t.Errorf("System Solving (2) Failed: %s", err)
		sol.Print()
	}
}
