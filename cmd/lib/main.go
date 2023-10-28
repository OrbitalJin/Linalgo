package main

import (
	"github.com/OrbitalJin/Linalgo/internal/matrix"
	"github.com/OrbitalJin/Linalgo/types"
)

func main() {
  mat := matrix.NewFromString("1 2 3 ; 4 5 6 ; 7 8 9")
  sub, err := mat.SubMatrix(
    types.Pos{Row: 0, Col: 1},
    types.Pos{Row: 2, Col: 2},
  )
  if err != nil {
    panic(err)
  }
  mat.Print()
  sub.Print()
}
