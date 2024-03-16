package ntt

import (
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/theme"
)

const MAX_TRACKBAR_WIDTH = 1000

type TrackBar struct {
	IWidget
	max      int
	position int
	onChange func(int)
}

func NewTrackBar(id string, row, col int, max int, syst ISystem) *TrackBar {
	if max > MAX_TRACKBAR_WIDTH {
		max = MAX_TRACKBAR_WIDTH
	}
	width := max + 1
	o := new(TrackBar)
	o.max = max
	o.IWidget = NewWidget(id, 1, width, syst)
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)

	//SetDefaultFuncFor_OnFocus_And_OnLostFocus(o)
	o.SetOnFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnDraw(func() {
		background, _ := o.GetChar(theme.TRACKBAR_BACKGROUND)
		cursor, _ := o.GetChar(theme.TRACKBAR_CURSOR)
		style := ClearWithStyle(o, o.IWidget, theme.STYLE_TRACKBAR, theme.STYLE_TRACKBAR_HOVERED, theme.STYLE_TRACKBAR_FOCUSED, theme.STYLE_TRACKBAR_FOCUSEDHOVERED)
		screenutils.DrawHorizontalLine(o.GetRuneCanvas(), o.GetFColorCanvas(), o.GetBColorCanvas(), 0, 0, width, background, style.FColor, style.BColor)
		o.GetRuneCanvas().SetRune(0, o.position, cursor)
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		o.SendMessage(message.MessageIdFocus, nil, id)
		return true
	})

	o.SetListener(message.MessageIdKey, func(messageParams interface{}) bool {
		if params, ok := messageParams.(message.ParamsKey); ok {
			switch params.Name {
			case "+", "Right":
				o.incPosition()
			case "-", "Left":
				o.decPosition()
			}
			return true
		}
		return false
	})

	return o
}

func (o *TrackBar) SetCursorPosition(p int) {
	if p >= 0 && p <= o.max {
		o.position = p
		o.doOnChange()
	}
}

func (o *TrackBar) CursorPosition() int {
	return o.position
}

func (o *TrackBar) SetOnChange(f func(int)) {
	o.onChange = f
}

func (o *TrackBar) doOnChange() {
	if o.onChange != nil {
		o.onChange(o.position)
	}
}

func (o *TrackBar) decPosition() {
	p := o.position
	p--
	o.SetCursorPosition(p)
}

func (o *TrackBar) incPosition() {
	p := o.position
	p++
	o.SetCursorPosition(p)
}
