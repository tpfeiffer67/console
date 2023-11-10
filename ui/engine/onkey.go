package engine

import "github.com/tpfeiffer67/console/ui/message"

type OnKey struct {
	onKey func(*message.ParamsKey) bool
}

func (o *OnKey) SetOnKey(f func(*message.ParamsKey) bool) {
	o.onKey = f
}

func (o *OnKey) GetOnKey() func(*message.ParamsKey) bool {
	return o.onKey
}

func (o *OnKey) doOnKey(params *message.ParamsKey) bool {
	if o.onKey != nil {
		return o.onKey(params)
	}
	return false
}
