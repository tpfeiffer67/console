package ntt

import (
	"math"

	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type HSLAPicker struct {
	IWidget
	huePos      int
	satPos      int
	lumPos      int
	alpha       int
	selector    int
	widthInChar int
	width2      int
	widthf2     float64
	property.OnChange
}

func (o *HSLAPicker) Color() screen.Color {
	return screen.NewTrueColorRGBA(o.RGBA())
}

func (o *HSLAPicker) RGBA() (uint8, uint8, uint8, uint8) {
	div := o.widthf2 - 1.0
	h := float64(o.huePos) / div
	s := float64(o.satPos) / div
	l := float64(o.lumPos) / div
	r, g, b := hslToRgb(h, s, l)
	a := uint8(math.Round(255 * float64(o.alpha) / div))
	return r, g, b, a
}

func (o *HSLAPicker) HSLA() (uint8, uint8, uint8, uint8) {
	div := o.widthf2 - 1.0
	h := uint8(math.Round(255 * float64(o.huePos) / div))
	s := uint8(math.Round(255 * float64(o.satPos) / div))
	l := uint8(math.Round(255 * float64(o.lumPos) / div))
	a := uint8(math.Round(255 * float64(o.alpha) / div))
	return h, s, l, a
}

func NewHSLAPicker(id string, row, col int, widthInChar int, syst ISystem) *HSLAPicker {
	o := newHSLAPicker(widthInChar)
	o.IWidget = NewWidget(id, 9, widthInChar+2, syst)
	o.SetPosition(row, col)
	o.SetCanMove(false)
	o.SetFocusable(true)

	runeSelectorLeft := '🞂'
	runeSelectorRight := '🞀'
	runePickerLeft := '⠇'
	runePickerRight := '⠸'

	style := screen.Style{FColor: theme.ToColor("#FFF"), BColor: theme.ToColor("#000")}
	o.SetDefaults(' ', style)

	runeCanvas := o.GetRuneCanvas()
	fColorCanvas := o.GetFColorCanvas()
	bColorCanvas := o.GetBColorCanvas()
	top := 0
	o.SetOnDraw(func() {
		o.Clear()
		for i := 1; i <= o.widthInChar; i++ {
			runeCanvas.SetRune(top+1, i, '▌')
			runeCanvas.SetRune(top+3, i, '▌')
			runeCanvas.SetRune(top+5, i, '▌')
			runeCanvas.SetRune(top+7, i, '▚')
		}

		A := screen.ColorSetter(fColorCanvas)
		B := screen.ColorSetter(bColorCanvas)
		setColor := func(ro, co int, h, s, l float64) {
			r, g, b := hslToRgb(h, s, l)
			c := screen.NewTrueColorRGB(r, g, b)
			A.SetColor(ro, co/2+1, c)
		}

		hue := float64(o.huePos) / o.widthf2
		for i := 0; i < o.width2; i++ {
			f := float64(i) / o.widthf2
			setColor(top+1, i, f, 1.0, 0.5) // hue bar
			setColor(top+3, i, hue, f, 0.5) // saturation bar
			setColor(top+5, i, hue, 1.0, f) // luminosity bar
			A, B = B, A
		}

		black := screen.NewTrueColorRGB(0, 0, 0)
		white := screen.NewTrueColorRGB(255, 255, 255)
		clr := o.Color()
		for i := 0; i < o.widthInChar; i++ {
			clr.A = uint8(math.Round(255 * float64(i) / float64(o.widthInChar)))
			aclr := screen.BlendAlpha(clr, black, clr)
			bColorCanvas.SetColor(top+7, i+1, aclr)
			aclr = screen.BlendAlpha(clr, white, clr)
			fColorCanvas.SetColor(top+7, i+1, aclr)
		}

		runeCanvas.SetRune(top+1+o.selector*2, 0, runeSelectorLeft)
		runeCanvas.SetRune(top+1+o.selector*2, o.widthInChar+1, runeSelectorRight)

		drawPickerPosition := func(r int, position int) {
			runeCanvas.SetRune(top+r, position/2+1, iifRune(position%2 == 0, runePickerLeft, runePickerRight))
		}
		drawPickerPosition(2, o.huePos)
		drawPickerPosition(4, o.satPos)
		drawPickerPosition(6, o.lumPos)
		drawPickerPosition(8, o.alpha)

		colorstring := o.Color().String()
		screenutils.DrawStyledString(0, 1, colorstring, o, style, theme.ToColor)
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		o.SendMessage(message.MessageIdFocus, nil, id)
		return true
	})

	selectPicker := func() *int {
		switch o.selector {
		default: //0
			return &o.huePos
		case 1:
			return &o.satPos
		case 2:
			return &o.lumPos
		case 3:
			return &o.alpha
		}
	}
	pick := selectPicker()

	o.SetListener(message.MessageIdKey, func(messageParams interface{}) bool {
		h, s, l, a := o.huePos, o.satPos, o.lumPos, o.alpha

		if params, ok := messageParams.(message.ParamsKey); ok {
			switch params.Name {
			case "Up":
				o.selector = dec(o.selector, 0, 3)
				pick = selectPicker()
			case "Down":
				o.selector = inc(o.selector, 0, 3)
				pick = selectPicker()
			case "H", "h":
				o.selector = 0
				pick = selectPicker()
			case "S", "s":
				o.selector = 1
				pick = selectPicker()
			case "L", "l":
				o.selector = 2
				pick = selectPicker()
			case "A", "a":
				o.selector = 3
				pick = selectPicker()
			case "Left":
				*pick = dec(*pick, 0, o.width2-1)
			case "Right":
				*pick = inc(*pick, 0, o.width2-1)
			case "Home":
				*pick = 0
			case "End":
				*pick = o.width2 - 1
			}
			if params.Rune >= '0' && params.Rune <= '9' {
				v := float64(params.Rune - '0')
				step := float64(o.width2) / 10.0
				*pick = int(math.Round(v * step))
			}
		}

		if h != o.huePos || s != o.satPos || l != o.lumPos || a != o.alpha {
			o.DoOnChange()
		}

		return false
	})

	return o
}

func newHSLAPicker(w int) *HSLAPicker {
	o := new(HSLAPicker)
	o.widthInChar = w
	o.width2 = w * 2
	o.widthf2 = float64(w) * 2.0
	o.huePos = 0
	o.satPos = o.width2 - 1
	o.lumPos = w
	o.alpha = o.width2 - 1
	return o
}

// [Go Playground - The Go Programming Language](https://go.dev/play/p/9q5yBNDh3W)
func hueToRgb(v1, v2, vH float64) float64 {
	if vH < 0 {
		vH += 1
	}
	if vH > 1 {
		vH -= 1
	}
	if (6 * vH) < 1 {
		return (v1 + (v2-v1)*6*vH)
	}
	if (2 * vH) < 1 {
		return v2
	}
	if (3 * vH) < 2 {
		return (v1 + (v2-v1)*((2.0/3.0)-vH)*6)
	}
	return v1
}

func hslToRgb(H, S, L float64) (uint8, uint8, uint8) {
	var r, g, b float64
	if S == 0 { //HSL from 0 to 1
		r = L * 255 //RGB results from 0 to 255
		g = L * 255
		b = L * 255
	} else {
		var v1, v2 float64
		if L < 0.5 {
			v2 = L * (1 + S)
		} else {
			v2 = (L + S) - (S * L)
		}

		v1 = 2*L - v2

		r = 255 * hueToRgb(v1, v2, H+(1.0/3.0))
		g = 255 * hueToRgb(v1, v2, H)
		b = 255 * hueToRgb(v1, v2, H-(1.0/3.0))
	}
	return uint8(math.Round(r)), uint8(math.Round(g)), uint8(math.Round(b))
}

// IifRune returns trueRune if the boolean is true, falseRune if not.
func iifRune(b bool, trueRune, falseRune rune) rune {
	if b {
		return trueRune
	}
	return falseRune
}

// Dec decrements an integer and return its new value.
// If the value goes below the min limit, the max value is returned.
func dec(v, min, max int) int {
	v--
	if v < min {
		return max
	}
	return v
}

// Inc decrements an integer and return its new value.
// If the value goes over the max limit, the min value is returned.
func inc(v int, min, max int) int {
	v++
	if v > max {
		return min
	}
	return v
}
