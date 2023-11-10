package property

type IVisible interface {
	Visible() bool
	SetVisible(bool)
	CanShow() bool
	SetCanShow(bool)
	CanHide() bool
	SetCanHide(bool)
	SetOnShow(func())
	DoOnShow()
	SetOnHide(func())
	DoOnHide()
}

type PropertyVisible struct {
	visible bool
	canShow bool
	canHide bool
	onShow  func()
	onHide  func()
}

func (o *PropertyVisible) Visible() bool { return o.visible }
func (o *PropertyVisible) SetVisible(b bool) {
	if !o.canShow && b {
		return
	}
	if !o.canHide && !b {
		return
	}
	if b != o.visible {
		o.visible = b
		switch b {
		case true:
			o.DoOnShow()
		case false:
			o.DoOnHide()
		}
	}
}

func (o *PropertyVisible) CanShow() bool     { return o.canShow }
func (o *PropertyVisible) SetCanShow(b bool) { o.canShow = b }
func (o *PropertyVisible) CanHide() bool     { return o.canHide }
func (o *PropertyVisible) SetCanHide(b bool) { o.canHide = b }

func (o *PropertyVisible) SetOnShow(f func()) { o.onShow = f }
func (o *PropertyVisible) DoOnShow() {
	if o.onShow != nil {
		o.onShow()
	}
}

func (o *PropertyVisible) SetOnHide(f func()) { o.onHide = f }
func (o *PropertyVisible) DoOnHide() {
	if o.onHide != nil {
		o.onHide()
	}
}
