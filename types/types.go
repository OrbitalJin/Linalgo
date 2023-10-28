package types

type MatrixType float32

type Transformer func(MatrixType) MatrixType

type Shape struct {
	Rows int
	Cols int
}

type Pos struct {
	Row int
	Col int
}