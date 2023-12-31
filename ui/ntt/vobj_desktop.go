package ntt

import (
	"github.com/tpfeiffer67/console/ui/theme"
)

type Desktop struct {
	IEntity
	theme.ITheme
	Background rune
}

// TODO add hovered style

func NewDesktop(id string, height, width int, syst ISystem) *Desktop {
	o := new(Desktop)
	o.IEntity = NewEntity(id, height, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_DESKTOP, theme.DESKTOP_BACKGROUND)
	o.SetZOrderLayer(-1000)
	o.SetZOrder(1)
	o.SetCanHide(false)
	o.SetCanMove(false)

	o.SetOnThemeChange(func() {
		o.update()
	})

	o.SetOnDraw(func() {
		style, _ := o.GetStyle(theme.STYLE_DESKTOP)
		o.SetDefaults(o.Background, style)
		o.Clear()
	})
	o.DoOnDraw()

	o.SetOnScreenResize(func(height, width int) {
		o.Resize(height, width)
	})
	return o
}

func (o *Desktop) update() {
	o.Background, _ = o.GetChar(theme.DESKTOP_BACKGROUND)
}
