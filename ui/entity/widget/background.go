package widget

import (
	"github.com/tpfeiffer67/console/ui/theme"
)

type Background struct {
	IWidget
	Background rune
}

func NewBackground(id string, height, width int, syst ISystem) *Background {
	o := new(Background)
	o.IWidget = NewWidget(id, height, width, syst)
	o.SetZOrderLayer(-1000)
	o.SetZOrder(1)
	o.SetCanHide(false)
	o.SetCanMove(false)

	o.SetOnThemeChange(func() {
		o.update()
	})

	o.SetOnDraw(func() {
		style, _ := o.GetStyle(theme.STYLE_BACKGROUND)
		o.SetDefaults(o.Background, style)
		o.Clear()
	})
	o.DoOnDraw()

	o.SetOnScreenResize(func(height, width int) {
		o.Resize(height, width)
	})
	return o
}

func (o *Background) update() {
	o.Background, _ = o.GetChar(theme.BACKGROUND_PATTERN)
}
