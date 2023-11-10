package tcellterm

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
	"github.com/tpfeiffer67/console/ui/message"
)

func StartTcell(messagesChannel chan message.Message) tcell.Screen {
	encoding.Register()

	tcellScreen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := tcellScreen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	tcellScreen.EnableMouse()
	tcellScreen.DisablePaste()

	go func() {
		for {
			tcellEvent := tcellScreen.PollEvent()
			switch eventType := tcellEvent.(type) {
			case *tcell.EventResize:
				width, height := eventType.Size()
				messagesChannel <- message.Message{MessageId: message.MessageIdScreenResize, TimeStamp: time.Now(), Params: message.ParamsScreenResize{Rows: height, Cols: width}}

			case *tcell.EventKey:
				key := int(eventType.Key())
				r := eventType.Rune()
				shift, ctrl, alt := getModifiers(eventType.Modifiers())
				name := getName(eventType)
				messagesChannel <- message.Message{MessageId: message.MessageIdKey, TimeStamp: time.Now(), Params: message.ParamsKey{
					Rune: r, Key: key, Shift: shift, Ctrl: ctrl, Alt: alt, Name: name}}

			case *tcell.EventMouse:
				col, row := eventType.Position()
				buttons := eventType.Buttons()
				b1 := iifButton(buttons, tcell.ButtonPrimary)
				b2 := iifButton(buttons, tcell.ButtonSecondary)
				b3 := iifButton(buttons, tcell.ButtonMiddle)
				shift, ctrl, alt := getModifiers(eventType.Modifiers())
				eventType.Buttons()
				messagesChannel <- message.Message{MessageId: message.MessageIdMouse, TimeStamp: time.Now(), Params: message.ParamsMouse{
					Row: row, Col: col, ButtonPrimary: b1, ButtonSecondary: b2, ButtonMiddle: b3, Shift: shift, Ctrl: ctrl, Alt: alt}}
			}
		}
	}()

	return tcellScreen
}

func iifButton(buttons, mask tcell.ButtonMask) bool {
	if buttons&mask == 0 {
		return false
	}
	return true
}

func getModifiers(mm tcell.ModMask) (shift, ctrl, alt bool) {
	shift = iifModMask(mm, tcell.ModShift)
	ctrl = iifModMask(mm, tcell.ModCtrl)
	alt = iifModMask(mm, tcell.ModAlt)
	return shift, ctrl, alt
}

func getName(ev *tcell.EventKey) string {
	k := ev.Key()
	switch k {
	case tcell.KeyRune:
		return string(ev.Rune())
	default:
		return ev.Name()
	}
}

func iifModMask(modifiers, mask tcell.ModMask) bool {
	if modifiers&mask == 0 {
		return false
	}
	return true
}
