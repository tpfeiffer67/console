package screenutils

import (
	"regexp"
	"strings"

	"github.com/tpfeiffer67/console/screen"
)

func DrawColorPattern(row, col int, colorPattern string, colorsetter screen.ColorSetter, decodeColor func(string) screen.Color) {
	lines := strings.Split(colorPattern, "\n")
	for _, s := range lines {
		c := col
		for i := 0; i < len(s); i++ {
			colorId := s[i : i+1]
			color := decodeColor(colorId)
			colorsetter.SetColor(row, c, color)
			c++
		}
		row++
	}
}

func DrawRawString(s string, row, col int, runesetter screen.RuneSetter) int {
	counter := 0
	runes := []rune(s)
	c := col
	for _, ru := range runes {
		if ru == '\n' {
			row++
			c = col
		} else {
			runesetter.SetRune(row, c, ru)
			c++
			counter++
		}
	}
	return counter
}

func DrawStyledString(row, col int, text string, canvas ICanvas, defaultStyle screen.Style, decodeColor func(string) screen.Color) int {
	runesetter, foreground, background, emphasiscanvas := canvas.GetRuneCanvas(), canvas.GetFColorCanvas(), canvas.GetBColorCanvas(), canvas.GetEmphasisCanvas()
	return drawStyledString(row, col, text, runesetter, foreground, background, emphasiscanvas, defaultStyle, decodeColor)
}

// The lines can contain formating markup
// <i> </i>	=> italic
// <u> </u>	=> underline
// <s> </s>	=> strikeout
// <fc=#ff1020> </fc>  forecolor
// <bc=#ff1020> </bc>  backcolor

// TODO Add multi-lines capabilities
// TODO Change color decoding
func drawStyledString(row, col int, text string, runesetter screen.RuneSetter, foreground screen.ColorSetter, background screen.ColorSetter, emphasiscanvas screen.EmphasisSetter, defaultStyle screen.Style, decodeColor func(string) screen.Color) int {
	const INDEX_OF_EQUAL_SIGN = 3 // Is the position of the equal sign in <fc= or <bc=
	var ru rune
	var markup string

	forecolor := defaultStyle.FColor
	backcolor := defaultStyle.BColor
	memoFColor := defaultStyle.FColor
	memoBColor := defaultStyle.BColor
	emphasis := defaultStyle.Emphasis
	runes := []rune(text)

	counter := 0
	index := 0
	c := col
	nl := func() { row++; c = col }

	for {
		ru, markup, index = getRuneOrMarkup(runes, index)

		if index == -1 {
			break
		}

		l := len(markup)
		if l > 0 {
			if l > INDEX_OF_EQUAL_SIGN {
				if markup[INDEX_OF_EQUAL_SIGN] == '=' {
					// TODO Factorize blocs below
					// TODO Add errors management
					s := markup[INDEX_OF_EQUAL_SIGN+1 : l-1]
					if markup[:INDEX_OF_EQUAL_SIGN] == "<fc" {
						memoFColor = forecolor
						forecolor = decodeColor(s)
					}
					if markup[:INDEX_OF_EQUAL_SIGN] == "<bc" {
						memoBColor = backcolor
						backcolor = decodeColor(s)
					}
				}
			}
			switch markup {
			case "<br>":
				nl()
			case "<i>":
				emphasis.Italic = true
			case "</i>":
				emphasis.Italic = false
			case "<u>":
				emphasis.Underline = true
			case "</u>":
				emphasis.Underline = false
			case "<s>":
				emphasis.Strikeout = true
			case "</s>":
				emphasis.Strikeout = false
			case "</fc>":
				forecolor = defaultStyle.FColor
			case "</bc>":
				backcolor = defaultStyle.BColor
			case "</c>":
				forecolor = defaultStyle.FColor
				backcolor = defaultStyle.BColor
			case "</mfc>":
				forecolor = memoFColor
			case "</mbc>":
				backcolor = memoBColor
			case "</mc>":
				forecolor = memoFColor
				backcolor = memoBColor
			}
		} else if ru == '\n' {
			nl()
		} else {
			runesetter.SetRune(row, c, ru)
			// If a color is TRANSPARENT then we do not set the value
			// It is possible to paint transparency by setting the alpha value to 0 (#0000)
			if !forecolor.IsTransparent() {
				foreground.SetColor(row, c, forecolor)
			}
			if !backcolor.IsTransparent() {
				background.SetColor(row, c, backcolor)
			}
			emphasiscanvas.SetEmphasis(row, c, emphasis)
			c++
			counter++
		}
	}
	return counter
}

// TODO Add Tests
func getRuneOrMarkup(runes []rune, pos int) (rune, string, int) {
	if pos < 0 || pos >= len(runes) {
		return rune(0), "", -1
	}
	ru := runes[pos]
	if ru == '<' {
		markup := getMarkup(runes, pos)
		if isValidMarkup(markup) {
			return rune(0), markup, pos + len(markup)
		}
	}
	// not a markup
	return ru, "", pos + 1
}

func getMarkup(runes []rune, pos int) string {
	var sb strings.Builder
	l := len(runes)
	for i := min(pos, l-1); i < l; i++ {
		ru := runes[i]
		sb.WriteRune(ru)
		if ru == '>' {
			return sb.String()
		}
	}
	return sb.String()
}

// TODO modify for the color fields
func isValidMarkup(markup string) bool {
	re := regexp.MustCompile(`^<br>$|^<i>$|^</i>$|^<u>$|^</u>$|^<s>$|^</s>$|^(<fc=|<bc=).*>$|^</fc>$|^</bc>$|^</c>$|^</mfc>$|^</mbc>$|^</mc>$`)
	return re.MatchString(markup)
}

func GetStyledStringLen(text string) int {
	var ru rune
	runes := []rune(text)
	index := 0
	result := 0
	for {
		ru, _, index = getRuneOrMarkup(runes, index)
		if index == -1 {
			break
		}
		if ru != rune(0) {
			result++
		}
	}
	return result
}
