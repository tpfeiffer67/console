package property

// TODO Refactoring
/*
type IOnMouse struct {
	IPropertyMouse
	IOnMouseClick
	IOnMouseClickDown
	IOnMouseClickUp
}
*/
type IMouse interface {
	Hovered() bool
	SetHovered(bool)
	HoveredGroup() bool
	SetHoveredGroup(bool)
	SetOnMouseEnter(func())
	DoOnMouseEnter()
	SetOnMouseExit(func())
	DoOnMouseExit()
}

/*
type OnMouse struct {
	Mouse
	OnMouseClick
	OnMouseClickDown
	OnMouseClickUp
}
*/
type Mouse struct {
	hovered      bool
	hoveredGroup bool
	onMouseEnter func()
	onMouseExit  func()
}

func (o *Mouse) Hovered() bool { return o.hovered }
func (o *Mouse) SetHovered(b bool) {
	if b != o.hovered {
		o.hovered = b
		switch b {
		case true:
			o.DoOnMouseEnter()
		case false:
			o.DoOnMouseExit()
		}
	}
}

func (o *Mouse) HoveredGroup() bool     { return o.hoveredGroup }
func (o *Mouse) SetHoveredGroup(b bool) { o.hoveredGroup = b }

func (o *Mouse) SetOnMouseEnter(f func()) { o.onMouseEnter = f }
func (o *Mouse) DoOnMouseEnter() {
	if o.onMouseEnter != nil {
		o.onMouseEnter()
	}
}

func (o *Mouse) SetOnMouseExit(f func()) { o.onMouseExit = f }
func (o *Mouse) DoOnMouseExit() {
	if o.onMouseExit != nil {
		o.onMouseExit()
	}
}

/*
type IOnMouseClickDown interface {
	SetOnMouseClickDown(f func(message.ParamsMouse))
	DoOnMouseClickDown(message.ParamsMouse)
}

type OnMouseClickDown struct {
	onMouseClickDown func(message.ParamsMouse)
}

func (o *OnMouseClickDown) SetOnMouseClickDown(f func(message.ParamsMouse)) { o.onMouseClickDown = f }
func (o *OnMouseClickDown) DoOnMouseClickDown(m message.ParamsMouse) {
	if o.onMouseClickDown != nil {
		o.onMouseClickDown(m)
	}
}

type IOnMouseClickUp interface {
	SetOnMouseClickUp(f func(message.ParamsMouse))
	DoOnMouseClickUp(message.ParamsMouse)
}

type OnMouseClickUp struct {
	onMouseClickUp func(message.ParamsMouse)
}

func (o *OnMouseClickUp) SetOnMouseClickUp(f func(message.ParamsMouse)) { o.onMouseClickUp = f }
func (o *OnMouseClickUp) DoOnMouseClickUp(m message.ParamsMouse) {
	if o.onMouseClickUp != nil {
		o.onMouseClickUp(m)
	}
}

type IOnMouseClick interface {
	SetOnMouseClick(f func(message.ParamsMouse))
	DoOnMouseClick(message.ParamsMouse)
}

type OnMouseClick struct {
	onMouseClick func(message.ParamsMouse)
}

func (o *OnMouseClick) SetOnMouseClick(f func(message.ParamsMouse)) { o.onMouseClick = f }
func (o *OnMouseClick) DoOnMouseClick(m message.ParamsMouse) {
	if o.onMouseClick != nil {
		o.onMouseClick(m)
	}
}
*/
