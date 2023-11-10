package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Button struct {
	IEntity
	theme.ITheme
	property.PropertyText
	down bool
}

func NewButton(id string, row, col int, height, width int, label string, syst ISystem) *Button {
	o := new(Button)
	o.IEntity = NewEntity(id, height, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_BUTTON, theme.STYLE_BUTTON_HOVERED, theme.STYLE_BUTTON_FOCUSED, theme.STYLE_BUTTON_FOCUSEDHOVERED, theme.BUTTON_FRAME, theme.SHADOW_ENABLED, theme.SHADOW_VALUE, theme.SHADOW_VERTICAL_OFFSET, theme.SHADOW_HORIZONTAL_OFFSET)
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
		style := ForEntity_GetStyleByItsStatus_AndClear(o, o, theme.STYLE_BUTTON, theme.STYLE_BUTTON_HOVERED, theme.STYLE_BUTTON_FOCUSED, theme.STYLE_BUTTON_FOCUSEDHOVERED)
		frame, _ := o.GetInt(theme.BUTTON_FRAME)
		screenutils.DrawFrame(o.GetRuneCanvas(), 0, 0, o.Height(), o.Width(), frame)
		screenutils.DrawStyledString(1, 1, o.Text(), o, style, theme.ToColor)
	})

	o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
		// TODO test Visible
		if !o.down {
			DrawShadowAccordingToTheTheme(sb, o.GetStencil(), pos, o)
			o.Render(sb, pos)
			return
		}
		verticalOffset, _ := o.GetInt(theme.SHADOW_VERTICAL_OFFSET)
		horizontalOffset, _ := o.GetInt(theme.SHADOW_HORIZONTAL_OFFSET)
		pos.Row = pos.Row + verticalOffset
		pos.Col = pos.Col + horizontalOffset
		o.Render(sb, pos)
	})

	o.SetListener(message.MessageIdMouseDown, func(messageParams interface{}) bool {
		o.down = true
		return true
	})

	o.SetListener(message.MessageIdMouseUp, func(messageParams interface{}) bool {
		o.down = false
		return true
	})

	return o
}
