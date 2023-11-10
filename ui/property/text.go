package property

type PropertyText struct {
	text     string
	onChange func()
}

func (o *PropertyText) Text() string { return o.text }

func (o *PropertyText) SetText(s string) {
	o.text = s
	if o.onChange != nil {
		o.onChange()
	}
}

func (o *PropertyText) SetOnTextChange(f func()) { o.onChange = f }
