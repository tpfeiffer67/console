package render

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/terminal"
)

func ScreenBufferToTerminal(o *screen.Buffer) int {
	counter := 0
	escseq := new(terminal.EscSeq)
	for row := 0; row < o.Height(); row++ {
		for col := 0; col < o.Width(); col++ {
			cell, previousCell := o.PeekCellCurrentAndPreviousValue(row, col)
			if cell != previousCell {
				// TODO Optimize
				// It is possible to optimize if several successive runes have same F and B colors
				drawCellAtPosition(o.Row+row+1, o.Col+col+1, cell, escseq)
			}
		}
		n, _ := escseq.Output()
		counter = counter + n
	}
	o.Swap()
	return counter
}

func ScreenBufferToTerminalForced(o *screen.Buffer) int {
	counter := 0
	escseq := new(terminal.EscSeq)
	for row := 0; row < o.Height(); row++ {
		for col := 0; col < o.Width(); col++ {
			cell, _ := o.PeekCell(row, col)
			drawCellAtPosition(o.Row+row+1, o.Col+col+1, cell, escseq)
		}
		n, _ := escseq.Output()
		counter = counter + n
	}
	o.Swap()
	return counter
}

func drawCellAtPosition(row, col int, cell screen.Cell, escseq *terminal.EscSeq) {
	escseq.SetCursorPosition(row, col)
	escseq.ClearStyle() // affects style (emphasis and colors)
	if cell.FColor.IsTrueColor() {
		escseq.SetForegroundTrueColor(cell.FColor.NRGBA)
	} else {
		escseq.SetForegroundColor(cell.FColor.Color256)
	}
	if cell.BColor.IsTrueColor() {
		escseq.SetBackgroundTrueColor(cell.BColor.NRGBA)
	} else {
		escseq.SetBackgroundColor(cell.BColor.Color256)
	}
	escseq.SetAttributes(cell.Emphasis.Italic, cell.Emphasis.Underline, cell.Emphasis.Strikeout)
	escseq.AddRune(cell.Rune)
}
