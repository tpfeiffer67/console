package property

type IOnScreenResize interface {
	SetOnScreenResize(func(int, int))
	DoOnScreenResize(int, int)
}

type OnScreenResize struct {
	onScreenResize func(int, int)
}

func (o *OnScreenResize) SetOnScreenResize(f func(int, int)) { o.onScreenResize = f }
func (o *OnScreenResize) DoOnScreenResize(h int, w int) {
	if o.onScreenResize != nil {
		o.onScreenResize(h, w)
	}
}
