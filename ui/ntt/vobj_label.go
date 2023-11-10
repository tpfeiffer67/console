package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Label struct {
	IEntity
	theme.ITheme
	property.PropertyText
}

func NewLabel(id string, row, col int, label string, syst ISystem) *Label {
	width := screenutils.GetStyledStringLen(label)
	o := new(Label)
	o.IEntity = NewEntity(id, 1, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_LABEL)
	o.SetOnTextChange(o.update)
	o.SetText(label)
	o.SetPosition(row, col)
	o.SetCanMove(false)

	o.SetOnDraw(func() {
		style, _ := o.ITheme.GetStyle(theme.STYLE_LABEL)
		o.SetDefaults(' ', style)
		o.Clear()
		screenutils.DrawStyledString(0, 0, o.Text(), o, style, theme.ToColor)
	})

	return o
}

func (o *Label) update() {
	width := screenutils.GetStyledStringLen(o.Text())
	o.SetWidth(width)
}
