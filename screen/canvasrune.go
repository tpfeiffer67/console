package screen

import (
	"fmt"
)

type RuneCanvas struct {
	Size
	DefaultValue    rune
	TransparentRune rune
	buffer          []rune
	Canvas          [][]rune
}

func NewRuneCanvas(height int, width int) *RuneCanvas {
	o := new(RuneCanvas)
	o.DefaultValue = ' '
	o.Resize(height, width)
	o.Clear()
	return o
}

func (o *RuneCanvas) SetHeight(height int) {
	if o.height != height {
		o.Resize(height, o.width)
	}
}

func (o *RuneCanvas) SetWidth(width int) {
	if o.width != width {
		o.Resize(o.height, width)
	}
}

func (o *RuneCanvas) Resize(height int, width int) {
	// TODO Check input values (not <0 and not > upper reasonable limit)
	o.height = height
	o.width = width
	l := height * width
	o.buffer = make([]rune, l)
	o.Canvas = make([][]rune, height)
	for i := 0; i < height; i++ {
		o.Canvas[i] = o.buffer[i*width : i*width+width]
	}
}

func (o *RuneCanvas) Clear() {
	o.Fill(o.DefaultValue)
}

func (o *RuneCanvas) Fill(value rune) {
	for i := 0; i < o.width*o.height; i++ {
		o.buffer[i] = value
	}
}

func (o *RuneCanvas) Render(b *Buffer, stencil *Stencil, position Coordinates, offset Coordinates) {
	for i := 0; i < stencil.Height(); i++ {
		for j := 0; j < stencil.Width(); j++ {
			if stencil.IsNotMasked(i, j) {
				r1 := position.Row + i
				c1 := position.Col + j
				r2 := offset.Row + i
				c2 := offset.Col + j
				if r1 >= 0 && r1 < b.height && c1 >= 0 && c1 < b.width {
					var r rune
					if r2 >= 0 && r2 < o.height && c2 >= 0 && c2 < o.width {
						r = o.Canvas[r2][c2]
					} else {
						r = o.DefaultValue
					}
					if r != o.TransparentRune {
						b.PokeRune(r1, c1, r)
					}
				}
			}
		}
	}
}

func (o *RuneCanvas) SetRune(row int, col int, ru rune) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		o.Canvas[row][col] = ru
	}
}

func (o *RuneCanvas) GetRune(row int, col int) (rune, bool) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		return o.Canvas[row][col], true
	}
	return rune(0), false
}

func (o *RuneCanvas) String() string {
	return fmt.Sprintf("RuneCanvas,width=%d,height=%d", o.width, o.height)
}
