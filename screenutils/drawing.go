package screenutils

import "github.com/tpfeiffer67/console/screen"

func FillRectangle(runes screen.RuneSetter, foreground screen.ColorSetter, background screen.ColorSetter, row, col, height, width int, ru rune, fcolor screen.Color, bcolor screen.Color) {
	for r := row; r < height; r++ {
		for c := col; c < width; c++ {
			runes.SetRune(r, c, ru)
			foreground.SetColor(r, c, fcolor)
			background.SetColor(r, c, bcolor)
		}
	}
}

func DrawHorizontalLine(runes screen.RuneSetter, foreground screen.ColorSetter, background screen.ColorSetter, row, col, width int, ru rune, fcolor screen.Color, bcolor screen.Color) {
	for c := col; c < width; c++ {
		runes.SetRune(row, c, ru)
		foreground.SetColor(row, c, fcolor)
		background.SetColor(row, c, bcolor)
	}
}

// DrawShadow paint a shadow directly to the screenBuffer
func DrawShadow(screenbuffer *screen.Buffer, stencil *screen.Stencil, position screen.Coordinates, rowOffset, colOffset int, shadowDensity uint8) {
	shadowPosition := screen.Coordinates{position.Row + rowOffset, position.Col + colOffset}
	fShadow := screen.NewFColorCanvas(stencil.Height(), stencil.Width())
	bShadow := screen.NewBColorCanvas(stencil.Height(), stencil.Width())
	fShadow.DefaultValue = screen.NewTrueColorRGBA(0, 0, 0, shadowDensity)
	bShadow.DefaultValue = screen.NewTrueColorRGBA(0, 0, 0, shadowDensity)
	fShadow.Clear()
	bShadow.Clear()
	fShadow.RenderMode = 3
	bShadow.RenderMode = 3
	fShadow.Render(screenbuffer, stencil, shadowPosition, screen.Coordinates{})
	bShadow.Render(screenbuffer, stencil, shadowPosition, screen.Coordinates{})
}

func DrawColorGradientsMatrix(colorcanvas screen.ColorSetter, row, col, height, width int, topLeft, topRight, bottomLeft, bottomRight screen.Color) {
	colorGradientsMatrix := BuildColorGradientMatrix(topLeft, topRight, bottomLeft, bottomRight, height, width)
	for r := row; r < height; r++ {
		for c := col; c < width; c++ {
			colorcanvas.SetColor(r, c, colorGradientsMatrix[r][c])
		}
	}
}
