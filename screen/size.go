package screen

import (
	"fmt"
)

type Size struct {
	height       int
	width        int
	onSizeChange func(int, int)
}

func NewSize(height, width int) *Size {
	o := Size{height, width, nil}
	return &o
}

func (o *Size) Height() int { return o.height }
func (o *Size) Width() int  { return o.width }

func (o *Size) SetHeight(height int) {
	if height > 0 && o.height != height {
		o.resize(height, o.width)
	}
}

func (o *Size) SetWidth(width int) {
	if width > 0 && o.width != width {
		o.resize(o.height, width)
	}
}

// TODO Add error handling
func (o *Size) Resize(height, width int) {
	if height < 1 || width < 1 {
		return
	}
	if o.height != height || o.width != width {
		o.resize(height, width)
	}
}

func (o *Size) resize(height, width int) {
	o.height = height
	o.width = width
	o.DoOnSizeChange(height, width)
}

func (o *Size) SetOnSizeChange(f func(int, int)) { o.onSizeChange = f }
func (o *Size) GetOnSizeChange() func(int, int)  { return o.onSizeChange }
func (o *Size) DoOnSizeChange(height, width int) {
	if o.onSizeChange != nil {
		o.onSizeChange(height, width)
	}
}

func (o *Size) String() string {
	return fmt.Sprintf("screen.Size{height=%d,width=%d}", o.height, o.width)
}
