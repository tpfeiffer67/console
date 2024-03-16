package screen

import (
	"fmt"
)

type Coordinates struct {
	Row int
	Col int
}

type IPosition interface {
	GetPosition() Coordinates
	Row() int
	Col() int
	SetRow(row int)
	SetCol(col int)
	SetPosition(row, col int)
	MoveRow(rows int)
	MoveCol(cols int)
	Move(rows, cols int)
	CanMove() bool
	SetCanMove(v bool)
}

type Position struct {
	coordinates Coordinates
	canMove     bool
}

func (o *Position) GetPosition() Coordinates {
	return o.coordinates
}

func (o *Position) Row() int {
	return o.coordinates.Row
}

func (o *Position) Col() int {
	return o.coordinates.Col
}

func (o *Position) SetRow(row int) {
	o.coordinates.Row = row
}

func (o *Position) SetCol(col int) {
	o.coordinates.Col = col
}

func (o *Position) SetPosition(row, col int) {
	o.SetRow(row)
	o.SetCol(col)
}

func (o *Position) MoveRow(rows int) {
	o.coordinates.Row = o.coordinates.Row + rows
}

func (o *Position) MoveCol(cols int) {
	o.coordinates.Col = o.coordinates.Col + cols
}

func (o *Position) Move(rows, cols int) {
	o.MoveRow(rows)
	o.MoveCol(cols)
}

func (o *Position) CanMove() bool     { return o.canMove }
func (o *Position) SetCanMove(v bool) { o.canMove = v }

func (o Coordinates) String() string {
	return fmt.Sprintf("Row: %d, Col %d", o.Row, o.Col)
}
