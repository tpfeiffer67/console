package widget

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type FlatButton struct {
	IWidget
	property.PropertyText
	property.OnSelect
}

func NewFlatButton(id string, row, col int, height, width int, label string, syst ISystem) *FlatButton {
	o := new(FlatButton)
	o.IWidget = NewWidget(id, height, width, syst)
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)
	o.SetText(label)

	//SetDefaultFuncFor_OnFocus_And_OnLostFocus(o)
	o.SetOnFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnDraw(func() {
		// TODO to improve
		style := ClearWithStyle(o, o.IWidget, theme.STYLE_BUTTON, theme.STYLE_BUTTON_HOVERED, theme.STYLE_BUTTON_FOCUSED, theme.STYLE_BUTTON_FOCUSEDHOVERED)
		if o.Selected() {
			style = ClearWithStyle(o, o.IWidget, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN, theme.STYLE_BUTTON_DOWN)
		}
		screenutils.DrawStyledString(0, 0, o.Text(), o, style, theme.ToColor)
	})

	o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
		o.Render(sb, pos)
	})

	unselectGroup := func(group int) {
		f := func(i any) {
			if e, ok := i.(property.Selecter); ok {
				if e.GetSelectGroup() == group {
					e.Unselect()
				}
			}
		}
		syst.CallFuncWithAllEntities(f)
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
