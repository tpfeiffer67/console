package screenutils

import (
	"github.com/tpfeiffer67/console/screen"
)

// ChromaticFilter
type ChromaticFilter struct {
	*screen.Size
	*screen.FColorCanvas
	*screen.BColorCanvas
	*screen.Stencil
}

func NewChromaticFilter(stencil *screen.Stencil, color screen.Color) *ChromaticFilter {
	o := new(ChromaticFilter)
	height, width := stencil.Height(), stencil.Width()
	o.Size = screen.NewSize(height, width)
	o.Size.SetOnSizeChange(o.onResize)
	o.FColorCanvas = screen.NewFColorCanvasWithDefaultValue(height, width, color)
	o.BColorCanvas = screen.NewBColorCanvasWithDefaultValue(height, width, color)
	o.FColorCanvas.RenderMode = 3
	o.BColorCanvas.RenderMode = 3
	o.Stencil = stencil
	return o
}

func (o *ChromaticFilter) Resize(height int, width int) { o.Size.Resize(height, width) }
func (o *ChromaticFilter) Height() int                  { return o.Size.Height() }
func (o *ChromaticFilter) Width() int                   { return o.Size.Width() }
func (o *ChromaticFilter) SetHeight(height int)         { o.Size.SetHeight(height) }
func (o *ChromaticFilter) SetWidth(width int)           { o.Size.SetWidth(width) }

func (o *ChromaticFilter) onResize(height int, width int) {
	o.FColorCanvas.Resize(height, width)
	o.BColorCanvas.Resize(height, width)
	o.Stencil.Resize(height, width)
}

func (o *ChromaticFilter) Clear() {
	o.FColorCanvas.Clear()
	o.BColorCanvas.Clear()
}

func (o *ChromaticFilter) Render(b *screen.Buffer, stencil *screen.Stencil, position screen.Coordinates, offset screen.Coordinates) {
	if stencil == nil {
		stencil = o.Stencil
	}
	o.FColorCanvas.Render(b, stencil, position, offset)
	o.BColorCanvas.Render(b, stencil, position, offset)
}
