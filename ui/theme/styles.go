package theme

import (
	"strings"

	"github.com/tpfeiffer67/console/screen"
)

// style_name,attributes,fcolor,bcolor
// attributes=ius => which means : i=italic, u=underline, s=strikethrough
// fcolor and bcolor : value form 0 to 255 or true color starting #
// Example: top_menu_bar_focused,,0,222
// TODO Add error return
// TODO Use StringToColor (value, name, index ...)
// TODO Create regex to validate format
func styleFromDescriptionString(description string) (string, screen.Style) {
	style := screen.Style{}

	fields := strings.Split(description, ",")
	if len(fields) != 4 {
		return "", style
	}
	if len(fields[0]) == 0 {
		return "", style
	}

	setStyleEmphasis(fields[1], &style)
	style.FColor = ToColor(fields[2])
	style.BColor = ToColor(fields[3])
	return fields[0], style
}

func setStyleEmphasis(s string, style *screen.Style) {
	if strings.Contains(s, "i") {
		style.Emphasis.Italic = true
	}
	if strings.Contains(s, "u") {
		style.Emphasis.Underline = true
	}
	if strings.Contains(s, "s") {
		style.Emphasis.Strikeout = true
	}
}
