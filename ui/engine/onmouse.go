package engine

import (
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/ntt"
)

type OnMouse struct {
	onMouse func(*message.ParamsMouse, message.ParamsMouse, ntt.IEntity) bool
}

func (o *OnMouse) SetOnMouse(f func(*message.ParamsMouse, message.ParamsMouse, ntt.IEntity) bool) {
	o.onMouse = f
}

func (o *OnMouse) GetOnMouse() func(*message.ParamsMouse, message.ParamsMouse, ntt.IEntity) bool {
	return o.onMouse
}

func (o *OnMouse) doOnMouse(params *message.ParamsMouse, memoParams message.ParamsMouse, v ntt.IEntity) bool {
	if o.onMouse != nil {
		return o.onMouse(params, memoParams, v)
	}
	return false
}
