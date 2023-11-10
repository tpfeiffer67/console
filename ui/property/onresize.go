package property

type IOnResize interface {
	SetOnResize(func(int, int))
	DoOnResize(int, int)
}

type OnResize struct {
	onResize func(int, int)
}

func (o *OnResize) SetOnResize(f func(int, int)) { o.onResize = f }
func (o *OnResize) DoOnResize(h int, w int) {
	if o.onResize != nil {
		o.onResize(h, w)
	}
}
