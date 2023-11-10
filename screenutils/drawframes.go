package screenutils

import "github.com/tpfeiffer67/console/screen"

func DrawColorFrame(o screen.ColorSetter, row, col, height, width int, color screen.Color) {
	for c := col; c < col+width; c++ {
		o.SetColor(row, c, color)
		o.SetColor(row+height-1, c, color)
	}
	for r := row + 1; r < row+height-1; r++ {
		o.SetColor(r, col, color)
		o.SetColor(r, col+width-1, color)
	}
}

func DrawColorGradientFrame(colorsetter screen.ColorSetter, row, col, height, width int, topLeft, topRight, bottomLeft, bottomRight screen.Color) {
	topGradients := BuildColorGradient(topLeft, topRight, width)
	bottomGradients := BuildColorGradient(bottomLeft, bottomRight, width)
	for i := 0; i < width; i++ {
		colorsetter.SetColor(row, col+i, topGradients[i])
		colorsetter.SetColor(row+height-1, col+i, bottomGradients[i])
	}
	leftGradients := BuildColorGradient(topLeft, bottomLeft, height)
	rightGradients := BuildColorGradient(topRight, bottomRight, height)
	//for i := 1; i < height-1; i++ {
	for i := height - 2; i > 0; i-- {
		colorsetter.SetColor(row+i, col, leftGradients[i])
		colorsetter.SetColor(row+i, col+width-1, rightGradients[i])
	}
}

func DrawRuneFrame(o screen.RuneSetter, row, col, height, width int, ru rune) {
	for c := col; c < col+width; c++ {
		o.SetRune(row, c, ru)
		o.SetRune(row+height-1, c, ru)
	}
	for r := row + 1; r < row+height-1; r++ {
		o.SetRune(r, col, ru)
		o.SetRune(r, col+width-1, ru)
	}
}

type BoxPattern struct {
	TopLeftCorner     rune
	TopBorder         rune
	TopRightCorner    rune
	LeftBorder        rune
	RightBorder       rune
	BottomLeftCorner  rune
	BottomBorder      rune
	BottomRightCorner rune
}

var boxPatterns = [14]BoxPattern{
	{'+', '-', '+', '|', '|', '+', '-', '+'},
	{'┌', '─', '┐', '│', '│', '└', '─', '┘'},
	{'╭', '─', '╮', '│', '│', '╰', '─', '╯'},
	{'┏', '━', '┓', '┃', '┃', '┗', '━', '┛'},
	{'╒', '═', '╕', '│', '│', '╘', '═', '╛'},
	{'╓', '─', '╖', '║', '║', '╙', '─', '╜'},
	{'╔', '═', '╗', '║', '║', '╚', '═', '╝'},
	{'▛', '▀', '▜', '▌', '▐', '▙', '▄', '▟'},
	{'▞', '▀', '▚', '▌', '▐', '▚', '▄', '▞'},
	{'█', '▀', '█', '█', '█', '█', '▄', '█'},
	{'░', '░', '░', '░', '░', '░', '░', '░'},
	{'▒', '▒', '▒', '▒', '▒', '▒', '▒', '▒'},
	{'▓', '▓', '▓', '▓', '▓', '▓', '▓', '▓'},
	{'█', '█', '█', '█', '█', '█', '█', '█'},
}

func DrawFrame(o screen.RuneSetter, row, col, height, width int, patternIndex int) {
	box := boxPatterns[patternIndex]

	o.SetRune(row+height-1, col, box.BottomLeftCorner)
	o.SetRune(row+height-1, col+width-1, box.BottomRightCorner)
	o.SetRune(row, col, box.TopLeftCorner)
	o.SetRune(row, col+width-1, box.TopRightCorner)

	for c := col + width - 2; c > col; c-- {
		o.SetRune(row, c, box.TopBorder)
		o.SetRune(row+height-1, c, box.BottomBorder)
	}
	for r := row + 1; r < row+height-1; r++ {
		o.SetRune(r, col, box.LeftBorder)
		o.SetRune(r, col+width-1, box.RightBorder)
	}
}
