package engine

import (
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
)

func (o *Engine) SendMessage(messageId message.MessageId, messageParams interface{}, objId string) {
	o.uiMessageChannel <- message.Message{MessageId: messageId, Params: messageParams, ObjId: objId}
}

func (o *Engine) getMetaUnderMousePosition(mouseParams message.ParamsMouse) interface{} {
	return o.Screen.Buffer.PeekMeta(mouseParams.Row, mouseParams.Col)
}

func (o *Engine) GetPointedObjectId() (string, bool) {
	if o.pointedEntity != nil {
		return o.pointedEntity.(property.IId).Id(), true // #directId
	}
	return "", false
}

func (o *Engine) GetPointedObject() (any, bool) {
	if o.pointedEntity != nil {
		return o.pointedEntity, true
	}
	return nil, false
}
