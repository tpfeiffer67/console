package screenutils

import (
	"fmt"

	"github.com/tpfeiffer67/console/screen"
)

// ComboCanvas is a Combined Canvas Type 1.
type ComboCanvas struct {
	*screen.Size
	//Position
	Origins screen.Coordinates
	*screen.RuneCanvas
	*screen.EmphasisCanvas
	*screen.FColorCanvas
	*screen.BColorCanvas
	*screen.Stencil
}

func NewComboCanvas(height int, width int) *ComboCanvas {
	o := new(ComboCanvas)
	o.Size = screen.NewSize(height, width)
	o.SetOnSizeChange(o.onResize)
	//o.Position = new(Position)
	o.RuneCanvas = screen.NewRuneCanvas(height, width)
	o.EmphasisCanvas = screen.NewEmphasisCanvas(height, width)
	o.FColorCanvas = screen.NewFColorCanvas(height, width)
	o.BColorCanvas = screen.NewBColorCanvas(height, width)
	o.Stencil = screen.NewStencil(height, width, false)
	o.Clear()
	return o
}

func (o *ComboCanvas) Resize(height int, width int) { o.Size.Resize(height, width) }
func (o *ComboCanvas) Height() int                  { return o.Size.Height() }
func (o *ComboCanvas) Width() int                   { return o.Size.Width() }
func (o *ComboCanvas) SetHeight(height int)         { o.Size.SetHeight(height) }
func (o *ComboCanvas) SetWidth(width int)           { o.Size.SetWidth(width) }

func (o *ComboCanvas) onResize(height int, width int) {
	o.RuneCanvas.Resize(height, width)
	o.EmphasisCanvas.Resize(height, width)
	o.FColorCanvas.Resize(height, width)
	o.BColorCanvas.Resize(height, width)
	o.Stencil.Resize(height, width)
}

func (o *ComboCanvas) Clear() {
	o.RuneCanvas.Clear()
	o.EmphasisCanvas.Clear()
	o.FColorCanvas.Clear()
	o.BColorCanvas.Clear()
}

func (o *ComboCanvas) Render(b *screen.Buffer, stencil *screen.Stencil, position screen.Coordinates, offset screen.Coordinates) {
	if stencil == nil {
		stencil = o.Stencil
	}
	o.RuneCanvas.Render(b, stencil, position, offset)
	o.EmphasisCanvas.Render(b, stencil, position, offset)
	o.FColorCanvas.Render(b, stencil, position, offset)
	o.BColorCanvas.Render(b, stencil, position, offset)
}

func (o *ComboCanvas) SetDefaults(r rune, style screen.Style) {
	o.RuneCanvas.DefaultValue = r
	o.EmphasisCanvas.DefaultValue = style.Emphasis
	o.FColorCanvas.DefaultValue = style.FColor
	o.BColorCanvas.DefaultValue = style.BColor
}

/*
func (o *ComboCanvas) GetCanvas() (*screen.RuneCanvas, *screen.FColorCanvas, *screen.BColorCanvas, *screen.EmphasisCanvas) {
	return o.RuneCanvas, o.FColorCanvas, o.BColorCanvas, o.EmphasisCanvas
}
*/
func (o *ComboCanvas) GetRuneCanvas() *screen.RuneCanvas         { return o.RuneCanvas }
func (o *ComboCanvas) GetFColorCanvas() *screen.FColorCanvas     { return o.FColorCanvas }
func (o *ComboCanvas) GetBColorCanvas() *screen.BColorCanvas     { return o.BColorCanvas }
func (o *ComboCanvas) GetEmphasisCanvas() *screen.EmphasisCanvas { return o.EmphasisCanvas }
func (o *ComboCanvas) GetStencil() *screen.Stencil               { return o.Stencil }
func (o *ComboCanvas) GetOrigins() screen.Coordinates            { return o.Origins }

func (o *ComboCanvas) String() string {
	return fmt.Sprintf("ComboCanvas,width=%d,height=%d", o.Width(), o.Height())
}
