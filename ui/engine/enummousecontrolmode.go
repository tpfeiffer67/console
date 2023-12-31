package engine

// This file was auto-generated by goenum.exe
// Do not modify this file directly.
// The name of the original description file used to generate this one is ./goenum/MouseControlMode.goenum
// The name of the template is ./goenum/goenum.template

import (
	"errors"
	"strings"
)

type MouseControlMode int

const (
	MouseControlModeCount = 3
	MouseControlModeMaxIndex = int(MouseControlModeResize)
	MouseControlModeLastValue = MouseControlModeResize
)

const (
	MouseControlModeDefault MouseControlMode = iota
	MouseControlModeMove
	MouseControlModeResize
)

func (v MouseControlMode) String() string {
	return [...]string{
		"MouseControlModeDefault",
		"MouseControlModeMove",
		"MouseControlModeResize",
	}[v]
}

func MouseControlModeFromString(s string) (MouseControlMode, error) {
	var suffix string
	if strings.HasPrefix(s, "MouseControlMode") {
		l := len("MouseControlMode")
		if l < len(s) {
			suffix = s[l:]
		}
	} else {
		suffix = s
	}
	switch suffix {
	case "Default":
		return MouseControlModeDefault, nil
	case "Move":
		return MouseControlModeMove, nil
	case "Resize":
		return MouseControlModeResize, nil
	}
	return MouseControlMode(0), errors.New("String does not correspond to any existing MouseControlMode values")
}
