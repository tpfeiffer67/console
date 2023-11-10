package screen

type EmphasisCanvas struct {
	Size
	DefaultValue Attributes
	buffer       []Attributes
	Canvas       [][]Attributes
}

func NewEmphasisCanvas(height, width int) *EmphasisCanvas {
	var o EmphasisCanvas
	o.Resize(height, width)
	o.Clear()
	return &o
}

func (o *EmphasisCanvas) SetHeight(height int) {
	if o.height != height {
		o.Resize(height, o.width)
	}
}

func (o *EmphasisCanvas) SetWidth(width int) {
	if o.width != width {
		o.Resize(o.height, width)
	}
}

func (o *EmphasisCanvas) Resize(height int, width int) {
	o.height = height
	o.width = width
	l := height * width
	o.buffer = make([]Attributes, l)
	o.Canvas = make([][]Attributes, height)
	for i := 0; i < height; i++ {
		o.Canvas[i] = o.buffer[i*width : i*width+width]
	}
}

func (o *EmphasisCanvas) Clear() {
	o.Fill(o.DefaultValue)
}

func (o *EmphasisCanvas) Fill(value Attributes) {
	for i := 0; i < o.width*o.height; i++ {
		o.buffer[i] = value
	}
}

func (o *EmphasisCanvas) Render(b *Buffer, stencil *Stencil, position Coordinates, offset Coordinates) {
	for i := 0; i < stencil.Height(); i++ {
		for j := 0; j < stencil.Width(); j++ {
			if stencil.IsNotMasked(i, j) {
				r1 := position.Row + i
				c1 := position.Col + j
				r2 := offset.Row + i
				c2 := offset.Col + j
				if r1 >= 0 && r1 < b.height && c1 >= 0 && c1 < b.width {
					var v Attributes
					if r2 >= 0 && r2 < o.height && c2 >= 0 && c2 < o.width {
						v = o.Canvas[r2][c2]
					} else {
						v = o.DefaultValue
					}
					b.PokeEmphasis(r1, c1, v)
				}
			}
		}
	}
}

func (o *EmphasisCanvas) SetEmphasis(row, col int, emphasis Attributes) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		o.Canvas[row][col] = emphasis
	}
}

func (o *EmphasisCanvas) GetEmphasis(row, col int) (Attributes, bool) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		return o.Canvas[row][col], true
	}
	return Attributes{}, false
}
