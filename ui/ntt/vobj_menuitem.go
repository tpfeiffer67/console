package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type MenuItem struct {
	IEntity
	theme.ITheme
	property.PropertyText
	Action string
}

func NewMenuItem(id string, row, col int, text string, syst ISystem) *MenuItem {
	width := screenutils.GetStyledStringLen(text)
	o := new(MenuItem)
	o.IEntity = NewEntity(id, 1, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_MENUITEM, theme.STYLE_MENUITEM_HOVERED, theme.STYLE_MENUITEM_FOCUSED, theme.STYLE_MENUITEM_FOCUSEDHOVERED)

	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)

	o.SetOnTextChange(o.updateText)
	o.SetText(text)

	o.SetOnFocus(func(foc property.IFocus) {
		if !o.FocusedGroup() {
			syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
		}
		if o.FocusedGroup() {
			l := syst.GetListOfBrotherhood(o.Id())
			for _, bro := range l {
				if child, b := syst.GetFirstChildEntity(bro.Id()); b {
					child.SetVisible(false)
				}
			}

			if child, b := syst.GetFirstChildEntity(o.Id()); b {
				child.SetVisible(true)
			}
		}
	})

	o.SetOnLostFocus(func(foc property.IFocus) {
		if !foc.FocusedGroup() {
			syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
		}
	})

	o.SetOnDraw(func() {
		style := ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_MENUITEM, theme.STYLE_MENUITEM_HOVERED, theme.STYLE_MENUITEM_FOCUSED, theme.STYLE_MENUITEM_FOCUSEDHOVERED)
		screenutils.DrawStyledString(0, 0, o.Text(), o, style, theme.ToColor)
	})

	o.SetListener(message.MessageIdMouseEnter, func(messageParams interface{}) bool {
		if o.FocusedGroup() {
			syst.Focus(o)
		}
		return true
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		if o.Action != "" {
			syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
			topMenu := syst.GetTheTopMostAncestorEntity(o.Id())
			syst.Focus(topMenu)
			// ici execution de l'action
			//syst.ExecuteAction(menuItem.Action)
			return true
		}
		syst.Focus(o)
		return true
	})

	return o
}

func (o *MenuItem) updateText() {
	width := screenutils.GetStyledStringLen(o.Text())
	o.SetWidth(width)
}
