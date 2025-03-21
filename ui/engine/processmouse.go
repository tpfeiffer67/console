package engine

import (
	"time"

	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
)

type mouseControl struct {
	pointer          *MousePointer
	pointedEntity    any
	mouseParams      message.ParamsMouse
	mouseControlMode MouseControlMode
	memoMouseRow     int
	memoMouseCol     int
	memoEntity       any
	memoEntityRow    int
	memoEntityCol    int
	memoEntityHeight int
	memoEntityWidth  int
	timeAtMouseDown  time.Time
}

/*
type MessageParamsMouse struct {
	Row             int
	Col             int
	ButtonPrimary   bool
	ButtonSecondary bool
	ButtonMiddle    bool
	Shift           bool
	Ctrl            bool
	Alt             bool
}
*/

func (o *Engine) processMouse(mouseParams message.ParamsMouse) {
	o.renderingMutex.Lock()
	defer o.renderingMutex.Unlock()
	if o.EditMode {
		o.processHoverChange(mouseParams, false)

		switch o.mouseControlMode {
		case MouseControlModeDefault:
			if !o.mouseParams.ButtonPrimary && mouseParams.ButtonPrimary {
				if o.pointedEntity != nil {
					if e, ok := o.pointedEntity.(screen.Positioner); ok {
						o.mouseControlMode = MouseControlModeMove
						o.memoMouseRow = mouseParams.Row
						o.memoMouseCol = mouseParams.Col
						o.memoEntityRow = e.Row()
						o.memoEntityCol = e.Col()
						o.memoEntity = o.pointedEntity
					}
				}
			}

			if !o.mouseParams.ButtonSecondary && mouseParams.ButtonSecondary {
				if o.pointedEntity != nil {
					if e, ok := o.pointedEntity.(screen.Sizer); ok {
						o.mouseControlMode = MouseControlModeResize
						o.memoMouseRow = mouseParams.Row
						o.memoMouseCol = mouseParams.Col
						o.memoEntityHeight = e.Height()
						o.memoEntityWidth = e.Width()
						o.memoEntity = o.pointedEntity
					}
				}
			}

		case MouseControlModeMove:
			if mouseParams.ButtonPrimary {
				if o.memoEntity != nil {
					if e, ok := o.memoEntity.(screen.Positioner); ok {
						e.SetPosition(o.memoEntityRow+mouseParams.Row-o.memoMouseRow, o.memoEntityCol+mouseParams.Col-o.memoMouseCol)
					}
				}
			} else {
				o.mouseControlMode = MouseControlModeDefault
			}

		case MouseControlModeResize:
			if mouseParams.ButtonSecondary {
				if o.memoEntity != nil {
					if e, ok := o.memoEntity.(screen.Sizer); ok {
						e.Resize(o.memoEntityHeight+mouseParams.Row-o.memoMouseRow, o.memoEntityWidth+mouseParams.Col-o.memoMouseCol)
					}
				}
			} else {
				o.mouseControlMode = MouseControlModeDefault
			}
		}
	} else { //----------------------------------------------------------------------------------------------------------------------------------------------
		o.processHoverChange(mouseParams, true)
		if o.doOnMouse(&mouseParams, o.mouseParams, o.pointedEntity) {
			return
		}

		if o.mouseParams.Row != mouseParams.Row || o.mouseParams.Col != mouseParams.Col {
			if mouseParams.ButtonPrimary {
				if o.memoEntity != nil {
					id := o.memoEntity.(property.IId).Id() // #directId
					mouseParamsExt := message.ParamsMouseExt{ParamsMouse: mouseParams, DeltaRow: mouseParams.Row - o.memoMouseRow, DeltaCol: mouseParams.Col - o.memoMouseCol}
					o.SendMessage(message.MessageIdMouseMove, mouseParamsExt, id)
				}
			}
		}

		if id, ok := o.GetPointedObjectId(); ok {
			if mouseParams.ButtonPrimary && !o.mouseParams.ButtonPrimary {
				o.timeAtMouseDown = time.Now()
				o.SendMessage(message.MessageIdMouseDown, mouseParams, id)
				o.memoEntity = o.pointedEntity
			}

			if !mouseParams.ButtonPrimary && o.mouseParams.ButtonPrimary {
				currentTime := time.Now()
				if currentTime.Sub(o.timeAtMouseDown) < time.Millisecond*500 {
					o.SendMessage(message.MessageIdMouseClick, mouseParams, id)
				}
				o.SendMessage(message.MessageIdMouseUp, mouseParams, id)
			}
		}
		o.memoMouseRow = mouseParams.Row
		o.memoMouseCol = mouseParams.Col
	}
	o.mouseParams = mouseParams
}

func (o *Engine) processHoverChange(mouseParams message.ParamsMouse, sendMessages bool) {
	a := o.getMetaUnderMousePosition(mouseParams)
	if pointedEntity, ok := a.(property.IMouse); ok {
		if o.pointedEntity != a { // a instead of pointedEntity because we compare two anys
			if o.pointedEntity != nil {
				o.updateHoveredGroup(o.pointedEntity.(property.IId).Id(), false) // #directId
				if e, ok := o.pointedEntity.(property.IMouse); ok {
					e.SetHovered(false)
				}
				if sendMessages {
					o.SendMessage(message.MessageIdMouseExit, mouseParams, o.pointedEntity.(property.IId).Id()) // #directId
				}
			}
			o.updateHoveredGroup(pointedEntity.(property.IId).Id(), true)
			pointedEntity.SetHovered(true)
			if sendMessages {
				o.SendMessage(message.MessageIdMouseEnter, mouseParams, pointedEntity.(property.IId).Id()) // #directId
			}
			o.pointedEntity = pointedEntity
		}
		return
	}
	o.pointedEntity = nil
}

func (o *Engine) updateHoveredGroup(id string, b bool) {
	l := o.listOfEntityAndHisAscendants(id)
	callFuncWithEachEntityOfTheList(l, func(a any) {
		if e, ok := a.(property.IMouse); ok {
			e.SetHoveredGroup(b)
		}
	})
}
