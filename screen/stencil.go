package screen

import (
	"fmt"
	"strconv"
)

type Stencil struct {
	Size
	//	TransparentRune rune
	// false, the default value, means that the character at this position is rendered
	withMask    bool
	maskEnabled bool
	maskBuffer  []bool
	mask        [][]bool
}

func NewStencil(height, width int, withMask bool) *Stencil {
	var o Stencil
	o.withMask = withMask
	if withMask {
		o.maskEnabled = true
	}
	o.Resize(height, width)
	return &o
}

func (o *Stencil) SetWidth(width int) {
	if o.width != width {
		o.Resize(o.height, width)
	}
}

func (o *Stencil) SetHeight(height int) {
	if o.height != height {
		o.Resize(height, o.width)
	}
}

func (o *Stencil) Resize(height, width int) {
	o.height = height
	o.width = width
	if o.withMask {
		l := height * width
		o.maskBuffer = make([]bool, l)
		o.mask = make([][]bool, height)
		for i := 0; i < height; i++ {
			o.mask[i] = o.maskBuffer[i*width : i*width+width]
		}
	}
}

func (o *Stencil) EnableMask()      { o.maskEnabled = true }
func (o *Stencil) DisableMask()     { o.maskEnabled = false }
func (o Stencil) MaskEnabled() bool { return o.maskEnabled }

func (o *Stencil) CellIsNotMasked(row, col int) bool {
	if o.withMask && o.maskEnabled {
		return !o.mask[row][col]
	}
	return true
}

func (o *Stencil) GetMask() [][]bool {
	return o.mask
}

func (o *Stencil) InvertMask() {
	if o.withMask {
		for i := 0; i < o.width*o.height; i++ {
			o.maskBuffer[i] = !o.maskBuffer[i]
		}
	}
}

func (o *Stencil) CutOutAShape(row int, col int, shape []string, transparentRune rune) {
	if o.withMask {
		for _, line := range shape {
			column := col
			for _, c := range line {
				if c != transparentRune {
					if row >= 0 && row < o.height && column >= 0 && column < o.width {
						o.mask[row][column] = true
						fmt.Println(row, column)
					}
				}
				column = column + 1
			}
			row = row + 1
		}
	}
}

func (o *Stencil) String() string {
	return fmt.Sprintf("Stencil,width=%d,height=%d,maskEnable=%s", o.width, o.height, strconv.FormatBool(o.maskEnabled))
}
