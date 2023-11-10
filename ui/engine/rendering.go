package engine

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/ntt"
)

func (o *Engine) renderEntitiesToScreenBuffer() {
	o.renderEntities("", 0, 0)
	o.renderMousePointer()
}

func (o *Engine) renderEntities(id string, row, col int) {
	l := o.listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(id)
	l = sortEntitiesByZOrder(l)
	o.drawEntities(l)
	for _, v := range l {
		r := row + v.Row()
		c := col + v.Col()
		v.DoOnRender(o.Screen.Buffer, screen.Coordinates{Row: r, Col: c})
		o.renderEntities(v.Id(), r, c)
	}
}

func (o *Engine) drawEntities(l []ntt.IEntity) {
	for _, v := range l {
		v.DoOnDraw()
	}
}

func (o *Engine) renderMousePointer() {
	o.pointer.RenderAtPosition(o.Screen.Buffer, screen.Coordinates{Row: o.mouseParams.Row, Col: o.mouseParams.Col})
}
