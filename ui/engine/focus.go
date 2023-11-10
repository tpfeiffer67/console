package engine

import (
	"github.com/tpfeiffer67/console/ui/ntt"
)

func (o *Engine) GetFocusedEntity() (ntt.IEntity, bool) {
	if o.FocusedEntity != nil {
		return o.FocusedEntity, true
	}
	return nil, false
}

func (o *Engine) GetFocusedEntityId() (string, bool) {
	if o.FocusedEntity != nil {
		return o.FocusedEntity.Id(), true
	}
	return "", false
}

func (o *Engine) FocusById(id string) bool {
	if v, ok := o.GetEntityById(id); ok {
		return o.Focus(v)
	}
	return false
}

func (o *Engine) Focus(v ntt.IEntity) bool {
	if v != nil {
		if v.Focusable() {
			if o.FocusedEntity != nil {
				o.FocusedEntity.SetFocused(false, v)
			}
			v.SetFocused(true, o.FocusedEntity)
			o.FocusedEntity = v
			return true
		}
	}
	return false
}

func (o *Engine) SetFocusedGroupFromTheTopMostAncestorEntity(id string, b bool) {
	top := o.GetTheTopMostAncestorEntity(id)
	o.CallFuncForEntityAndAllItsDescendants(top.Id(), func(n ntt.IEntity) { n.SetFocusedGroup(b) })
}

func (o *Engine) FocusParent() bool {
	if o.FocusedEntity != nil {
		return o.FocusById(o.FocusedEntity.Parent())
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
		currentFocusOrder := o.FocusedEntity.FocusOrder()
		p := o.FocusedEntity.Parent()
		l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(p)
		l = sortEntitiesByFocusOrder(l)
		for _, v := range l {
			if v.FocusOrder() > currentFocusOrder {
				return o.Focus(v)
			}
		}
	}
	return false
}

func (o *Engine) FocusPrevious() bool {
	if o.FocusedEntity != nil {
		currentFocusOrder := o.FocusedEntity.FocusOrder()
		p := o.FocusedEntity.Parent()
		l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(p)
		l = sortEntitiesByReverseFocusOrder(l)
		for _, v := range l {
			if v.FocusOrder() < currentFocusOrder {
				return o.Focus(v)
			}
		}
	}
	return false
}
