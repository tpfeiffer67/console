package property

import "github.com/tpfeiffer67/console/screen"

type IOnRender interface {
	SetOnRender(func(*screen.Buffer, screen.Coordinates))
	GetOnRender() func(*screen.Buffer, screen.Coordinates)
	DoOnRender(*screen.Buffer, screen.Coordinates)
}

type OnRender struct {
	onRender func(*screen.Buffer, screen.Coordinates)
}

func (o *OnRender) SetOnRender(f func(*screen.Buffer, screen.Coordinates)) { o.onRender = f }

func (o *OnRender) GetOnRender() func(*screen.Buffer, screen.Coordinates) { return o.onRender }

func (o *OnRender) DoOnRender(buf *screen.Buffer, pos screen.Coordinates) {
	if o.onRender != nil {
		o.onRender(buf, pos)
	}
}
