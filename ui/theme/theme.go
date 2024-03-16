package theme

import (
	"fmt"
	"strconv"

	"github.com/tpfeiffer67/console/screen"
)

type ITheme interface {
	IValuesMap
	//	UpdateTheme(ValuesMap)
	//	AppendValues(valuesId ...string)
	//	ValuesId() []string // TODO rename in ValueIds
}

type Theme struct {
	ITheme
	//valuesId []string
}

/*
func NewTheme(valuesId ...string) *Theme {
	o := new(Theme)
	o.valuesId = valuesId
	o.UpdateTheme(CurrentTheme)
	return o
}*/

// TODO add method to change theme
func NewTheme() *Theme {
	o := new(Theme)
	o.ITheme = CurrentTheme
	//	o.valuesId = valuesId
	//	o.UpdateTheme(CurrentTheme)
	return o
}

/*
	func (o *Theme) AppendValues(valuesId ...string) {
		o.valuesId = append(o.valuesId, valuesId...)
		o.UpdateTheme(CurrentTheme)
	}

	func (o *Theme) UpdateTheme(sourceTheme ValuesMap) {
		o.IValuesMap = make(ValuesMap)
		o.CopyFrom(sourceTheme, o.valuesId...)
	}

	func (o *Theme) ValuesId() []string {
		return o.valuesId
	}
*/
var CurrentTheme ValuesMap

func init() {
	SetThemeDefault()
	//SaveToFile("test.json")
}

func SetTheme(name string) error {
	switch name {
	case "default":
		SetThemeDefault()
	case "terminal":
		SetThemeTerminal()
	default:
		return fmt.Errorf("theme %s not found", name)
	}
	return nil
}

func setFuncToFindTheNameOfAColor(m map[string]interface{}) {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	screen.SetFuncToFindTheNameOfAColor(func(color screen.Color) (string, bool) {
		for _, k := range keys {
			if c, ok := CurrentTheme.GetColor(k); ok {
				if c == color {
					return k, true
				}
			}
		}
		return "", false
	})
}

func SetThemeDefault() {
	CurrentTheme = make(map[string]interface{})

	// https://coolors.co/gradient-palette/000f1c-415970?number=9
	// https://coolors.co/gradient-palette/415970-93cbff?number=8
	// primary: blue
	addColor("P0", "#000F1C")
	addColor("P1", "#081827")
	addColor("P2", "#102231")
	addColor("P3", "#182B3C")
	addColor("P4", "#213446")
	addColor("P5", "#293D51")
	addColor("P6", "#31475B")
	addColor("P7", "#395066")
	addColor("P8", "#415970")
	addColor("P9", "#4D6984")
	addColor("P10", "#587A99")
	addColor("P11", "#648AAD")
	addColor("P12", "#709AC2")
	addColor("P13", "#7CAAD6")
	addColor("P14", "#87BBEB")
	addColor("P15", "#93CBFF")

	// https://coolors.co/gradient-palette/100700-ff6600?number=9
	// https://coolors.co/gradient-palette/ff8800-663600?number=8
	// secondary: orange
	addColor("S0", "#100700")
	addColor("S1", "#2E1300")
	addColor("S2", "#4C1F00")
	addColor("S3", "#6A2B00")
	addColor("S4", "#883700")
	addColor("S5", "#A54200")
	addColor("S6", "#C34E00")
	addColor("S7", "#E15A00")
	addColor("S8", "#FF6600")
	addColor("S9", "#FF781E")
	addColor("S10", "#FF8B3D")
	addColor("S11", "#FF9D5B")
	addColor("S12", "#FFAF7A")
	addColor("S13", "#FFC198")
	addColor("S14", "#FFD4B7")
	addColor("S15", "#FFE6D5")

	// gray
	addColor("BL", "#000000")
	addColor("WH", "#FFFFFF")
	addColor("G0", "#000000")
	addColor("G1", "#111111")
	addColor("G2", "#222222")
	addColor("G3", "#333333")
	addColor("G4", "#444444")
	addColor("G5", "#555555")
	addColor("G6", "#666666")
	addColor("G7", "#777777")
	addColor("G8", "#888888")
	addColor("G9", "#999999")
	addColor("G10", "#AAAAAA")
	addColor("G11", "#BBBBBB")
	addColor("G12", "#CCCCCC")
	addColor("G13", "#DDDDDD")
	addColor("G14", "#EEEEEE")
	addColor("G15", "#FFFFFF")

	// gray 2
	// https://coolors.co/gradient-palette/001d23-bad9df?number=8
	addColor("g0", "#001D23")
	addColor("g1", "#1B383E")
	addColor("g2", "#355359")
	addColor("g3", "#506E74")
	addColor("g4", "#6A888E")
	addColor("g5", "#85A3A9")
	addColor("g6", "#9FBEC4")
	addColor("g7", "#BAD9DF")

	setFuncToFindTheNameOfAColor(CurrentTheme)

	addStyle(STYLE_BUTTON, ",P1,P4")
	addStyle(STYLE_BUTTON_HOVERED, ",P2,P9")
	addStyle(STYLE_BUTTON_FOCUSED, ",S2,S9")
	addStyle(STYLE_BUTTON_FOCUSEDHOVERED, ",S2,S12")
	addStyle(STYLE_BUTTON_DOWN, ",S12,S2")
	addStyle(STYLE_CHECKBOX, ",P0,P4")
	addStyle(STYLE_CHECKBOX_HOVERED, ",P0,P6")
	addStyle(STYLE_CHECKBOX_FOCUSED, ",S0,S4")
	addStyle(STYLE_CHECKBOX_FOCUSEDHOVERED, ",S0,S6")
	addStyle(STYLE_DESKTOP, ",P8,0")
	addStyle(STYLE_IMAGE_BACKGROUND, ",$1,TRANSPARENT")
	addStyle(STYLE_LABEL, ",#76b5c5,TRANSPARENT")

	addStyle(STYLE_MENUBAR, ",P2,P9")
	addStyle(STYLE_MENUBAR_HOVERED, ",P2,P9")
	addStyle(STYLE_MENUBAR_FOCUSED, ",P2,P9")
	addStyle(STYLE_MENUBAR_FOCUSEDHOVERED, ",P2,P9")
	addStyle(STYLE_MENUITEM, ",P2,P9")
	addStyle(STYLE_MENUITEM_HOVERED, ",P2,P9")
	addStyle(STYLE_MENUITEM_FOCUSED, ",P2,P9")
	addStyle(STYLE_MENUITEM_FOCUSEDHOVERED, ",P14,P2")
	addStyle(STYLE_MENUDROPDOWN, ",P2,P9")
	addStyle(STYLE_MENUDROPDOWN_HOVERED, ",P2,P9")
	addStyle(STYLE_MENUDROPDOWN_FOCUSED, ",P2,P9")
	addStyle(STYLE_MENUDROPDOWN_FOCUSEDHOVERED, ",P2,P9")

	addStyle(STYLE_MINIMIZINGBAR, ",#f4dc6280,#f4dc6280")
	addStyle(STYLE_PANEL, ",P0,P3")
	addStyle(STYLE_PANEL_HOVERED, ",P0,P7")
	addStyle(STYLE_PANEL_FOCUSED, ",S2,S5")
	addStyle(STYLE_PANEL_FOCUSEDHOVERED, ",S2,S10")
	addStyle(STYLE_STATUSBAR, ",P4,P7")
	addStyle(STYLE_TRACKBAR, ",P4,P7")
	addStyle(STYLE_TRACKBAR_HOVERED, ",P1,P5")
	addStyle(STYLE_TRACKBAR_FOCUSED, ",S5,S7")
	addStyle(STYLE_TRACKBAR_FOCUSEDHOVERED, ",S1,S6")

	CurrentTheme[BUTTON_FRAME] = 7
	CurrentTheme[CHECKBOX_NOTCHECKED] = "☐ "
	CurrentTheme[CHECKBOX_CHECKED] = "🗹 "
	CurrentTheme[DESKTOP_BACKGROUND] = '●'
	CurrentTheme[IMAGE_BACKGROUND] = '░'
	CurrentTheme[MENUBAR_FIRST_ITEM_POSITION] = 2
	CurrentTheme[MENUBAR_SPACE_BETWEEN_ITEMS] = 3
	CurrentTheme[MENUDROPDOWN_FRAME] = 7
	CurrentTheme[MINIMIZINGBAR_STICKER_WIDTH] = 30
	CurrentTheme[MINIMIZING_CLICK_WIDTH] = 10
	CurrentTheme[PANEL_FRAME] = 6
	CurrentTheme[STATUSBAR_ITEMS_SEPARATOR] = " | "
	CurrentTheme[TRACKBAR_CURSOR] = '▼'
	CurrentTheme[TRACKBAR_BACKGROUND] = '-'
	CurrentTheme[SHADOW_ENABLED] = true
	CurrentTheme[SHADOW_VALUE] = uint8(120)
	CurrentTheme[SHADOW_VERTICAL_OFFSET] = 1
	CurrentTheme[SHADOW_HORIZONTAL_OFFSET] = 2
}

func SetThemeTerminal() {
	CurrentTheme = make(map[string]interface{})

	// https://coolors.co/gradient-palette/001500-04ff00?number=9
	// https://coolors.co/gradient-palette/04ff00-c0ffbf?number=8
	// primary: green
	addColor("P0", "#001500")
	addColor("P1", "#013200")
	addColor("P2", "#015000")
	addColor("P3", "#026D00")
	addColor("P4", "#028A00")
	addColor("P5", "#03A700")
	addColor("P6", "#03C500")
	addColor("P7", "#04E200")
	addColor("P8", "#04FF00")
	addColor("P9", "#1FFF1B")
	addColor("P10", "#3AFF37")
	addColor("P11", "#55FF52")
	addColor("P12", "#6FFF6D")
	addColor("P13", "#8AFF88")
	addColor("P14", "#A5FFA4")
	addColor("P15", "#C0FFBF")

	addColor("S0", "#001500")
	addColor("S1", "#013200")
	addColor("S2", "#015000")
	addColor("S3", "#026D00")
	addColor("S4", "#028A00")
	addColor("S5", "#03A700")
	addColor("S6", "#03C500")
	addColor("S7", "#04E200")
	addColor("S8", "#04FF00")
	addColor("S9", "#1FFF1B")
	addColor("S10", "#3AFF37")
	addColor("S11", "#55FF52")
	addColor("S12", "#6FFF6D")
	addColor("S13", "#8AFF88")
	addColor("S14", "#A5FFA4")
	addColor("S15", "#C0FFBF")

	// gray
	addColor("BL", "#000000")
	addColor("WH", "#FFFFFF")
	addColor("G0", "#000000")
	addColor("G1", "#111111")
	addColor("G2", "#222222")
	addColor("G3", "#333333")
	addColor("G4", "#444444")
	addColor("G5", "#555555")
	addColor("G6", "#666666")
	addColor("G7", "#777777")
	addColor("G8", "#888888")
	addColor("G9", "#999999")
	addColor("G10", "#AAAAAA")
	addColor("G11", "#BBBBBB")
	addColor("G12", "#CCCCCC")
	addColor("G13", "#DDDDDD")
	addColor("G14", "#EEEEEE")
	addColor("G15", "#FFFFFF")

	// gray 2
	// https://coolors.co/gradient-palette/001d23-bad9df?number=8
	addColor("g0", "#001D23")
	addColor("g1", "#1B383E")
	addColor("g2", "#355359")
	addColor("g3", "#506E74")
	addColor("g4", "#6A888E")
	addColor("g5", "#85A3A9")
	addColor("g6", "#9FBEC4")
	addColor("g7", "#BAD9DF")

	addStyle(STYLE_BUTTON, ",P1,P0")
	addStyle(STYLE_BUTTON_HOVERED, ",P2,P0")
	addStyle(STYLE_BUTTON_FOCUSED, ",S2,S9")
	addStyle(STYLE_BUTTON_FOCUSEDHOVERED, ",S2,S12")
	addStyle(STYLE_BUTTON_DOWN, ",S12,S2")
	addStyle(STYLE_CHECKBOX, ",P0,P0")
	addStyle(STYLE_CHECKBOX_HOVERED, ",P0,P0")
	addStyle(STYLE_CHECKBOX_FOCUSED, ",S0,S4")
	addStyle(STYLE_CHECKBOX_FOCUSEDHOVERED, ",S0,S6")
	addStyle(STYLE_DESKTOP, ",#111122,0")
	addStyle(STYLE_IMAGE_BACKGROUND, ",$1,TRANSPARENT")
	addStyle(STYLE_LABEL, ",#76b5c5,TRANSPARENT")
	addStyle(STYLE_MENUBAR, ",P0,P4")
	addStyle(STYLE_MENUBAR_HOVERED, ",P0,P0")
	addStyle(STYLE_MENUBAR_FOCUSED, ",S0,S4")
	addStyle(STYLE_MENUBAR_FOCUSEDHOVERED, ",S0,S6")
	addStyle(STYLE_MENUITEM, ",P0,P0")
	addStyle(STYLE_MENUITEM_HOVERED, ",P0,P0")
	addStyle(STYLE_MENUITEM_FOCUSED, ",S0,S4")
	addStyle(STYLE_MENUITEM_FOCUSEDHOVERED, ",S0,S6")
	addStyle(STYLE_MENUDROPDOWN, ",P0,P4")
	addStyle(STYLE_MENUDROPDOWN_HOVERED, ",P0,P6")
	addStyle(STYLE_MENUDROPDOWN_FOCUSED, ",S0,S4")
	addStyle(STYLE_MENUDROPDOWN_FOCUSEDGROUP, ",S0,S4")
	addStyle(STYLE_MENUDROPDOWN_FOCUSEDHOVERED, ",S0,S6")
	addStyle(STYLE_PANEL, ",P0,P0")
	addStyle(STYLE_PANEL_HOVERED, ",P0,P0")
	addStyle(STYLE_PANEL_FOCUSED, ",S2,S5")
	addStyle(STYLE_PANEL_FOCUSEDHOVERED, ",S2,S10")
	addStyle(STYLE_STATUSBAR, ",P4,P7")
	addStyle(STYLE_TRACKBAR, ",P4,P7")
	addStyle(STYLE_TRACKBAR_HOVERED, ",P1,P5")
	addStyle(STYLE_TRACKBAR_FOCUSED, ",S5,S7")
	addStyle(STYLE_TRACKBAR_FOCUSEDHOVERED, ",S1,S6")

	CurrentTheme[BUTTON_FRAME] = 0
	CurrentTheme[CHECKBOX_NOTCHECKED] = "[ ]  "
	CurrentTheme[CHECKBOX_CHECKED] = "[x]  "
	CurrentTheme[DESKTOP_BACKGROUND] = '°'
	CurrentTheme[IMAGE_BACKGROUND] = '░'
	CurrentTheme[MENUBAR_FIRST_ITEM_POSITION] = 2
	CurrentTheme[MENUBAR_SPACE_BETWEEN_ITEMS] = 3
	CurrentTheme[MENUDROPDOWN_FRAME] = 7
	CurrentTheme[MINIMIZINGBAR_STICKER_WIDTH] = 30
	CurrentTheme[MINIMIZING_CLICK_WIDTH] = 10
	CurrentTheme[PANEL_FRAME] = 0
	CurrentTheme[STATUSBAR_ITEMS_SEPARATOR] = " : "
	CurrentTheme[TRACKBAR_CURSOR] = '^'
	CurrentTheme[TRACKBAR_BACKGROUND] = '.'
	CurrentTheme[SHADOW_ENABLED] = true
	CurrentTheme[SHADOW_VALUE] = uint8(120)
	CurrentTheme[SHADOW_VERTICAL_OFFSET] = 1
	CurrentTheme[SHADOW_HORIZONTAL_OFFSET] = 2
}

func addStyle(styleId, styleDescriptionString string) {
	s := styleId + "," + styleDescriptionString
	name, style := styleFromDescriptionString(s)
	CurrentTheme[name] = style
}

func addColor(key string, colorStr string) {
	CurrentTheme[key] = ToColor(colorStr)
}

/*
	ToColor converts a string to a color.

- TRANSPARENT
- $0 to $255 : console color
- #RGB
- #RGBA
- #RRGGBB
- #RRGGBBAA
- any other string
*/
func ToColorMap(s string, def screen.Color, colormap ValuesMap) screen.Color {
	if s == "TRANSPARENT" {
		return screen.Transparent
	}

	if s[:1] == "#" {
		colorNRGBA, _ := ParseHexColor(s)
		return screen.NewTrueColor(colorNRGBA)
	}

	if s[:1] == "$" {
		color256, _ := strconv.Atoi(s[1:])
		if color256 >= 0 && color256 < 256 {
			return screen.NewColor(color256)
		}
	}

	c, _ := colormap.GetColorDef(s, def)
	return c
}

func ToColor(s string) screen.Color {
	return ToColorMap(s, screen.Color{}, CurrentTheme)
}
