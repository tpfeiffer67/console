package terminal

import (
	"fmt"
	"image/color"
	"os"
	"strings"
)

// ANSI Escape Sequences
// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

type EscSeq struct {
	sequence strings.Builder
}

// Usage EscSec.CursorInvisible().Output()

func (o *EscSeq) Output() (int, error) {
	n, err := os.Stdout.WriteString(o.sequence.String())
	o.sequence.Reset()
	return n, err
}

func (o *EscSeq) Reset() *EscSeq {
	o.sequence.Reset()
	return o
}

func (o *EscSeq) CursorInvisible() *EscSeq {
	o.sequence.WriteString("\033[?25l")
	return o
}

func (o *EscSeq) CursorVisible() *EscSeq {
	o.sequence.WriteString("\033[?25h")
	return o
}

func (o *EscSeq) CursorSave() *EscSeq {
	o.sequence.WriteString("\033[s")
	return o
}

func (o *EscSeq) CursorRestore() *EscSeq {
	o.sequence.WriteString("\033[u")
	return o
}

func (o *EscSeq) Clear() *EscSeq {
	o.sequence.WriteString("\033[2J")
	return o
}

func (o *EscSeq) SetMode(mode int) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[%dm", mode)
	return o
}

func (o *EscSeq) SetForegroundColor(color int) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[38;5;%dm", color)
	return o
}

func (o *EscSeq) SetBackgroundColor(color int) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[48;5;%dm", color)
	return o
}

func (o *EscSeq) SetForegroundTrueColor(c color.NRGBA) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
	return o
}

func (o *EscSeq) SetBackgroundTrueColor(c color.NRGBA) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[48;2;%d;%d;%dm", c.R, c.G, c.B)
	return o
}

func (o *EscSeq) SetCursorPosition(row int, col int) *EscSeq {
	fmt.Fprintf(&o.sequence, "\033[%d;%dH", row, col)
	return o
}

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#colors--graphics-mode
func (o *EscSeq) SetAttributes(italic bool, underline bool, strikeout bool) *EscSeq {
	if italic {
		o.sequence.WriteString("\033[3m")
	}
	if underline {
		o.sequence.WriteString("\033[4m")
	}
	if strikeout {
		o.sequence.WriteString("\033[9m")
	}
	return o
}

func (o *EscSeq) ClearStyle() *EscSeq {
	o.sequence.WriteString("\033[0m")
	return o
}

func (o *EscSeq) AddString(s string) *EscSeq {
	o.sequence.WriteString(s)
	return o
}

func (o *EscSeq) AddRune(r rune) *EscSeq {
	o.sequence.WriteRune(r)
	return o
}
