package engine

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/property"
)

/*
type EntityRenderer interface {
	DoOnRender(screen.Buffer, screen.Coordinates) // provided by OnRender
	Parent()                                      // provided by PropertyId
	Visible()                                     // provided by PropertyVisible
	ZOrderSum()                                   // provided by PropertyZOrder
	DoOnDraw()                                    // non
}
*/

func (o *Engine) renderEntitiesToScreenBuffer() {
	o.renderEntities("", 0, 0)
	o.renderMousePointer()
}

type Renderer interface {
	property.IId
	property.IOnRender
	screen.PositionGetter
}

func (o *Engine) renderEntities(id string, row, col int) {
	l := o.listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(id)
	l = sortEntitiesByZOrder(l)
	o.drawEntities(l) // TODO remove from here
	for _, a := range l {
		if e, ok := a.(Renderer); ok {
			r := row + e.Row()
			c := col + e.Col()
			e.DoOnRender(o.Screen.Buffer, screen.Coordinates{Row: r, Col: c})
			o.renderEntities(e.Id(), r, c) // recursiv call
		}
	}
}

func (o *Engine) drawEntities(l []any) {
	for _, a := range l {
		if e, ok := a.(property.IOnDraw); ok {
			e.DoOnDraw()
		}
	}
}

func (o *Engine) renderMousePointer() {
	o.pointer.RenderAtPosition(o.Screen.Buffer, screen.Coordinates{Row: o.mouseParams.Row, Col: o.mouseParams.Col})
}
