package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

/*
	type Paneler interface {
		IWidget
		theme.ITheme
		//SizeManager
		Minimize(any)
		RestoreSize()
	}
*/
type Panel struct {
	IWidget
	//SizeManager
	ISystem
}

func NewPanel(id string, row, col int, height, width int, syst ISystem) *Panel {
	/* #minimize
	clickWidth := 0
	*/
	o := new(Panel)
	o.IWidget = NewWidget(id, height, width, syst)
	//o.SizeManager = new(SizeManagement) #minimizer
	o.ISystem = syst

	o.SetPosition(row, col)
	o.SetCanMove(true)
	o.SetFocusable(true)

	//SetDefaultFuncFor_OnFocus_And_OnLostFocus(o)
	o.SetOnFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc any) {
		syst.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})

	o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
		o.Render(sb, pos)
	})

	o.SetOnDraw(func() {
		ClearWithStyle(o, o.IWidget, theme.STYLE_PANEL, theme.STYLE_PANEL_HOVERED, theme.STYLE_PANEL_FOCUSED, theme.STYLE_PANEL_FOCUSEDHOVERED)
	})

	/* #minimize
	onThemeChanged := func() {
		clickWidth, _ = o.GetIntDef(theme.MINIMIZING_CLICK_WIDTH, DEFAULT_MINIMIZE_CLICKZONE_WIDTH)
	}
	o.SetOnThemeChange(onThemeChanged)
	onThemeChanged()
	*/

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		o.SendMessage(message.MessageIdZOrderTop, nil, id)
		o.SendMessage(message.MessageIdFocus, nil, id)
		return true
	})

	mouseResizeHeightMode := false
	mouseResizeWidthMode := false
	mouseMove := false
	o.SetListener(message.MessageIdMouseDown, func(messageParams interface{}) bool {
		o.ZOrderToTop(o.Id())
		mouseResizeHeightMode = false
		mouseResizeWidthMode = false
		if params, ok := messageParams.(message.ParamsMouse); ok {
			/*if o.SizeStatus() == SizeStatusMinimized {
				o.SendMessage(message.MessageIdRestoreSize, nil, o.Id())
				return true
			} */ //#minimizer

			relativeRow := params.Row - o.Row() // TODO fix bug when minimized
			relativeCol := params.Col - o.Col()
			/* #minimize
			if relativeRow == 0 && relativeCol < clickWidth {
				o.SendMessage(message.MessageIdMinimize, nil, o.Id())
				return true
			}
			*/
			if relativeRow == o.Height()-1 { // bottom border
				mouseResizeHeightMode = true
				if relativeCol == o.Width()-1 { // left-bottom corner
					mouseResizeWidthMode = true
				}
				return true
			}
			if relativeCol == o.Width()-1 { // left border only
				mouseResizeWidthMode = true
				return true
			}
		}
		mouseMove = true
		return true
	})

	o.SetListener(message.MessageIdMouseUp, func(messageParams interface{}) bool {
		mouseResizeHeightMode = false
		mouseResizeWidthMode = false
		mouseMove = false
		return true
	})

	o.SetListener(message.MessageIdMouseMove, func(messageParams interface{}) bool {
		if params, ok := messageParams.(message.ParamsMouseExt); ok {
			if mouseMove {
				o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: params.DeltaRow, Cols: params.DeltaCol}, o.Id())
				return true
			}

			h := o.Height()
			w := o.Width()
			if mouseResizeHeightMode {
				if params.Row > o.Row() {
					h = params.Row - o.Row() + 1
					//o.SetHeight(params.Row - o.Row() + 1)
				}
			}
			if mouseResizeWidthMode {
				if params.Col > o.Col() {
					w = params.Col - o.Col() + 1
					//o.SetWidth(params.Col - o.Col() + 1)
				}
			}

			o.SendMessage(message.MessageIdResize, message.ParamsResize{Rows: h, Cols: w}, o.Id())
		}
		return true
	})

	return o
}

func NewPanelWithShadow(id string, row, col int, height, width int, syst ISystem) *Panel {
	o := NewPanel(id, row, col, height, width, syst)
	AddShadow(o)
	return o
}

type FramePanel struct {
	*Panel
	property.PropertyText
}

func NewFramePanel(id string, row, col int, height, width int, title string, syst ISystem) *FramePanel {
	o := new(FramePanel)
	o.Panel = NewPanel(id, row, col, height, width, syst)
	o.SetText(title)

	o.SetOnDraw(func() {
		style := ClearWithStyle(o, o.IWidget, theme.STYLE_PANEL, theme.STYLE_PANEL_HOVERED, theme.STYLE_PANEL_FOCUSED, theme.STYLE_PANEL_FOCUSEDHOVERED)
		frame, _ := o.GetInt(theme.PANEL_FRAME)
		screenutils.DrawFrame(o.GetRuneCanvas(), 0, 0, o.Height(), o.Width(), frame)
		// TODO No magic number, put row et col (0, 2) into theme
		screenutils.DrawStyledString(0, 2, o.Text(), o, style, theme.ToColor)
	})
	return o
}

func NewFramePanelWithShadow(id string, row, col int, height, width int, title string, syst ISystem) *FramePanel {
	o := NewFramePanel(id, row, col, height, width, title, syst)
	AddShadow(o)
	return o
}

func NewColorGradientPanel(id string, row, col int, height, width int, topleft, topright, bottomLeft, bottomright screen.Color, syst ISystem) *Panel {
	o := NewPanelWithShadow(id, row, col, height, width, syst)

	o.SetOnDraw(func() {
		o.Clear()
		screenutils.DrawColorGradientsMatrix(o.GetBColorCanvas(), 0, 0, o.Height(), o.Width(), topleft, topright, bottomLeft, bottomright)
	})
	return o
}

type IShadowable interface {
	property.IOnRender
	screenutils.IStencil
	theme.ITheme
}

func AddShadow(a any) {
	/*
		var thme theme.ITheme
		if thme, ok := a.(theme.ITheme); ok {
			thme.AppendValues(theme.SHADOW_ENABLED, theme.SHADOW_VALUE, theme.SHADOW_VERTICAL_OFFSET, theme.SHADOW_HORIZONTAL_OFFSET)
		}*/
	if e, ok := a.(IShadowable); ok {
		f := e.GetOnRender()
		e.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
			DrawShadowAccordingToTheTheme(sb, e.GetStencil(), pos, e)
			f(sb, pos)
		})
	}
}

type IFramable interface {
	screen.SizeGetter
	screenutils.IHasRuneCanvas
	property.IOnDraw
	theme.ITheme
}

// TODO review this method "SetOnDraw" to add a frame
func AddFrame(a any) {
	if e, ok := a.(IFramable); ok {
		e.SetOnDraw(func() {
			ClearWithStyle(a, e, theme.STYLE_PANEL, theme.STYLE_PANEL_HOVERED, theme.STYLE_PANEL_FOCUSED, theme.STYLE_PANEL_FOCUSEDHOVERED)
			frame, _ := e.GetInt(theme.PANEL_FRAME)
			screenutils.DrawFrame(e.GetRuneCanvas(), 0, 0, e.Height(), e.Width(), frame)
		})
	}
}

/*
	type ColorFilter struct {
		IEntity
		ISystem
	}

	func (o *ColorFilter) SetColor(c screen.Color) {
		o.GetFColorCanvas().DefaultValue = c
		o.GetBColorCanvas().DefaultValue = c
		o.Clear()
	}

	func NewColorFilter(id string, row, col int, height, width int, c screen.Color, syst ISystem) *ColorFilter {
		o := new(ColorFilter)
		o.IEntity = NewEntity(id, height, width, syst)
		o.ISystem = syst

		o.SetPosition(row, col)
		o.SetCanMove(true)
		o.SetFocusable(true)

		o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
			o.Render(sb, pos)
		})

		// TODO Factorisation with NewPanel (all SetListener)
		o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
			o.SendMessage(message.MessageIdZOrderTop, nil, id)
			o.SendMessage(message.MessageIdFocus, nil, id)
			return true
		})

		mouseResizeHeightMode := false
		mouseResizeWidthMode := false
		o.SetListener(message.MessageIdMouseDown, func(messageParams interface{}) bool {
			o.ZOrderToTop(o.Id())
			mouseResizeHeightMode = false
			mouseResizeWidthMode = false
			if params, ok := messageParams.(message.ParamsMouse); ok {
				relativeRow := params.Row - o.Row()
				relativeCol := params.Col - o.Col()
				if relativeRow == o.Height()-1 {
					mouseResizeHeightMode = true
				}
				if relativeCol == o.Width()-1 {
					mouseResizeWidthMode = true
				}
				return true
			}
			return true
		})

		o.SetListener(message.MessageIdMouseMove, func(messageParams interface{}) bool {
			if params, ok := messageParams.(message.ParamsMouseExt); ok {
				if !mouseResizeHeightMode && !mouseResizeWidthMode {
					o.SendMessage(message.MessageIdMove, message.ParamsMove{Rows: params.DeltaRow, Cols: params.DeltaCol}, o.Id())
					return true
				}
				h := o.Height()
				w := o.Width()
				if mouseResizeHeightMode {
					if params.Row > o.Row() {
						h = params.Row - o.Row() + 1
					}
				}
				if mouseResizeWidthMode {
					if params.Col > o.Col() {
						w = params.Col - o.Col() + 1
					}
				}

				o.SendMessage(message.MessageIdResize, message.ParamsResize{Rows: h, Cols: w}, o.Id())
			}
			return true
		})

		o.SetColor(c)

		o.SetOnDraw(func() {
			o.GetFColorCanvas().Clear()
			o.GetBColorCanvas().Clear()
		})

		o.SetOnRender(func(sb *screen.Buffer, pos screen.Coordinates) {
			if o.Visible() {
				sb.DefineMeta(o)
				o.GetFColorCanvas().Render(sb, o.GetStencil(), pos, screen.Coordinates{})
				sb.ResetMeta()
				o.GetBColorCanvas().Render(sb, o.GetStencil(), pos, screen.Coordinates{})
			}
		})

		return o
	}
*/
func (o *Panel) Minimize(a any) {
	/*
		if o.SizeStatus() != SizeStatusMinimized {
			o.MemorizeNormalStatus(o.Row(), o.Col(), o.Height(), o.Width(), o.Parent())
			o.SetSizeStatus(SizeStatusMinimized)
			o.SetParent(e.Id())
			e.SetVisible(true)
		}
	*/
}

func (o *Panel) RestoreSize() {
	/*
		if o.SizeStatus() == SizeStatusMinimized {
			r, c, h, w, p := o.GetMemorizedNormalStatus()
			o.SetParent(p)
			o.SetPosition(r, c)
			o.Resize(h, w)
			o.SetSizeStatus(SizeStatusNormal)
			o.ZOrderToTop(o.Id())
		}
	*/
}
