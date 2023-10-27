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
}
