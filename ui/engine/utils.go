package engine

import (
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/ntt"
)

func (o *Engine) SendMessage(messageId message.MessageId, messageParams interface{}, objId string) {
	o.uiMessageChannel <- message.Message{MessageId: messageId, Params: messageParams, ObjId: objId}
}

func (o *Engine) getMetaUnderMousePosition(mouseParams message.ParamsMouse) interface{} {
	return o.Screen.Buffer.PeekMeta(mouseParams.Row, mouseParams.Col)
}

func (o *Engine) GetPointedObjectId() (string, bool) {
	if o.pointedEntity != nil {
		return o.pointedEntity.Id(), true
	}
	return "", false
}

func (o *Engine) GetPointedObject() (ntt.IEntity, bool) {
	if o.pointedEntity != nil {
		return o.pointedEntity, true
	}
	return nil, false
}
