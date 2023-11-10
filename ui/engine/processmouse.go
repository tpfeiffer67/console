package engine

import (
	"time"

	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/ntt"
)

type mouseControl struct {
	pointer          *ntt.Pointer
	pointedEntity    ntt.IEntity
	mouseParams      message.ParamsMouse
	mouseControlMode MouseControlMode
	memoMouseRow     int
	memoMouseCol     int
	memoEntity       ntt.IEntity
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
					o.mouseControlMode = MouseControlModeMove
					o.memoMouseRow = mouseParams.Row
					o.memoMouseCol = mouseParams.Col
					o.memoEntityRow = o.pointedEntity.Row()
					o.memoEntityCol = o.pointedEntity.Col()
					o.memoEntity = o.pointedEntity
				}
			}

			if !o.mouseParams.ButtonSecondary && mouseParams.ButtonSecondary {
				if o.pointedEntity != nil {
					o.mouseControlMode = MouseControlModeResize
					o.memoMouseRow = mouseParams.Row
					o.memoMouseCol = mouseParams.Col
					o.memoEntityHeight = o.pointedEntity.Height()
					o.memoEntityWidth = o.pointedEntity.Width()
					o.memoEntity = o.pointedEntity
				}
			}

		case MouseControlModeMove:
			if mouseParams.ButtonPrimary {
				if o.memoEntity != nil {
					o.memoEntity.SetPosition(o.memoEntityRow+mouseParams.Row-o.memoMouseRow, o.memoEntityCol+mouseParams.Col-o.memoMouseCol)
				}
			} else {
				o.mouseControlMode = MouseControlModeDefault
			}

		case MouseControlModeResize:
			if mouseParams.ButtonSecondary {
				if o.memoEntity != nil {
					o.memoEntity.Resize(o.memoEntityHeight+mouseParams.Row-o.memoMouseRow, o.memoEntityWidth+mouseParams.Col-o.memoMouseCol)
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
					id := o.memoEntity.Id()
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
	x := o.getMetaUnderMousePosition(mouseParams)
	if pointedEntity, ok := x.(ntt.IEntity); ok {
		if o.pointedEntity != pointedEntity {
			if o.pointedEntity != nil {
				//o.pointedEntity.NextPointedEntity = pointedEntity
				o.updateHoveredGroup(o.pointedEntity.Id(), false)
				o.pointedEntity.SetHovered(false)
				if sendMessages {
					o.SendMessage(message.MessageIdMouseExit, mouseParams, o.pointedEntity.Id())
				}
			}
			o.updateHoveredGroup(pointedEntity.Id(), true)
			pointedEntity.SetHovered(true)
			if sendMessages {
				o.SendMessage(message.MessageIdMouseEnter, mouseParams, pointedEntity.Id())
			}
			o.pointedEntity = pointedEntity
		}
		return
	}
	o.pointedEntity = nil
}

func (o *Engine) updateHoveredGroup(id string, b bool) {
	l := o.listOfEntityAndHisAscendants(id)
	callFuncWithEachEntityOfTheList(l, func(n ntt.IEntity) { n.SetHoveredGroup(b) })
}
