package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Pointer struct {
	*screen.RuneCanvas
	*screen.EmphasisCanvas
	*screen.FColorCanvas
	*screen.Stencil
	Visible bool
}

func NewPointer() *Pointer {
	o := new(Pointer)
	// TODO To define in Theme (DrawBlock rune + color)
	o.RuneCanvas = screen.NewRuneCanvas(1, 1)
	o.RuneCanvas.DefaultValue = '▓'
	o.RuneCanvas.Clear()
	o.EmphasisCanvas = screen.NewEmphasisCanvas(1, 1)
	o.EmphasisCanvas.Clear()
	o.FColorCanvas = screen.NewFColorCanvas(1, 1)
	o.FColorCanvas.DefaultValue = theme.ToColor("#FFF")
	o.FColorCanvas.Clear()
	o.Stencil = screen.NewStencil(1, 1, false)
	o.Visible = false // TODO set to visible
	return o
}

func (o *Pointer) RenderAtPosition(sb *screen.Buffer, pos screen.Coordinates) {
	if o.Visible {
		o.RuneCanvas.Render(sb, o.Stencil, pos, screen.Coordinates{})
		o.EmphasisCanvas.Render(sb, o.Stencil, pos, screen.Coordinates{})
		o.FColorCanvas.Render(sb, o.Stencil, pos, screen.Coordinates{})
	}
}
