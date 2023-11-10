package engine

type OnModeChange struct {
	onModeChange func(mode bool)
}

func (o *OnModeChange) SetOnModeChange(f func(bool)) {
	o.onModeChange = f
}

func (o *OnModeChange) GetOnModeChange() func(bool) {
	return o.onModeChange
}

func (o *OnModeChange) doOnModeChange(mode bool) {
	if o.onModeChange != nil {
		o.onModeChange(mode)
	}
}
