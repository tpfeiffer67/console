package engine

import (
	"sort"

	"github.com/tpfeiffer67/console/ui/property"
)

type Entities map[string]any

func (o Entities) GetEntityById(id string) (any, bool) {
	if e, ok := o[id]; ok {
		return e, true
	}
	return nil, false
}

func (o Entities) GetListOfChildren(id string) []any {
	l := o.listOfEntitiesWhoseParentHasTheGivenId(id)
	return l
}

func (o Entities) GetListOfBrotherhood(id string) []any {
	if parentId, ok := o.getParentEntityId(id); ok {
		l := o.listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(parentId)
		return l
	}
	return make([]any, 0)
}

func (o Entities) getParentEntityId(id string) (string, bool) {
	if a, ok := o.GetEntityById(id); ok {
		if e, ok := a.(property.IParent); ok {
			return e.Parent(), true
		}
	}
	return "", false
}

func (o Entities) listOfEntitiesWhoseParentHasTheGivenId(id string) []any {
	return o.filter(func(a any) bool {
		if e, ok := a.(property.IParent); ok {
			return (e.Parent() == id)
		}
		return false
	})
}

func (o Entities) listOfEntityAndHisAscendants(id string) []any {
	l := make([]any, 0)
	for {
		if a, ok := o.GetEntityById(id); ok {
			l = append(l, a)
			if e, ok := a.(property.IParent); ok {
				id = e.Parent()
				if id == "" {
					break
				}
			} else {
				break
			}
		}
	}
	return l
}

func (o Entities) listOfEntities_WhoseParentHasTheGivenId_AndWhichAreVisible(id string) []any {
	return o.filter(func(a any) bool {
		if e, ok := a.(property.IParentVisible); ok {
			return (e.Parent() == id && e.Visible())
		} else {
			return false
		}
	})
}

func (o Entities) CallFuncForEachChildrenEntity(id string, f func(any)) {
	l := o.listOfEntitiesWhoseParentHasTheGivenId(id)
	l = sortEntitiesByOrder(l)
	callFuncWithEachEntityOfTheList(l, f)
}

func (o Entities) CallFuncForEntityAndHisAscendants(id string, f func(any)) {
	l := o.listOfEntityAndHisAscendants(id)
	callFuncWithEachEntityOfTheList(l, f)
}

func (o Entities) GetFirstChildEntity(id string) (any, bool) {
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
	for _, a := range o {
		if e, ok := a.(property.IParent); ok {
			if e.Parent() == id {
				c++
			}
		}
	}
	return c
}

func (o Entities) CallFuncWithAllEntities(f func(any)) {
	for _, e := range o {
		f(e)
	}
}

/*
TODO remove

	func (o Entities) CallFuncWithAllEntitiesAsInterface(f func(i interface{})) {
		for _, i := range o {
			f(i)
		}
	}
*/
func callFuncWithEachEntityOfTheList(l []any, f func(any)) {
	for _, e := range l {
		f(e)
	}
}

func getEntityIndex(l []any, id string) (int, bool) {
	for i, a := range l {
		if a.(property.IId).Id() == id { // fatal error if not an IId  #directId
			return i, true
		}
	}
	return -1, false
}

func (o Entities) GetTheTopMostAncestorEntity(id string) any {
	for {
		if a, ok := o.GetEntityById(id); ok {
			if e, ok := a.(property.IParent); ok {
				id = e.Parent()
				if id == "" {
					return a
				}
			}
		}
	}
}

func (o Entities) CallFuncForEntityAndAllItsDescendants(id string, f func(any)) {
	if n, ok := o.GetEntityById(id); ok {
		o.callFuncForEntityAndAllItsDescendants(n, f)
	}
}

func (o Entities) callFuncForEntityAndAllItsDescendants(a any, f func(any)) {
	f(a)
	l := o.listOfEntitiesWhoseParentHasTheGivenId(a.(property.IId).Id()) // #directId
	for _, e := range l {
		o.callFuncForEntityAndAllItsDescendants(e, f)
	}
}

func (o Entities) listEntities_WhoseParentIsId(id string) []any {
	return o.filter(func(a any) bool {
		if e, ok := a.(property.IParent); ok {
			return (e.Parent() == id)
		}
		return false
	})
}

func (o Entities) listEntities_WhoseParentIsId_AndWhichAreVisible_AndWhichAreFocusable(id string) []any {
	return o.filter(func(a any) bool {
		if e, ok := a.(property.IParentVisibleFocusable); ok {
			return (e.Parent() == id && e.Visible() && e.Focusable())
		}
		return false
	})
}

func (o Entities) listEntities_WhoseParentIsId_AndWhichAreVisible_AndWithGivenZOrderLayer(id string, zlayer int) []any {
	return o.filter(func(a any) bool {
		if e, ok := a.(property.IParentVisibleZOrder); ok {
			return (e.Parent() == id && e.Visible() && e.ZOrderLayer() == zlayer)
		}
		return false
	})
}

func sortEntitiesByOrder(l []any) []any {
	sort.SliceStable(l, func(i, j int) bool {
		ai := l[i]
		aj := l[j]
		if ei, ok := ai.(property.IOrder); ok {
			if ej, ok := aj.(property.IOrder); ok {
				return ei.Order() < ej.Order()
			}
		}
		return false
	})
	return l
}

func sortEntitiesByZOrder(l []any) []any {
	sort.SliceStable(l, func(i, j int) bool {
		ai := l[i]
		aj := l[j]
		if ei, ok := ai.(property.IZOrder); ok {
			if ej, ok := aj.(property.IZOrder); ok {
				return ei.ZOrderSum() < ej.ZOrderSum()
			}
		}
		return false
	})
	return l
}

func sortEntitiesByFocusOrder(l []any) []any {
	sort.SliceStable(l, func(i, j int) bool {
		ai := l[i]
		aj := l[j]
		if ei, ok := ai.(property.IFocus); ok {
			if ej, ok := aj.(property.IFocus); ok {
				return ei.FocusOrder() < ej.FocusOrder()
			}
		}
		return false
	})
	return l
}

func sortEntitiesByReverseFocusOrder(l []any) []any {
	sort.SliceStable(l, func(i, j int) bool {
		ai := l[i]
		aj := l[j]
		if ei, ok := ai.(property.IFocus); ok {
			if ej, ok := aj.(property.IFocus); ok {
				return ei.FocusOrder() > ej.FocusOrder()
			}
		}
		return false
	})
	return l
}

func (o Entities) ZOrderUp(id string) bool {
	if a, ok := o.GetEntityById(id); ok {
		if e, ok := a.(property.IParentVisibleZOrder); ok {
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
	}
	return false
}

func (o Entities) ZOrderDown(id string) bool {
	if a, ok := o.GetEntityById(id); ok {
		if e, ok := a.(property.IParentVisibleZOrder); ok {
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
	}
	return false // true means ZOrder has not changed
}

func swapZOrder(a1, a2 any) {
	if e1, ok := a1.(property.IZOrder); ok {
		if e2, ok := a2.(property.IZOrder); ok {
			z1 := e1.ZOrder()
			z2 := e2.ZOrder()
			e1.SetZOrder(z2)
			e2.SetZOrder(z1)
		}
	}
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
	/* TODO revoir Theme
	f := func(i interface{}) {
		if e, ok := i.(ntt.IEntityWithTheme); ok {
			e.UpdateTheme(theme.CurrentTheme)
			e.DoOnThemeChange()
		}
	}
	o.callFuncWithAllEntities(f)
	*/
}

func (o Entities) filter(filter func(any) bool) []any {
	l := make([]any, 0)
	for _, e := range o {
		if filter(e) {
			l = append(l, e)
		}
	}
	return l
}
