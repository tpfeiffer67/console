package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type ISystem interface {
	CallFuncWithAllEntities(func(any))
	CallFuncForEachChildrenEntity(string, func(any))
	GetFirstChildEntity(string) (any, bool)
	GetListOfChildren(string) []any
	GetListOfBrotherhood(string) []any
	SetFocusedGroupFromTheTopMostAncestorEntity(string, bool)
	GetTheTopMostAncestorEntity(string) any
	Focus(any) bool
	ZOrderToTop(string)
	GetTheme() theme.ITheme
}

type IWidget interface {
	screen.Sizer
	theme.ITheme
	screen.Positioner
	property.IId
	property.IParent
	property.IVisible
	property.IOrder
	property.IZOrder
	property.IFocus
	property.IMouse
	property.IOnResize
	property.IOnScreenResize
	property.IOnDraw
	property.IOnRender
	property.IOnKeyPress
	property.IOnThemeChange
	message.IMessageListener
	message.IMessageSender
	screenutils.ICanvas
	Render(*screen.Buffer, screen.Coordinates)
}

type Widget struct {
	ISystem
	theme.ITheme
	screen.Position
	property.PropertyId
	property.PropertyVisible
	property.PropertyOrder
	property.PropertyZOrder
	property.Focus
	property.Mouse
	property.OnResize
	property.OnScreenResize
	property.OnDraw
	property.OnRender
	property.OnKeyPress
	property.OnThemeChange
	message.MessageListener
	message.MessageSender
	*screenutils.ComboCanvas
}

func NewWidget(id string, height int, width int, syst ISystem) *Widget {
	o := new(Widget)
	o.ITheme = syst.GetTheme()
	o.SetId(id)
	o.ISystem = syst
	o.MessageListener = message.NewMessageListener()
	o.ComboCanvas = screenutils.NewComboCanvas(height, width)
	o.SetOnRender(o.Render)
	o.SetCanShow(true)
	o.SetCanHide(true)
	o.SetVisible(true)
	return o
}

func (o *Widget) Render(sb *screen.Buffer, pos screen.Coordinates) {
	if o.Visible() {
		sb.DefineMeta(o)
		o.RuneCanvas.Render(sb, o.Stencil, pos, o.Origins)
		sb.ResetMeta()
		o.EmphasisCanvas.Render(sb, o.Stencil, pos, o.Origins)
		o.FColorCanvas.Render(sb, o.Stencil, pos, o.Origins)
		o.BColorCanvas.Render(sb, o.Stencil, pos, o.Origins)
	}
}
