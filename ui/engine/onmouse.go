package engine

import (
	"github.com/tpfeiffer67/console/ui/message"
)

type OnMouse struct {
	onMouse func(*message.ParamsMouse, message.ParamsMouse, any) bool
}

func (o *OnMouse) SetOnMouse(f func(*message.ParamsMouse, message.ParamsMouse, any) bool) {
	o.onMouse = f
}

func (o *OnMouse) GetOnMouse() func(*message.ParamsMouse, message.ParamsMouse, any) bool {
	return o.onMouse
}

func (o *OnMouse) doOnMouse(params *message.ParamsMouse, memoParams message.ParamsMouse, v any) bool {
	if o.onMouse != nil {
		return o.onMouse(params, memoParams, v)
	}
	return false
}
