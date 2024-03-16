package engine

import (
	"github.com/tpfeiffer67/console/ui/property"
)

func (o *Engine) GetFocusedEntity() (any, bool) {
	if o.FocusedEntity != nil {
		return o.FocusedEntity, true
	}
	return nil, false
}

func (o *Engine) GetFocusedEntityId() (string, bool) {
	if o.FocusedEntity != nil {
		return o.FocusedEntity.(property.IId).Id(), true // #directId
	}
	return "", false
}

func (o *Engine) FocusById(id string) bool {
	if v, ok := o.GetEntityById(id); ok {
		return o.Focus(v)
	}
	return false
}

func (o *Engine) Focus(a any) bool {
	if a != nil {
		if e, ok := a.(property.IFocus); ok {
			if e.Focusable() {
				if o.FocusedEntity != nil {
					if fe, ok := o.FocusedEntity.(property.IFocus); ok {
						fe.SetFocused(false, e)
					}
				}
				e.SetFocused(true, o.FocusedEntity)
				o.FocusedEntity = e
				return true
			}
		}
	}
	return false
}

func (o *Engine) SetFocusedGroupFromTheTopMostAncestorEntity(id string, b bool) {
	top := o.GetTheTopMostAncestorEntity(id)
	o.CallFuncForEntityAndAllItsDescendants(top.(property.IId).Id(), // #directId
		func(a any) {
			if e, ok := a.(property.IFocus); ok {
				e.SetFocusedGroup(b)
			}
		})
}

func (o *Engine) FocusParent() bool {
	if o.FocusedEntity != nil {
		if e, ok := o.FocusedEntity.(property.IParent); ok {
			return o.FocusById(e.Parent())
		}
	}
	return false
}

func (o *Engine) FocusChild() bool {
	if id, ok := o.GetFocusedEntityId(); ok {
		l := o.listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(id)
		l = sortEntitiesByFocusOrder(l)
		if len(l) > 0 {
			return o.Focus(l[0])
		}
	}
	return false
}

func (o *Engine) FocusNext() bool {
	if o.FocusedEntity != nil {
		if e, ok := o.FocusedEntity.(property.IParentVisibleFocusable); ok {
			currentFocusOrder := e.FocusOrder()
			p := e.Parent()
			l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(p)
			l = sortEntitiesByFocusOrder(l)
			for _, a := range l {
				if e, ok := a.(property.IFocus); ok {
					if e.FocusOrder() > currentFocusOrder {
						return o.Focus(e)
					}
				}
			}
		}
	}
	return false
}

func (o *Engine) FocusPrevious() bool {
	if o.FocusedEntity != nil {
		if e, ok := o.FocusedEntity.(property.IParentVisibleFocusable); ok {
			currentFocusOrder := e.FocusOrder()
			p := e.Parent()
			l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(p)
			l = sortEntitiesByReverseFocusOrder(l)
			for _, a := range l {
				if e, ok := a.(property.IFocus); ok {
					if e.FocusOrder() < currentFocusOrder {
						return o.Focus(e)
					}
				}
			}
		}
	}
	return false
}
