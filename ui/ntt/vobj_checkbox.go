package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Checkbox struct {
	IWidget
	property.PropertyText
	checked      bool
	CheckedBox   string
	UncheckedBox string
	onChange     func(bool)
}

func (o *Checkbox) SetChecked(b bool) {
	o.checked = b
	o.doOnChange()
}

func (o *Checkbox) Checked() bool {
	return o.checked
}

func (o *Checkbox) SetOnChange(f func(bool)) {
	o.onChange = f
}

func (o *Checkbox) doOnChange() {
	if o.onChange != nil {
		o.onChange(o.checked)
	}
}

func NewCheckbox(id string, row, col int, label string, syst ISystem) *Checkbox {
	o := new(Checkbox)
	o.IWidget = NewWidget(id, 1, 0, syst) // The actual width is unknown here. It depends on the design of the box.
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)
	o.SetOnTextChange(o.update)
	o.SetText(label)

	// TODO create function setDefaultGetFocusAndLostFocus()
	o.SetOnFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnThemeChange(func() {
		o.update()
	})

	o.SetOnDraw(func() {
		style := ClearWithStyle(o, o, theme.STYLE_CHECKBOX, theme.STYLE_CHECKBOX_HOVERED, theme.STYLE_CHECKBOX_FOCUSED, theme.STYLE_CHECKBOX_FOCUSEDHOVERED)
		box := o.getBox()
		l := screenutils.DrawStyledString(0, 0, box, o, style, theme.ToColor)
		screenutils.DrawStyledString(0, l, o.Text(), o, style, theme.ToColor)
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		o.checked = !o.checked
		o.doOnChange()
		return true
	})

	return o
}

func (o *Checkbox) getBox() string {
	if o.checked {
		s, _ := o.GetString(theme.CHECKBOX_CHECKED)
		return s
	}
	s, _ := o.GetString(theme.CHECKBOX_NOTCHECKED)
	return s
}

func (o *Checkbox) update() {
	textWidth := screenutils.GetStyledStringLen(o.Text())
	o.CheckedBox, _ = o.GetString(theme.CHECKBOX_CHECKED)
	o.UncheckedBox, _ = o.GetString(theme.CHECKBOX_NOTCHECKED)
	w1 := screenutils.GetStyledStringLen(o.CheckedBox)
	w2 := screenutils.GetStyledStringLen(o.UncheckedBox)
	boxWidth := max(w1, w2)
	o.SetWidth(boxWidth + textWidth)
}
