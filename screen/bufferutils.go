package screen

func (o *Buffer) FillRectangleWithFColor(row, col, height, width int, color Color) {
	for r := row; r < height; r++ {
		for c := col; c < width; c++ {
			o.PokeFColor(r, c, color)
		}
	}
}

func (o *Buffer) FillRectangleWithBColor(row, col, height, width int, color Color) {
	for r := row; r < height; r++ {
		for c := col; c < width; c++ {
			o.PokeBColor(r, c, color)
		}
	}
}
