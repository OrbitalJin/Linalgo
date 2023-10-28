package types

type Transformer func(float32) float32

type Shape struct {
	Rows int
	Cols int
}