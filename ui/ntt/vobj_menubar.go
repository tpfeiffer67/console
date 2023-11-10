package ntt

import (
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type MenuBar struct {
	IEntity
	theme.ITheme
}

func NewMenuBar(id string, syst ISystem) *MenuBar {
	o := new(MenuBar)
	o.IEntity = NewEntity(id, 1, 50, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_MENUBAR, theme.STYLE_MENUBAR_HOVERED, theme.STYLE_MENUBAR_FOCUSED, theme.STYLE_MENUBAR_FOCUSEDHOVERED, theme.MENUBAR_FIRST_ITEM_POSITION, theme.MENUBAR_SPACE_BETWEEN_ITEMS)

	o.SetZOrderLayer(1000)
	o.SetZOrder(3)
	o.SetCanMove(false)
	o.SetFocusable(true)

	//o.setDefaultGetFocusAndLostFocus()
	o.SetOnFocus(func(foc property.IFocus) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc property.IFocus) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnDraw(func() {
		ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_MENUBAR, theme.STYLE_MENUBAR_HOVERED, theme.STYLE_MENUBAR_FOCUSED, theme.STYLE_MENUBAR_FOCUSEDHOVERED)
		row, _ := o.GetIntDef(theme.MENUBAR_FIRST_ITEM_POSITION, 1)
		sep, _ := o.GetInt(theme.MENUBAR_SPACE_BETWEEN_ITEMS)
		syst.CallFuncForEachChildrenEntity(id, func(v IEntity) {
			v.SetPosition(0, row)
			w := v.Width()
			row = row + w + sep
		})
	})

	// Positioning of the MenuBar at the top of the screen
	o.SetOnScreenResize(func(height, width int) {
		o.Resize(1, width)
		o.SetPosition(0, 0)
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		syst.Focus(o)
		return true
	})

	return o
}
