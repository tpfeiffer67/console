package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/theme"
)

/*
func SetDefaultFuncFor_OnFocus_And_OnLostFocus(o IEntity) {
	o.SetOnFocus(func(foc property.IFocus) {
		o.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), true)
	})

	o.SetOnLostFocus(func(foc property.IFocus) {
		o.SetFocusedGroupFromTheTopMostAncestorEntity(o.Id(), false)
	})
}
*/

// TODO Combine IEntity and IValuesMap
func ForEntity_GetStyleByItsStatus_AndClear(o IEntity, vm theme.IValuesMap, normal, hovered, focused, focusedhovered string) screen.Style {
	style := ForEntity_GetStyleByItsStatus(o, vm, normal, hovered, focused, focusedhovered)
	o.SetDefaults(' ', style)
	o.Clear()
	return style
}

func ForEntity_GetStyleByItsStatus(o IEntity, vm theme.IValuesMap, normal, hovered, focused, focusedhovered string) screen.Style {
	style := func(name string) screen.Style {
		st, _ := vm.GetStyleDef(name, screen.Style{})
		return st
	}

	if o.FocusedGroup() {
		if o.HoveredGroup() {
			return style(focusedhovered)
		}
		if o.Focused() {
			s := style(focused)
			// TODO modify
			s.BColor, s.FColor = s.FColor, s.BColor
			return s
		}
		return style(focused)
	}
	if o.HoveredGroup() {
		return style(hovered)
	}
	return style(normal)
}
