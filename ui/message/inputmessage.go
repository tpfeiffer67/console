package message

import (
	"time"
)

type InputMessageHeader struct {
	TimeStamp time.Time `json:"ts"`
}

type ParamsKey struct {
	Rune  rune   `json:"rune"`
	Key   int    `json:"key"`
	Shift bool   `json:"shift"`
	Ctrl  bool   `json:"ctrl"`
	Alt   bool   `json:"alt"`
	Name  string `json:"name"`
}

type ParamsMouse struct {
	Row             int  `json:"row"`
	Col             int  `json:"col"`
	ButtonPrimary   bool `json:"button1"`
	ButtonSecondary bool `json:"button2"`
	ButtonMiddle    bool `json:"button3"`
	Shift           bool `json:"shift"`
	Ctrl            bool `json:"ctrl"`
	Alt             bool `json:"alt"`
}

type ParamsScreenSize struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type InputMessageKey struct {
	InputMessageHeader
	ParamsKey
}

type InputMessageMouse struct {
	InputMessageHeader
	ParamsMouse
}

type InputMessageScreenSize struct {
	InputMessageHeader
	ParamsScreenSize
}

type InputMessage interface {
	FuncInputMessage()
}

func (o InputMessageHeader) FuncInputMessage() {}
