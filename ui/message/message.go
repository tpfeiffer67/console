package message

import (
	"time"
)

type Message struct {
	MessageId MessageId
	Params    interface{}
	ObjId     string
	TimeStamp time.Time
}

type ParamsScreenResize struct {
	Rows int
	Cols int
}

// Faire un SetSize
type ParamsResize struct {
	Rows int
	Cols int
}

type ParamsMove struct {
	Rows int
	Cols int
}

type ParamsSetPosition struct {
	Row int
	Col int
}

type ParamsKey struct {
	Rune  rune
	Key   int
	Shift bool
	Ctrl  bool
	Alt   bool
	Name  string
}

type ParamsSetTheme struct {
	Name string
}

type ParamsMouse struct {
	Row             int
	Col             int
	ButtonPrimary   bool
	ButtonSecondary bool
	ButtonMiddle    bool
	Shift           bool
	Ctrl            bool
	Alt             bool
}

type ParamsMouseExt struct {
	ParamsMouse
	DeltaRow int
	DeltaCol int
}

type FuncProcessMessage = func(interface{}) bool

type IMessageListener interface {
	SetListener(MessageId, func(interface{}) bool)
	GetListener(MessageId) (func(interface{}) bool, bool)
	ProcessMessage(MessageId, interface{}) bool
}

type MessageListener struct {
	listener map[MessageId]func(interface{}) bool
}

func NewMessageListener() MessageListener {
	var o MessageListener
	o.listener = make(map[MessageId]func(interface{}) bool)
	return o
}

func (o *MessageListener) SetListener(messageId MessageId, f func(interface{}) bool) {
	o.listener[messageId] = f
}

func (o *MessageListener) GetListener(messageId MessageId) (func(interface{}) bool, bool) {
	r, ok := o.listener[messageId]
	return r, ok
}

func (o *MessageListener) ProcessMessage(messageId MessageId, messageParams interface{}) bool {
	if m, ok := o.listener[messageId]; ok {
		return m(messageParams)
	}
	return false
}

type IMessageSender interface {
	SendMessage(MessageId, interface{}, string)
	SetSendMessage(func(MessageId, interface{}, string))
}

type MessageSender struct {
	sendMessage func(MessageId, interface{}, string)
}

func (o *MessageSender) SendMessage(messageId MessageId, messageParams interface{}, entityId string) {
	if o.sendMessage != nil {
		o.sendMessage(messageId, messageParams, entityId)
	}
}

func (o *MessageSender) SetSendMessage(f func(MessageId, interface{}, string)) {
	o.sendMessage = f
}
