package property

type IOnDraw interface {
	SetOnDraw(func())
	GetOnDraw() func()
	DoOnDraw()
}

type OnDraw struct {
	onDraw func()
}

func (o *OnDraw) SetOnDraw(f func()) { o.onDraw = f }

func (o *OnDraw) GetOnDraw() func() { return o.onDraw }

func (o *OnDraw) DoOnDraw() {
	if o.onDraw != nil {
		o.onDraw()
	}
}
