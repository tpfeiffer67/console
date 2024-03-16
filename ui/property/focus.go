package property

type IFocus interface {
	Focused() bool
	SetFocused(bool, any)
	FocusedGroup() bool
	SetFocusedGroup(b bool)
	Focusable() bool
	SetFocusable(bool)
	FocusOrder() int
	SetFocusOrder(int)
	SetOnFocus(func(any))
	DoOnFocus(any)
	SetOnLostFocus(func(any))
	DoOnLostFocus(any)
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
	onFocus          func(any)
	onLostFocus      func(any)
	onFocusGroup     func()
	onLostFocusGroup func()
}

func (o *Focus) Focused() bool { return o.focused }
func (o *Focus) SetFocused(b bool, foc any) {
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

func (o *Focus) SetOnFocus(f func(any)) { o.onFocus = f }
func (o *Focus) DoOnFocus(previous any) {
	if o.onFocus != nil {
		o.onFocus(previous)
	}
}

func (o *Focus) SetOnLostFocus(f func(any)) { o.onLostFocus = f }
func (o *Focus) DoOnLostFocus(next any) {
	if o.onLostFocus != nil {
		o.onLostFocus(next)
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
