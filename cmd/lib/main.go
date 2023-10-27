package main

import (
	"github.com/OrbitalJin/Linalgo/internal/matrix"
)

func main() {
  m1 := matrix.NewFromString("-1 -2 -3 ; -4 -5 -6 ; -7 -8 -9")
  m2 := matrix.NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
  
  m1.Print();
  m2.Print();
  m1.Sub(m2);
  m1.Print();
  m1.T()
  m1.Print();
}
