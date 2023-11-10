package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/theme"
)

type MenuDropDown struct {
	IEntityWithTheme
}

func NewMenuDropDown(id string, row, col int, height, width int, syst ISystem) *MenuDropDown {
	o := new(MenuDropDown)
	o.IEntityWithTheme = NewPanelWithShadow(id, row, col, height, width, syst)
	o.AppendValues(theme.STYLE_MENUDROPDOWN, theme.STYLE_MENUDROPDOWN_HOVERED, theme.STYLE_MENUDROPDOWN_FOCUSED, theme.STYLE_MENUDROPDOWN_FOCUSEDHOVERED, theme.MENUDROPDOWN_FRAME)

	o.SetVisible(false)
	o.SetCanMove(false)
	o.SetFocusable(true)

	o.SetOnDraw(func() {
		ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_MENUDROPDOWN, theme.STYLE_MENUDROPDOWN_HOVERED, theme.STYLE_MENUDROPDOWN_FOCUSED, theme.STYLE_MENUDROPDOWN_FOCUSEDHOVERED)
		frame, _ := o.GetInt(theme.MENUDROPDOWN_FRAME)
		screenutils.DrawFrame(o.GetRuneCanvas(), 0, 0, o.Height(), o.Width(), frame)
	})

	o.SetOnLostFocusGroup(func() {
		o.SetVisible(false)
	})

	o.SetOnShow(func() {
		maxWidth := 0
		row := 1
		syst.CallFuncForEachChildrenEntity(id, func(v IEntity) {
			maxWidth = max(maxWidth, v.Width())
			v.SetPosition(row, 1)
			row++
		})
		syst.CallFuncForEachChildrenEntity(id, func(v IEntity) {
			v.SetWidth(maxWidth)
		})
		// TODO Suppress magic numbers 2 and 1. Depends if frame or not.
		o.SetWidth(maxWidth + 2)
		o.SetHeight(row + 1)
	})
	return o
}
