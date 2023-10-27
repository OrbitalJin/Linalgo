package main

import "github.com/OrbitalJin/Linalgo/internal/matrix"

func main() {
  mat1 := matrix.NewFromString("1 2 3 ; 4 5 6")
	mat2 := matrix.NewFromString("-1 -2 ; -3 -4 ; -5 -6")
	mat3, err := mat1.Dot(mat2)
  if err != nil {
    panic(err)
  }
  mat3.Print()
	I3, err := matrix.NewIdentity(3);
	mat := matrix.NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
	res, err := mat.Dot(I3);
  if err != nil {
    panic(err)
  }
  I3.Print()
  res.Print()
}
