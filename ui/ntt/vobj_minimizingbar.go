package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type MinimizingBar struct {
	IWidget
	ISystem
}

func (o *MinimizingBar) AttachChild(a any) {
	if e, ok := a.(property.IParent); ok {
		e.SetParent(o.Id())
	}
}

func NewMinimizingBar(id string, syst ISystem) *MinimizingBar {
	o := new(MinimizingBar)
	o.IWidget = NewWidget(id, 1, 1, syst) // TODO Provide the width of the screen on the creation
	o.ISystem = syst
	o.SetZOrderLayer(1000)
	o.SetZOrder(2)
	o.SetCanMove(false)

	o.SetOnDraw(func() {
		style, _ := o.GetStyle(theme.STYLE_MINIMIZINGBAR)
		o.SetDefaults(' ', style)
		o.Clear()

		stickerWidth, _ := o.GetIntDef(theme.MINIMIZINGBAR_STICKER_WIDTH, 10)
		row, col := 0, 0
		syst.CallFuncForEachChildrenEntity(id, func(a any) {
			if col+stickerWidth > o.Width() {
				row++
				col = 0
			}
			if e, ok := a.(screen.SizeSetterAndPositionSetter); ok {
				e.SetPosition(row, col)
				e.Resize(1, stickerWidth)
			}
			col = col + stickerWidth
		})
		height := row + 1
		o.Resize(height, o.Width())
		o.SetPosition(-height, 0)
	})

	// The MinimizedBar bar is positioned at the bottom of the screen above the statusBar
	o.SetOnScreenResize(func(height, width int) {
		o.Resize(o.Height(), width)
		o.SetPosition(-1, 0)
	})

	o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
		// if no Vobjs are minimized, the minimizedBar can be hidden
		_, b := o.GetFirstChildEntity(o.Id())
		o.SetVisible(b)

		if o.Visible() {
			sb.DefineMeta(o)
			o.GetFColorCanvas().Render(sb, o.GetStencil(), pos, screen.Coordinates{})
			sb.ResetMeta()
			o.GetBColorCanvas().Render(sb, o.GetStencil(), pos, screen.Coordinates{})
		}
	})

	return o
}
