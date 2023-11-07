package system

import (
	"fmt"

	"github.com/OrbitalJin/Linalgo/linalgo/matrix"
	t "github.com/OrbitalJin/Linalgo/types"
)

type System struct {
	Equations matrix.Matrix
	Results   matrix.Matrix
}

func New(eqs, res matrix.Matrix) (*System, error) {
	if eqs.Rows != res.Rows {
		return nil, fmt.Errorf("system: row size mismatch, r1=%d, r2=%d", eqs.Rows, res.Rows)
	}
	return &System{
		Equations: eqs,
		Results:   res,
	}, nil
}

func (s *System) Solve() (*matrix.Matrix, error) {
	if s.Equations.Rows != s.Results.Rows {
		return nil, fmt.Errorf(
			"system: row size mismatch, r1=%d, r2=%d",
			s.Equations.Rows,
			s.Results.Rows,
		)
	}

	aug, err := s.Equations.Augment(&s.Results)
	if err != nil {
		return nil, fmt.Errorf("system solving exception: %s", err)
	}

	res := s.gaussJordan(aug)
	sol, err := res.GetCol(-1)
	if err != nil {
		return nil, fmt.Errorf("system solving exception: %s", err)
	}
	return sol, nil
}

// Gauss Jordan Elimination
func (s *System) gaussJordan(augmented *matrix.Matrix) *matrix.Matrix {
	lead := 0
	for r := 0; r < augmented.Rows; r++ {
		if lead >= augmented.Cols {
			return augmented
		}
		i := r
		for {
			if i == augmented.Rows {
				i = r
				lead++
				if lead == augmented.Cols {
					return augmented
				}
			}
			if augmented.Get(t.Pos{Row: i, Col: lead}) != 0 {
				break
			}
			i++
		}
		augmented.SwapRows(i, r)
		f := augmented.Get(t.Pos{Row: r, Col: lead})
		augmented.ScaleRowBy(r, 1/f)
		for i = 0; i < augmented.Rows; i++ {
			if i != r {
				f = augmented.Get(t.Pos{Row: i, Col: lead})
				for j := 0; j < augmented.Cols; j++ {
					val := augmented.Get(t.Pos{Row: r, Col: j})
					augmented.Set(
						t.Pos{Row: i, Col: j},
						augmented.Get(t.Pos{Row: i, Col: j}) - val * f)
				}
			}
		}
		lead++
	}

	// Normalize pivot elements to 1
	for r := 0; r < augmented.Rows; r++ {
		leadCol := -1
		for c := 0; c < augmented.Cols; c++ {
			if augmented.Get(t.Pos{Row: r, Col: c}) != 0 {
				leadCol = c
				break
			}
		}
		if leadCol != -1 {
			f := augmented.Get(t.Pos{Row: r, Col: leadCol})
			augmented.ScaleRowBy(r, 1/f)
		}
	}

	return augmented
}