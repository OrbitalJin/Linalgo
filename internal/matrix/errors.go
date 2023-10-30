package matrix

import (
	"fmt"

	"github.com/OrbitalJin/Linalgo/types"
)

func squareError(s types.Shape, reason string) error {
	return fmt.Errorf(
		"incompatible shape (%d, %d), matrix must be square to compute it's determinat",
		s.Rows, s.Cols,
	)
}
