package property

import (
	"github.com/tpfeiffer67/console/ui/message"
)

// We use the same structure than for the message

type IOnKeyPress interface {
	SetOnKeyPress(func(message.ParamsKey))
	DoOnKeyPress(message.ParamsKey)
}

type OnKeyPress struct {
	onKeyPress func(message.ParamsKey)
}

func (o *OnKeyPress) SetOnKeyPress(f func(message.ParamsKey)) { o.onKeyPress = f }
func (o *OnKeyPress) DoOnKeyPress(m message.ParamsKey) {
	if o.onKeyPress != nil {
		o.onKeyPress(m)
	}
}
