package engine

// This file was auto-generated by goenum.exe
// Do not modify this file directly.
// The name of the original description file used to generate this one is ./goenum/EngineMode.goenum
// The name of the template is ./goenum/goenum.template

import (
	"errors"
	"strings"
)

type EngineMode int

const (
	EngineModeCount = 3
	EngineModeMaxIndex = int(EngineModeResize)
	EngineModeLastValue = EngineModeResize
)

const (
	EngineModeNormal EngineMode = iota
	EngineModeMove
	EngineModeResize
)

func (v EngineMode) String() string {
	return [...]string{
		"EngineModeNormal",
		"EngineModeMove",
		"EngineModeResize",
	}[v]
}

func EngineModeFromString(s string) (EngineMode, error) {
	var suffix string
	if strings.HasPrefix(s, "EngineMode") {
		l := len("EngineMode")
		if l < len(s) {
			suffix = s[l:]
		}
	} else {
		suffix = s
	}
	switch suffix {
	case "Normal":
		return EngineModeNormal, nil
	case "Move":
		return EngineModeMove, nil
	case "Resize":
		return EngineModeResize, nil
	}
	return EngineMode(0), errors.New("String does not correspond to any existing EngineMode values")
}
