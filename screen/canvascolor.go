package screen

import (
	"fmt"
)

// ColorCanvas is the canvas that contains the color buffer.
type ColorCanvas struct {
	Size
	RenderMode   int
	DefaultValue Color
	buffer       []Color
	Canvas       [][]Color
}

// FColorCanvas is a ColorCanvas for the foreground colors.
type FColorCanvas struct {
	ColorCanvas
}

// BColorCanvas is a ColorCanvas for the background colors.
type BColorCanvas struct {
	ColorCanvas
}

// NewFColorCanvas creates a new ColorCanvas.
func NewFColorCanvas(height, width int) *FColorCanvas {
	o := new(FColorCanvas)
	o.Resize(height, width)
	o.Clear()
	return o
}

func NewFColorCanvasWithDefaultValue(height, width int, defaultValue Color) *FColorCanvas {
	o := new(FColorCanvas)
	o.DefaultValue = defaultValue
	o.Resize(height, width)
	o.Clear()
	return o
}

func NewBColorCanvas(height, width int) *BColorCanvas {
	o := new(BColorCanvas)
	o.Resize(height, width)
	o.Clear()
	return o
}

func NewBColorCanvasWithDefaultValue(height, width int, defaultValue Color) *BColorCanvas {
	o := new(BColorCanvas)
	o.DefaultValue = defaultValue
	o.Resize(height, width)
	o.Clear()
	return o
}

func (o *ColorCanvas) SetHeight(height int) {
	if o.height != height {
		o.Resize(height, o.width)
	}
}

func (o *ColorCanvas) SetWidth(width int) {
	if o.width != width {
		o.Resize(o.height, width)
	}
}

func (o *ColorCanvas) Resize(height, width int) {
	o.height = height
	o.width = width
	l := height * width
	o.buffer = make([]Color, l)
	o.Canvas = make([][]Color, height)
	for i := 0; i < height; i++ {
		o.Canvas[i] = o.buffer[i*width : i*width+width]
	}
}

func (o *ColorCanvas) Clear() {
	o.Fill(o.DefaultValue)
}

func (o *ColorCanvas) Fill(value Color) {
	for i := 0; i < o.width*o.height; i++ {
		o.buffer[i] = value
	}
}

func (o *FColorCanvas) Render(b *Buffer, stencil *Stencil, position Coordinates, offset Coordinates) {
	o.render(b, stencil, position, offset, func(r1, c1, r2, c2 int) {
		c := o.Canvas[r2][c2]
		if c.Color256 != TRANSPARENT {
			switch o.RenderMode {
			case 1:
				b.PokeFColor(r1, c1, c)
			default:
				colorToBlendWith, _ := b.PeekFColor(r1, c1)
				blendedColor := BlendAlpha(c, colorToBlendWith, c)
				b.PokeFColor(r1, c1, blendedColor)
			}
		}
	})
}

func (o *BColorCanvas) Render(b *Buffer, stencil *Stencil, position Coordinates, offset Coordinates) {
	o.render(b, stencil, position, offset, func(r1, c1, r2, c2 int) {
		c := o.Canvas[r2][c2]
		if c.Color256 != TRANSPARENT {
			switch o.RenderMode {
			case 1:
				b.PokeBColor(r1, c1, c)
				colorToBlendWith, _ := b.PeekBColor(r1, c1)
				blendedColor := BlendAlpha(c, colorToBlendWith, c)
				b.PokeBColor(r1, c1, blendedColor)

				colorToBlendWith, _ = b.PeekFColor(r1, c1)
				blendedColor = BlendAlpha(c, colorToBlendWith, colorToBlendWith)
				b.PokeFColor(r1, c1, blendedColor)

			default:
				colorToBlendWith, _ := b.PeekBColor(r1, c1)
				blendedColor := BlendAlpha(c, colorToBlendWith, c)
				b.PokeBColor(r1, c1, blendedColor)
			}
		}
	})
}

func (o *ColorCanvas) render(b *Buffer, stencil *Stencil, position Coordinates, offset Coordinates, renderFunc func(int, int, int, int)) {
	for i := 0; i < stencil.Height(); i++ {
		for j := 0; j < stencil.Width(); j++ {
			if stencil.CellIsNotMasked(i, j) {
				r1 := position.Row + i
				c1 := position.Col + j
				r2 := offset.Row + i
				c2 := offset.Col + j
				if r1 >= 0 && r1 < b.height && c1 >= 0 && c1 < b.width { // inside the screen
					if r2 >= 0 && r2 < o.height && c2 >= 0 && c2 < o.width {
						renderFunc(r1, c1, r2, c2)
					}
				}
			}
		}
	}
}

func BlendAlpha(colorAlpha, colorToBlendWith, colorDefault Color) Color {
	if colorAlpha.IsTrueColor() {
		if colorAlpha.A == 255 {
			return colorAlpha
		}
		if colorAlpha.A == 0 { // completely transparent
			return colorToBlendWith
		}
		if !colorToBlendWith.IsTrueColor() {
			colorToBlendWith, _ = getTrueColorCorrespondenceForColor256(colorToBlendWith.Color256)
		}
		alpha := float64(colorAlpha.A) / 255.0
		alphacompl := 1.0 - alpha
		r := uint8((float64(colorAlpha.R) * alpha) + (float64(colorToBlendWith.R) * alphacompl))
		g := uint8((float64(colorAlpha.G) * alpha) + (float64(colorToBlendWith.G) * alphacompl))
		b := uint8((float64(colorAlpha.B) * alpha) + (float64(colorToBlendWith.B) * alphacompl))
		return NewTrueColorRGB(r, g, b)
	}
	// If the color is not type TrueColor, then no alpha blending.
	return colorDefault
}

func (o *ColorCanvas) SetColor(row, col int, color Color) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		o.Canvas[row][col] = color
	}
}

func (o *ColorCanvas) GetColor(row, col int) (Color, bool) {
	if verifyIfRowAndColAreInsideLimits(row, col, o.Height(), o.Width()) {
		return o.Canvas[row][col], true
	}
	return Color{}, false
}

func (o *ColorCanvas) String() string {
	return fmt.Sprintf("ColorCanvas,width=%d,height=%d,DefaultValue=%s", o.width, o.height, o.DefaultValue.String())
}
