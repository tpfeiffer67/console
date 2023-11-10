/*
Package screen provides a mid-level library.
*/
package screen

import (
	"errors"
	"fmt"
)

// Buffer stores a matrix of Cells to render to the terminal.
type Buffer struct {
	Coordinates
	height       int
	width        int
	current      int
	previous     int
	DefaultValue Cell
	currentMeta  interface{}
	page         [2][]Cell
	matrix       [2][][]Cell
	meta         []interface{}
	metaMatrix   [][]interface{}
}

func (o *Buffer) Height() int { return o.height }
func (o *Buffer) Width() int  { return o.width }

func (o *Buffer) DefineMeta(meta interface{}) {
	o.currentMeta = meta
}

func (o *Buffer) ResetMeta() {
	o.currentMeta = nil
}

func (o *Buffer) setMeta(row, col int) {
	if o.currentMeta != nil {
		o.metaMatrix[row][col] = o.currentMeta
	}
}

func (o *Buffer) PeekMeta(row, col int) interface{} {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		return o.metaMatrix[row][col]
	}
	return nil
}

func (o *Buffer) PokeCell(row, col int, cell Cell) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		o.matrix[o.current][row][col] = cell
		o.setMeta(row, col)
	}
}

func (o *Buffer) PokeRune(row, col int, r rune) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		o.matrix[o.current][row][col].Rune = r
		o.setMeta(row, col)
	}
}

func (o *Buffer) PokeFColor(row, col int, color Color) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		o.matrix[o.current][row][col].FColor = color
		o.setMeta(row, col)
	}
}

func (o *Buffer) PokeBColor(row, col int, color Color) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		o.matrix[o.current][row][col].BColor = color
		o.setMeta(row, col)
	}
}

func (o *Buffer) PokeEmphasis(row, col int, emphasis Attributes) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		o.matrix[o.current][row][col].Emphasis = emphasis
		o.setMeta(row, col)
	}
}

func (o *Buffer) PeekCell(row, col int) (Cell, error) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		return o.matrix[o.current][row][col], nil
	}
	return o.DefaultValue, errors.New("the coordinates provided are outside the limits")
}

func (o *Buffer) PeekRune(row, col int) (rune, error) {
	cell, err := o.PeekCell(row, col)
	return cell.Rune, err
}

func (o *Buffer) PeekFColor(row, col int) (Color, error) {
	cell, err := o.PeekCell(row, col)
	return cell.FColor, err
}

func (o *Buffer) PeekBColor(row, col int) (Color, error) {
	cell, err := o.PeekCell(row, col)
	return cell.BColor, err
}

func (o *Buffer) PeekEmphasis(row, col int) (Attributes, error) {
	cell, err := o.PeekCell(row, col)
	return cell.Emphasis, err
}

func (o *Buffer) PeekCellCurrentAndPreviousValue(row, col int) (Cell, Cell) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.height, o.width) {
		return o.matrix[o.current][row][col], o.matrix[o.previous][row][col]
	}
	return Cell{}, Cell{}
}

func NewBuffer(height, width int) *Buffer {
	o := new(Buffer)
	o.Resize(height, width)
	o.DefaultValue.Rune = ' '
	o.clear()
	o.Swap()
	o.clear()
	o.clearMeta()
	return o
}

func (o *Buffer) Resize(height, width int) {
	if height > 0 && width > 0 && (o.height != height || o.width != width) {
		o.height = height
		o.width = width
		o.resize(0)
		o.resize(1)
		o.resizeMeta()
	}
}

func (o *Buffer) resize(index int) {
	l := o.height * o.width
	o.page[index] = make([]Cell, l)
	o.matrix[index] = make([][]Cell, o.height)
	for i := 0; i < o.height; i++ {
		o.matrix[index][i] = o.page[index][i*o.width : i*o.width+o.width]
	}
}

func (o *Buffer) resizeMeta() {
	l := o.height * o.width
	o.meta = make([]interface{}, l)
	o.metaMatrix = make([][]interface{}, o.height)
	for i := 0; i < o.height; i++ {
		o.metaMatrix[i] = o.meta[i*o.width : i*o.width+o.width]
	}
}

func (o *Buffer) Clear() {
	o.clear()
	o.clearMeta()
}

func (o *Buffer) clear() {
	for i := 0; i < o.height*o.width; i++ {
		o.page[o.current][i] = o.DefaultValue
	}
}

func (o *Buffer) clearMeta() {
	for i := 0; i < o.height*o.width; i++ {
		o.meta[i] = nil
	}
}

func (o *Buffer) Swap() {
	if o.current == 0 {
		o.current = 1
		o.previous = 0
	} else {
		o.current = 0
		o.previous = 1
	}
}

func (o *Buffer) String() string {
	return fmt.Sprintf("Buffer,row=%d,col=%d,width=%d,height=%d", o.Row, o.Col, o.width, o.height)
}
