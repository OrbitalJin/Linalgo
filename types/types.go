package types

type Transformer func(float32) float32

type Shape struct {
	Rows int
	Cols int
}

type Pos struct {
	Row int
	Col int
}