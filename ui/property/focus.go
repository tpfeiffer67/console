package property

type IFocus interface {
	Focused() bool
	SetFocused(bool, IFocus)
	FocusedGroup() bool
	SetFocusedGroup(b bool)
	Focusable() bool
	SetFocusable(bool)
	FocusOrder() int
	SetFocusOrder(int)
	SetOnFocus(func(IFocus))
	DoOnFocus(IFocus)
	SetOnLostFocus(func(IFocus))
	DoOnLostFocus(IFocus)
	SetOnFocusGroup(func())
	DoOnFocusGroup()
	SetOnLostFocusGroup(func())
	DoOnLostFocusGroup()
}

type Focus struct {
	focused          bool
	focusedGroup     bool
	focusable        bool
	focusOrder       int
	onFocus          func(IFocus)
	onLostFocus      func(IFocus)
	onFocusGroup     func()
	onLostFocusGroup func()
}

func (o *Focus) Focused() bool { return o.focused }
func (o *Focus) SetFocused(b bool, foc IFocus) {
	if !o.Focusable() && b {
		return
	}
	if b != o.focused {
		o.focused = b
		switch b {
		case true:
			o.DoOnFocus(foc)
		case false:
			o.DoOnLostFocus(foc)
		}
	}
}

func (o *Focus) FocusedGroup() bool { return o.focusedGroup }
func (o *Focus) SetFocusedGroup(b bool) {
	if !o.Focusable() && b {
		return
	}
	if b != o.focusedGroup {
		o.focusedGroup = b
		switch b {
		case true:
			o.DoOnFocusGroup()
		case false:
			o.DoOnLostFocusGroup()
		}

	}
}

func (o *Focus) Focusable() bool     { return o.focusable }
func (o *Focus) SetFocusable(b bool) { o.focusable = b }
func (o *Focus) FocusOrder() int     { return o.focusOrder }
func (o *Focus) SetFocusOrder(i int) { o.focusOrder = i }

func (o *Focus) SetOnFocus(f func(IFocus)) { o.onFocus = f }
func (o *Focus) DoOnFocus(previousIFocus IFocus) {
	if o.onFocus != nil {
		o.onFocus(previousIFocus)
	}
}

func (o *Focus) SetOnLostFocus(f func(IFocus)) { o.onLostFocus = f }
func (o *Focus) DoOnLostFocus(nextIFocus IFocus) {
	if o.onLostFocus != nil {
		o.onLostFocus(nextIFocus)
	}
}

func (o *Focus) SetOnFocusGroup(f func()) { o.onFocusGroup = f }
func (o *Focus) DoOnFocusGroup() {
	if o.onFocusGroup != nil {
		o.onFocusGroup()
	}
}

func (o *Focus) SetOnLostFocusGroup(f func()) { o.onLostFocusGroup = f }
func (o *Focus) DoOnLostFocusGroup() {
	if o.onLostFocusGroup != nil {
		o.onLostFocusGroup()
	}
}
