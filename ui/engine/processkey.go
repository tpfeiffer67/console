package engine

import (
	"github.com/tpfeiffer67/console/ui/message"
)

const F10 = 288
const TAB = 9
const ESCAPE = 27
const SHIFT_TAB = 278
const ARROW_UP = 257
const ARROW_DOWN = 258
const ARROW_RIGHT = 259
const ARROW_LEFT = 260

func (o *Engine) processKey(keyParams message.ParamsKey) {
	o.renderingMutex.Lock()
	defer o.renderingMutex.Unlock()

	// It is possible to modify the params here.
	// If the function returns true, the rest of the processing is not done.
	if o.doOnKey(&keyParams) {
		return
	}

	o.DoOnHotKey(keyParams)

	// The priority goes to the focused entity.
	if a, ok := o.GetFocusedEntity(); ok {
		if e, ok := a.(message.IMessageListener); ok {
			if e.ProcessMessage(message.MessageIdKey, keyParams) {
				return
			}
		}
	}

	switch keyParams.Key {
	case F10:
		if keyParams.Shift && keyParams.Ctrl && keyParams.Alt {
			o.EditMode = !o.EditMode
			o.doOnModeChange(o.EditMode)
		}

	case TAB:
		o.SendMessage(message.MessageIdFocusNext, nil, "")
	case SHIFT_TAB:
		o.SendMessage(message.MessageIdFocusPrevious, nil, "")
	case 266: // Page Up
		o.SendMessage(message.MessageIdFocusParent, nil, "")
	case 267: // Page Down
		o.SendMessage(message.MessageIdFocusChild, nil, "")

	// TODO Add management of CanMove
	case ARROW_UP:
		if id, ok := o.GetFocusedEntityId(); ok {
			o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: -1, Cols: 0}, id)
		}
	case ARROW_DOWN:
		if id, ok := o.GetFocusedEntityId(); ok {
			o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: 1, Cols: 0}, id)
		}
	case ARROW_LEFT:
		if id, ok := o.GetFocusedEntityId(); ok {
			o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: 0, Cols: -1}, id)
		}
	case ARROW_RIGHT:
		if id, ok := o.GetFocusedEntityId(); ok {
			o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: 0, Cols: 1}, id)
		}
	case ESCAPE:
		o.SendMessage(message.MessageIdExit, nil, "")

	case 1:
		o.SendMessage(message.MessageIdSetTheme, message.ParamsSetTheme{Name: "default"}, "")
	case 2:
		o.SendMessage(message.MessageIdSetTheme, message.ParamsSetTheme{Name: "terminal"}, "")
	case 4:
		o.SendMessage(message.MessageIdFocus, nil, "menuitem22")
	}
}
