package engine

import (
	"sort"

	"github.com/tpfeiffer67/console/ui/ntt"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Entities map[string]ntt.IEntity

func (o Entities) GetEntityById(id string) (ntt.IEntity, bool) {
	if e, ok := o[id]; ok {
		return e, true
	}
	return nil, false
}

func (o Entities) GetListOfChildren(id string) []ntt.IEntity {
	l := o.listOfEntitiesWhoseParentHasTheGivenId(id)
	return l
}

func (o Entities) GetListOfBrotherhood(id string) []ntt.IEntity {
	if parentId, ok := o.getParentEntityId(id); ok {
		l := o.listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(parentId)
		return l
	}
	return make([]ntt.IEntity, 0)
}

func (o Entities) getParentEntityId(id string) (string, bool) {
	if n, ok := o.GetEntityById(id); ok {
		return n.Parent(), true
	}
	return "", false
}

func (o Entities) listOfEntitiesWhoseParentHasTheGivenId(id string) []ntt.IEntity {
	return o.filter(func(e ntt.IEntity) bool {
		return (e.Parent() == id)
	})
}

func (o Entities) listOfEntityAndHisAscendants(id string) []ntt.IEntity {
	l := make([]ntt.IEntity, 0)
	for {
		if e, ok := o.GetEntityById(id); ok {
			l = append(l, e)
			id = e.Parent()
			if id == "" {
				break
			}
		}
	}
	return l
}

func (o Entities) listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(id string) []ntt.IEntity {
	return o.filter(func(e ntt.IEntity) bool {
		return (e.Parent() == id && e.Visible())
	})
}

func (o Entities) CallFuncForEachChildrenEntity(id string, f func(e ntt.IEntity)) {
	l := o.listOfEntitiesWhoseParentHasTheGivenId(id)
	l = sortEntitiesByOrder(l)
	callFuncWithEachEntityOfTheList(l, f)
}

func (o Entities) CallFuncForEntityAndHisAscendants(id string, f func(e ntt.IEntity)) {
	l := o.listOfEntityAndHisAscendants(id)
	callFuncWithEachEntityOfTheList(l, f)
}

func (o Entities) GetFirstChildEntity(id string) (ntt.IEntity, bool) {
	l := o.listOfEntitiesWhoseParentHasTheGivenId(id)
	l = sortEntitiesByOrder(l)
	if len(l) > 0 {
		return l[0], true
	}
	return nil, false
}

func (o Entities) nextOrder(id string) int {
	return o.countChildrenForGivenParentId(id) + 1
}

func (o Entities) countChildrenForGivenParentId(id string) int {
	c := 0
	for _, e := range o {
		if e.Parent() == id {
			c++
		}
	}
	return c
}

func (o Entities) callFuncWithAllEntities(f func(e ntt.IEntity)) {
	for _, e := range o {
		f(e)
	}
}

func (o Entities) CallFuncWithAllEntitiesAsInterface(f func(i interface{})) {
	for _, i := range o {
		f(i)
	}
}

func callFuncWithEachEntityOfTheList(l []ntt.IEntity, f func(e ntt.IEntity)) {
	for _, e := range l {
		f(e)
	}
}

func getEntityIndex(l []ntt.IEntity, id string) (int, bool) {
	for i, e := range l {
		if e.Id() == id {
			return i, true
		}
	}
	return -1, false
}

func (o Entities) GetTheTopMostAncestorEntity(id string) ntt.IEntity {
	for {
		if e, ok := o.GetEntityById(id); ok {
			id = e.Parent()
			if id == "" {
				return e
			}
		}
	}
}

func (o Entities) CallFuncForEntityAndAllItsDescendants(id string, f func(e ntt.IEntity)) {
	if n, ok := o.GetEntityById(id); ok {
		o.callFuncForEntityAndAllItsDescendants(n, f)
	}
}

func (o Entities) callFuncForEntityAndAllItsDescendants(n ntt.IEntity, f func(e ntt.IEntity)) {
	f(n)
	l := o.listOfEntitiesWhoseParentHasTheGivenId(n.Id())
	for _, e := range l {
		o.callFuncForEntityAndAllItsDescendants(e, f)
	}
}

func (o Entities) listEntities_WhoseParentIsId(id string) []ntt.IEntity {
	return o.filter(func(e ntt.IEntity) bool {
		return (e.Parent() == id)
	})
}

func (o Entities) listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(id string) []ntt.IEntity {
	return o.filter(func(e ntt.IEntity) bool {
		return (e.Parent() == id && e.Visible() && e.Focusable())
	})
}

func (o Entities) listEntities_WhoseParentIsId_AndWhichAreVisible_AndWithGivenZOrderLayer(id string, zlayer int) []ntt.IEntity {
	return o.filter(func(e ntt.IEntity) bool {
		return (e.Parent() == id && e.Visible() && e.ZOrderLayer() == zlayer)
	})
}

func sortEntitiesByOrder(l []ntt.IEntity) []ntt.IEntity {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].Order() < l[j].Order()
	})
	return l
}

func sortEntitiesByZOrder(l []ntt.IEntity) []ntt.IEntity {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].ZOrderSum() < l[j].ZOrderSum()
	})
	return l
}

func sortEntitiesByFocusOrder(l []ntt.IEntity) []ntt.IEntity {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].FocusOrder() < l[j].FocusOrder()
	})
	return l
}

func sortEntitiesByReverseFocusOrder(l []ntt.IEntity) []ntt.IEntity {
	sort.SliceStable(l, func(i, j int) bool {
		return l[i].FocusOrder() > l[j].FocusOrder()
	})
	return l
}

// TODO Factorize ZOrderUp and ZOrderDown
func (o Entities) ZOrderUp(id string) bool {
	if e, ok := o.GetEntityById(id); ok {
		parentId := e.Parent()
		l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWithGivenZOrderLayer(parentId, e.ZOrderLayer())
		l = sortEntitiesByZOrder(l)
		i := len(l)
		if i > 1 {
			if ei, ok := getEntityIndex(l, id); ok {
				if ei < i-1 {
					swapZOrder(e, l[ei+1])
					return true
				}
			}
		}
	}
	return false
}

func (o Entities) ZOrderDown(id string) bool {
	if e, ok := o.GetEntityById(id); ok {
		parentId := e.Parent()
		l := o.listEntities_WhoseParentIsId_AndWhichAreVisible_AndWithGivenZOrderLayer(parentId, e.ZOrderLayer())
		l = sortEntitiesByZOrder(l)
		i := len(l)
		if i > 1 {
			if ei, ok := getEntityIndex(l, id); ok {
				if ei > 0 {
					swapZOrder(e, l[ei-1])
					return true // true means ZOrder has changed
				}
			}
		}
	}
	return false // true means ZOrder has not changed
}

func swapZOrder(e1, e2 ntt.IEntity) {
	z1 := e1.ZOrder()
	z2 := e2.ZOrder()
	e1.SetZOrder(z2)
	e2.SetZOrder(z1)
}

func (o Entities) ZOrderToTop(id string) {
	for o.ZOrderUp(id) {
	}
}

func (o Entities) ZOrderToBottom(id string) {
	for o.ZOrderDown(id) {
	}
}

func (o Entities) UpdateEntitiesTheme() {
	f := func(i interface{}) {
		if e, ok := i.(ntt.IEntityWithTheme); ok {
			e.UpdateTheme(theme.CurrentTheme)
			e.DoOnThemeChange()
		}
	}
	o.CallFuncWithAllEntitiesAsInterface(f)
}

func (o Entities) filter(filter func(ntt.IEntity) bool) []ntt.IEntity {
	l := make([]ntt.IEntity, 0)
	for _, e := range o {
		if filter(e) {
			l = append(l, e)
		}
	}
	return l
}
