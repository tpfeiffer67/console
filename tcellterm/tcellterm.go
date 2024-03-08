package tcellterm

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/tpfeiffer67/console/ui/message"
)

func StartTcell(messagesChannel chan message.InputMessage) tcell.Screen {
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
				m := message.InputMessageScreenSize{}
				//m.MessageId = message.MessageIdScreenResize
				m.TimeStamp = time.Now()
				m.Width, m.Height = eventType.Size()
				messagesChannel <- m

			case *tcell.EventKey:
				m := message.InputMessageKey{}
				//m.MessageId = message.MessageIdKey
				m.TimeStamp = time.Now()
				m.Rune = eventType.Rune()
				m.Key = int(eventType.Key())
				m.Shift, m.Ctrl, m.Alt = getModifiers(eventType.Modifiers())
				m.Name = getName(eventType)
				messagesChannel <- m

			case *tcell.EventMouse:
				m := message.InputMessageMouse{}
				//m.MessageId = message.MessageIdMouse
				m.TimeStamp = time.Now()
				m.Col, m.Row = eventType.Position()
				m.ButtonPrimary, m.ButtonSecondary, m.ButtonMiddle = mouseButtons(eventType.Buttons())
				m.Shift, m.Ctrl, m.Alt = getModifiers(eventType.Modifiers())
				messagesChannel <- m
			}
		}
	}()

	return tcellScreen
}

func mouseButtons(buttons tcell.ButtonMask) (bool, bool, bool) {
	return iifButton(buttons, tcell.ButtonPrimary), iifButton(buttons, tcell.ButtonSecondary), iifButton(buttons, tcell.ButtonMiddle)
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
