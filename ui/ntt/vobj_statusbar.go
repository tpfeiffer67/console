package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/theme"
)

type StatusBar struct {
	IEntity
	theme.ITheme
}

func NewStatusBar(id string, syst ISystem) *TrackBar {
	o := new(TrackBar)
	o.IEntity = NewEntity(id, 1, 50, syst) // TODO Provide the width of the screen on the create
	o.ITheme = theme.NewTheme(theme.STYLE_STATUSBAR, theme.STATUSBAR_ITEMS_SEPARATOR)
	o.SetZOrderLayer(1000)
	o.SetZOrder(1)
	o.SetCanMove(false)

	o.SetOnDraw(func() {
		style, _ := o.GetStyle(theme.STYLE_STATUSBAR)
		o.SetDefaults(' ', style)
		o.Clear()

		// The children of the StatusBar will be rendered after the StatusBar.
		// So, here we will just render the background and the vertical separation bars of the StatusBar.
		// For that, we need the widths of each children Entity.
		// In the same pass, we also set the new column position for the Entity and calculate the position for the next Entity.

		sep, _ := theme.CurrentTheme.GetString(theme.STATUSBAR_ITEMS_SEPARATOR)
		sepWidth := screenutils.GetStyledStringLen(sep)
		col := 0
		syst.CallFuncForEachChildrenEntity(id, func(v IEntity) {
			v.SetPosition(0, col)
			w := v.Width()
			col = col + w
			screenutils.DrawStyledString(0, col, sep, o, style, theme.ToColor)
			col = col + sepWidth
		})
	})

	// The StatusBar bar is positioned at the bottom of the screen.
	o.SetOnScreenResize(func(height, width int) {
		o.Resize(1, width)
		o.SetPosition(height-1, 0)
	})

	return o
}
