package widget

import (
	"github.com/tpfeiffer67/console/screen"
)

type ColorBox struct {
	IWidget
}

func NewColorBox(id string, row, col, height, width int, syst ISystem) *ColorBox {
	o := new(ColorBox)
	o.IWidget = NewWidget(id, height, width, syst)
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(false)

	o.SetColor(screen.Color{})

	return o
}

func (o *ColorBox) SetColor(c screen.Color) {
	style := screen.Style{}
	style.BColor = c
	o.SetDefaults(' ', style)
	o.Clear()
}
