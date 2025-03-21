package widget

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/property"
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

func ClearWithStyle(a any, vm theme.IValuesMap, normal, hovered, focused, focusedhovered string) screen.Style {
	style := GetStyleByStatus(a, vm, normal, hovered, focused, focusedhovered)
	if e, ok := a.(screenutils.ICombosCanvas); ok {
		e.SetDefaults(' ', style)
		e.Clear()
	}
	return style
}

type StatusStyler interface {
	property.IFocus
	property.IMouse
}

func GetStyleByStatus(a any, vm theme.IValuesMap, normal, hovered, focused, focusedhovered string) screen.Style {
	style := func(name string) screen.Style {
		st, _ := vm.GetStyleDef(name, screen.Style{})
		return st
	}

	if e, ok := a.(StatusStyler); ok {
		if e.FocusedGroup() {
			if e.HoveredGroup() {
				return style(focusedhovered)
			}
			if e.Focused() {
				s := style(focused)
				// TODO modify to define a custom style, not just inverting
				s.BColor, s.FColor = s.FColor, s.BColor
				return s
			}
			return style(focused)
		}
		if e.HoveredGroup() {
			return style(hovered)
		}
	}

	return style(normal)
}
