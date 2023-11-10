package property

type IOnThemeChange interface {
	SetOnThemeChange(func())
	DoOnThemeChange()
}

type OnThemeChange struct {
	onThemeChange func()
}

func (o *OnThemeChange) SetOnThemeChange(f func()) { o.onThemeChange = f }
func (o *OnThemeChange) DoOnThemeChange() {
	if o.onThemeChange != nil {
		o.onThemeChange()
	}
}
