package property

type IOnChange interface {
	SetOnChange(func())
	DoOnChange()
}

type OnChange struct {
	onChange func()
}

func (o *OnChange) SetOnChange(f func()) { o.onChange = f }
func (o *OnChange) DoOnChange() {
	if o.onChange != nil {
		o.onChange()
	}
}
