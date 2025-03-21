package theme

import (
	"errors"

	"github.com/tpfeiffer67/console/screen"
)

type IValuesMap interface {
	GetMap() ValuesMap
	CopyFrom(ValuesMap, ...string)
	GetBool(string) (bool, bool)
	GetBoolDef(string, bool) (bool, bool)
	GetInt(string) (int, bool)
	GetIntDef(string, int) (int, bool)
	GetString(string) (string, bool)
	GetStringDef(string, string) (string, bool)
	GetUInt8(string) (uint8, bool)
	GetUInt8Def(string, uint8) (uint8, bool)
	GetChar(string) (rune, bool)
	GetCharDef(string, rune) (rune, bool)
	GetStyle(string) (screen.Style, bool)
	GetStyleDef(string, screen.Style) (screen.Style, bool)
	GetColor(string) (screen.Color, bool)
	GetColorDef(string, screen.Color) (screen.Color, bool)
	GetAny(string) (any, error)
}

type ValuesMap map[string]any

func (o ValuesMap) GetMap() ValuesMap {
	return o
}

func (o ValuesMap) CopyFrom(source ValuesMap, fields ...string) {
	for _, field := range fields {
		if v, ok := source[field]; ok {
			o[field] = v
		}
	}
}

func (o ValuesMap) GetBool(name string) (bool, bool) {
	return o.GetBoolDef(name, false)
}

func (o ValuesMap) GetBoolDef(name string, def bool) (bool, bool) {
	if v, ok := o[name].(bool); ok {
		return v, true
	}
	return false, false
}

func (o ValuesMap) GetInt(name string) (int, bool) {
	return o.GetIntDef(name, 0)
}

func (o ValuesMap) GetIntDef(name string, def int) (int, bool) {
	if v, ok := o[name].(int); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetString(name string) (string, bool) {
	return o.GetStringDef(name, "")
}

func (o ValuesMap) GetStringDef(name string, def string) (string, bool) {
	if v, ok := o[name].(string); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetUInt8(name string) (uint8, bool) {
	return o.GetUInt8Def(name, 0)
}

func (o ValuesMap) GetUInt8Def(name string, def uint8) (uint8, bool) {
	if v, ok := o[name].(uint8); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetChar(name string) (rune, bool) {
	return o.GetCharDef(name, ' ')
}

func (o ValuesMap) GetCharDef(name string, def rune) (rune, bool) {
	if v, ok := o[name].(rune); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetStyle(name string) (screen.Style, bool) {
	return o.GetStyleDef(name, screen.Style{})
}

func (o ValuesMap) GetStyleDef(name string, def screen.Style) (screen.Style, bool) {
	if v, ok := o[name].(screen.Style); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetColor(name string) (screen.Color, bool) {
	return o.GetColorDef(name, screen.Color{})
}

func (o ValuesMap) GetColorDef(name string, def screen.Color) (screen.Color, bool) {
	if v, ok := o[name].(screen.Color); ok {
		return v, true
	}
	return def, false
}

func (o ValuesMap) GetAny(name string) (any, error) {
	if a, ok := o[name]; ok {
		return a, nil
	}
	return nil, errors.New("value not found")
}
