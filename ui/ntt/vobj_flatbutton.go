package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type FlatButton struct {
	IEntity
	theme.ITheme
	property.PropertyText
	property.OnSelect
}

func NewFlatButton(id string, row, col int, height, width int, label string, syst ISystem) *FlatButton {
	o := new(FlatButton)
	o.IEntity = NewEntity(id, height, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_BUTTON, theme.STYLE_BUTTON_HOVERED, theme.STYLE_BUTTON_FOCUSED, theme.STYLE_BUTTON_FOCUSEDHOVERED, theme.STYLE_BUTTON_DOWN)
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)
	o.SetText(label)

	//SetDefaultFuncFor_OnFocus_And_OnLostFocus(o)
	o.SetOnFocus(func(foc property.IFocus) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc property.IFocus) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnDraw(func() {
		// TODO to improve
		style := ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_BUTTON, theme.STYLE_BUTTON_HOVERED, theme.STYLE_BUTTON_FOCUSED, theme.STYLE_BUTTON_FOCUSEDHOVERED)
		if o.Selected() {
			style = ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN)
		}
		screenutils.DrawStyledString(0, 0, o.Text(), o, style, theme.ToColor)
	})

	o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
		o.Render(sb, pos)
	})

	unselectGroup := func(group int) {
		f := func(i interface{}) {
			if e, ok := i.(property.Selecter); ok {
				if e.GetSelectGroup() == group {
					e.Unselect()
				}
			}
		}
		syst.CallFuncWithAllEntitiesAsInterface(f)
	}

	o.SetListener(message.MessageIdMouseDown, func(messageParams interface{}) bool {
		switch o.Selected() {
		case true:
			o.Unselect()
		default:
			group := o.GetSelectGroup()
			if group != 0 {
				unselectGroup(group)
			}
			o.Select()
		}

		return true
	})

	return o
}
