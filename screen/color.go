package screen

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
)

const TRUE_COLOR = 256
const TRANSPARENT = 257

var Transparent = Color{TRANSPARENT, color.NRGBA{}}

type Color struct {
	Color256 int
	color.NRGBA
}

func NewColor(color256 int) Color {
	return Color{color256, color.NRGBA{}}
}

func NewTrueColor(truecolor color.NRGBA) Color {
	return Color{TRUE_COLOR, truecolor}
}

func NewTrueColorRGB(r, g, b uint8) Color {
	return NewTrueColor(color.NRGBA{r, g, b, 255})
}

func NewTrueColorRGBA(r, g, b, a uint8) Color {
	return NewTrueColor(color.NRGBA{r, g, b, a})
}

func (o *Color) SetColor(color256 int) {
	o.Color256 = color256
}

func (o *Color) SetTrueColor(truecolor color.NRGBA) {
	o.Color256 = TRUE_COLOR
	o.NRGBA = truecolor
}

func (o *Color) IsTrueColor() bool {
	return o.Color256 == TRUE_COLOR
}

func (o *Color) SetTransparent() {
	o.Color256 = TRANSPARENT
}

func (o *Color) IsTransparent() bool {
	return o.Color256 == TRANSPARENT
}

func (o *Color) IsColor256() bool {
	if o.Color256 >= 0 && o.Color256 < 256 {
		return true
	}
	return false
}

func (o *Color) GetRGB() (uint8, uint8, uint8, bool) {
	if o.IsTrueColor() {
		return o.R, o.G, o.B, true
	}
	if o.IsColor256() {
		color, err := getTrueColorCorrespondenceForColor256(o.Color256)
		if err == nil {
			return color.R, color.G, color.B, true
		}
		return 0, 0, 0, false
	}
	return 0, 0, 0, false
}

func (o *Color) GetRGBfloat64() (float64, float64, float64, bool) {
	r, g, b, ok := o.GetRGB()
	if ok {
		factor := 1.0 / 255.0
		return float64(r) * factor, float64(g) * factor, float64(b) * factor, true
	}
	return float64(0), float64(0), float64(0), false
}

func getTrueColorCorrespondenceForColor256(color256 int) (Color, error) {
	if color256 >= 0 && color256 < 256 {
		return CorrespondanceMapANSI256ColorsToTrueColors[color256], nil
	}
	return Color{}, errors.New("invalid ANSI256 color index")
}

func init() {
	CorrespondanceMapANSI256ColorsToTrueColors = defaultCorrespondenceMap()
}

var CorrespondanceMapANSI256ColorsToTrueColors map[int]Color

func defaultCorrespondenceMap() map[int]Color {
	return map[int]Color{0: Color{TRUE_COLOR, color.NRGBA{0, 0, 0, 255}},
		1:   Color{TRUE_COLOR, color.NRGBA{128, 0, 0, 255}},
		2:   Color{TRUE_COLOR, color.NRGBA{0, 128, 0, 255}},
		3:   Color{TRUE_COLOR, color.NRGBA{128, 128, 0, 255}},
		4:   Color{TRUE_COLOR, color.NRGBA{0, 0, 128, 255}},
		5:   Color{TRUE_COLOR, color.NRGBA{128, 0, 128, 255}},
		6:   Color{TRUE_COLOR, color.NRGBA{0, 128, 128, 255}},
		7:   Color{TRUE_COLOR, color.NRGBA{192, 192, 192, 255}},
		8:   Color{TRUE_COLOR, color.NRGBA{128, 128, 128, 255}},
		9:   Color{TRUE_COLOR, color.NRGBA{255, 0, 0, 255}},
		10:  Color{TRUE_COLOR, color.NRGBA{0, 255, 0, 255}},
		11:  Color{TRUE_COLOR, color.NRGBA{255, 255, 0, 255}},
		12:  Color{TRUE_COLOR, color.NRGBA{0, 0, 255, 255}},
		13:  Color{TRUE_COLOR, color.NRGBA{255, 0, 255, 255}},
		14:  Color{TRUE_COLOR, color.NRGBA{0, 255, 255, 255}},
		15:  Color{TRUE_COLOR, color.NRGBA{255, 255, 255, 255}},
		16:  Color{TRUE_COLOR, color.NRGBA{0, 0, 0, 255}},
		17:  Color{TRUE_COLOR, color.NRGBA{0, 0, 95, 255}},
		18:  Color{TRUE_COLOR, color.NRGBA{0, 0, 135, 255}},
		19:  Color{TRUE_COLOR, color.NRGBA{0, 0, 175, 255}},
		20:  Color{TRUE_COLOR, color.NRGBA{0, 0, 215, 255}},
		21:  Color{TRUE_COLOR, color.NRGBA{0, 0, 255, 255}},
		22:  Color{TRUE_COLOR, color.NRGBA{0, 95, 0, 255}},
		23:  Color{TRUE_COLOR, color.NRGBA{0, 95, 95, 255}},
		24:  Color{TRUE_COLOR, color.NRGBA{0, 95, 135, 255}},
		25:  Color{TRUE_COLOR, color.NRGBA{0, 95, 175, 255}},
		26:  Color{TRUE_COLOR, color.NRGBA{0, 95, 215, 255}},
		27:  Color{TRUE_COLOR, color.NRGBA{0, 95, 255, 255}},
		28:  Color{TRUE_COLOR, color.NRGBA{0, 135, 0, 255}},
		29:  Color{TRUE_COLOR, color.NRGBA{0, 135, 95, 255}},
		30:  Color{TRUE_COLOR, color.NRGBA{0, 135, 135, 255}},
		31:  Color{TRUE_COLOR, color.NRGBA{0, 135, 175, 255}},
		32:  Color{TRUE_COLOR, color.NRGBA{0, 135, 215, 255}},
		33:  Color{TRUE_COLOR, color.NRGBA{0, 135, 255, 255}},
		34:  Color{TRUE_COLOR, color.NRGBA{0, 175, 0, 255}},
		35:  Color{TRUE_COLOR, color.NRGBA{0, 175, 95, 255}},
		36:  Color{TRUE_COLOR, color.NRGBA{0, 175, 135, 255}},
		37:  Color{TRUE_COLOR, color.NRGBA{0, 175, 175, 255}},
		38:  Color{TRUE_COLOR, color.NRGBA{0, 175, 215, 255}},
		39:  Color{TRUE_COLOR, color.NRGBA{0, 175, 255, 255}},
		40:  Color{TRUE_COLOR, color.NRGBA{0, 215, 0, 255}},
		41:  Color{TRUE_COLOR, color.NRGBA{0, 215, 95, 255}},
		42:  Color{TRUE_COLOR, color.NRGBA{0, 215, 135, 255}},
		43:  Color{TRUE_COLOR, color.NRGBA{0, 215, 175, 255}},
		44:  Color{TRUE_COLOR, color.NRGBA{0, 215, 215, 255}},
		45:  Color{TRUE_COLOR, color.NRGBA{0, 215, 255, 255}},
		46:  Color{TRUE_COLOR, color.NRGBA{0, 255, 0, 255}},
		47:  Color{TRUE_COLOR, color.NRGBA{0, 255, 95, 255}},
		48:  Color{TRUE_COLOR, color.NRGBA{0, 255, 135, 255}},
		49:  Color{TRUE_COLOR, color.NRGBA{0, 255, 175, 255}},
		50:  Color{TRUE_COLOR, color.NRGBA{0, 255, 215, 255}},
		51:  Color{TRUE_COLOR, color.NRGBA{0, 255, 255, 255}},
		52:  Color{TRUE_COLOR, color.NRGBA{95, 0, 0, 255}},
		53:  Color{TRUE_COLOR, color.NRGBA{95, 0, 95, 255}},
		54:  Color{TRUE_COLOR, color.NRGBA{95, 0, 135, 255}},
		55:  Color{TRUE_COLOR, color.NRGBA{95, 0, 175, 255}},
		56:  Color{TRUE_COLOR, color.NRGBA{95, 0, 215, 255}},
		57:  Color{TRUE_COLOR, color.NRGBA{95, 0, 255, 255}},
		58:  Color{TRUE_COLOR, color.NRGBA{95, 95, 0, 255}},
		59:  Color{TRUE_COLOR, color.NRGBA{95, 95, 95, 255}},
		60:  Color{TRUE_COLOR, color.NRGBA{95, 95, 135, 255}},
		61:  Color{TRUE_COLOR, color.NRGBA{95, 95, 175, 255}},
		62:  Color{TRUE_COLOR, color.NRGBA{95, 95, 215, 255}},
		63:  Color{TRUE_COLOR, color.NRGBA{95, 95, 255, 255}},
		64:  Color{TRUE_COLOR, color.NRGBA{95, 135, 0, 255}},
		65:  Color{TRUE_COLOR, color.NRGBA{95, 135, 95, 255}},
		66:  Color{TRUE_COLOR, color.NRGBA{95, 135, 135, 255}},
		67:  Color{TRUE_COLOR, color.NRGBA{95, 135, 175, 255}},
		68:  Color{TRUE_COLOR, color.NRGBA{95, 135, 215, 255}},
		69:  Color{TRUE_COLOR, color.NRGBA{95, 135, 255, 255}},
		70:  Color{TRUE_COLOR, color.NRGBA{95, 175, 0, 255}},
		71:  Color{TRUE_COLOR, color.NRGBA{95, 175, 95, 255}},
		72:  Color{TRUE_COLOR, color.NRGBA{95, 175, 135, 255}},
		73:  Color{TRUE_COLOR, color.NRGBA{95, 175, 175, 255}},
		74:  Color{TRUE_COLOR, color.NRGBA{95, 175, 215, 255}},
		75:  Color{TRUE_COLOR, color.NRGBA{95, 175, 255, 255}},
		76:  Color{TRUE_COLOR, color.NRGBA{95, 215, 0, 255}},
		77:  Color{TRUE_COLOR, color.NRGBA{95, 215, 95, 255}},
		78:  Color{TRUE_COLOR, color.NRGBA{95, 215, 135, 255}},
		79:  Color{TRUE_COLOR, color.NRGBA{95, 215, 175, 255}},
		80:  Color{TRUE_COLOR, color.NRGBA{95, 215, 215, 255}},
		81:  Color{TRUE_COLOR, color.NRGBA{95, 215, 255, 255}},
		82:  Color{TRUE_COLOR, color.NRGBA{95, 255, 0, 255}},
		83:  Color{TRUE_COLOR, color.NRGBA{95, 255, 95, 255}},
		84:  Color{TRUE_COLOR, color.NRGBA{95, 255, 135, 255}},
		85:  Color{TRUE_COLOR, color.NRGBA{95, 255, 175, 255}},
		86:  Color{TRUE_COLOR, color.NRGBA{95, 255, 215, 255}},
		87:  Color{TRUE_COLOR, color.NRGBA{95, 255, 255, 255}},
		88:  Color{TRUE_COLOR, color.NRGBA{135, 0, 0, 255}},
		89:  Color{TRUE_COLOR, color.NRGBA{135, 0, 95, 255}},
		90:  Color{TRUE_COLOR, color.NRGBA{135, 0, 135, 255}},
		91:  Color{TRUE_COLOR, color.NRGBA{135, 0, 175, 255}},
		92:  Color{TRUE_COLOR, color.NRGBA{135, 0, 215, 255}},
		93:  Color{TRUE_COLOR, color.NRGBA{135, 0, 255, 255}},
		94:  Color{TRUE_COLOR, color.NRGBA{135, 95, 0, 255}},
		95:  Color{TRUE_COLOR, color.NRGBA{135, 95, 95, 255}},
		96:  Color{TRUE_COLOR, color.NRGBA{135, 95, 135, 255}},
		97:  Color{TRUE_COLOR, color.NRGBA{135, 95, 175, 255}},
		98:  Color{TRUE_COLOR, color.NRGBA{135, 95, 215, 255}},
		99:  Color{TRUE_COLOR, color.NRGBA{135, 95, 255, 255}},
		100: Color{TRUE_COLOR, color.NRGBA{135, 135, 0, 255}},
		101: Color{TRUE_COLOR, color.NRGBA{135, 135, 95, 255}},
		102: Color{TRUE_COLOR, color.NRGBA{135, 135, 135, 255}},
		103: Color{TRUE_COLOR, color.NRGBA{135, 135, 175, 255}},
		104: Color{TRUE_COLOR, color.NRGBA{135, 135, 215, 255}},
		105: Color{TRUE_COLOR, color.NRGBA{135, 135, 255, 255}},
		106: Color{TRUE_COLOR, color.NRGBA{135, 175, 0, 255}},
		107: Color{TRUE_COLOR, color.NRGBA{135, 175, 95, 255}},
		108: Color{TRUE_COLOR, color.NRGBA{135, 175, 135, 255}},
		109: Color{TRUE_COLOR, color.NRGBA{135, 175, 175, 255}},
		110: Color{TRUE_COLOR, color.NRGBA{135, 175, 215, 255}},
		111: Color{TRUE_COLOR, color.NRGBA{135, 175, 255, 255}},
		112: Color{TRUE_COLOR, color.NRGBA{135, 215, 0, 255}},
		113: Color{TRUE_COLOR, color.NRGBA{135, 215, 95, 255}},
		114: Color{TRUE_COLOR, color.NRGBA{135, 215, 135, 255}},
		115: Color{TRUE_COLOR, color.NRGBA{135, 215, 175, 255}},
		116: Color{TRUE_COLOR, color.NRGBA{135, 215, 215, 255}},
		117: Color{TRUE_COLOR, color.NRGBA{135, 215, 255, 255}},
		118: Color{TRUE_COLOR, color.NRGBA{135, 255, 0, 255}},
		119: Color{TRUE_COLOR, color.NRGBA{135, 255, 95, 255}},
		120: Color{TRUE_COLOR, color.NRGBA{135, 255, 135, 255}},
		121: Color{TRUE_COLOR, color.NRGBA{135, 255, 175, 255}},
		122: Color{TRUE_COLOR, color.NRGBA{135, 255, 215, 255}},
		123: Color{TRUE_COLOR, color.NRGBA{135, 255, 255, 255}},
		124: Color{TRUE_COLOR, color.NRGBA{175, 0, 0, 255}},
		125: Color{TRUE_COLOR, color.NRGBA{175, 0, 95, 255}},
		126: Color{TRUE_COLOR, color.NRGBA{175, 0, 135, 255}},
		127: Color{TRUE_COLOR, color.NRGBA{175, 0, 175, 255}},
		128: Color{TRUE_COLOR, color.NRGBA{175, 0, 215, 255}},
		129: Color{TRUE_COLOR, color.NRGBA{175, 0, 255, 255}},
		130: Color{TRUE_COLOR, color.NRGBA{175, 95, 0, 255}},
		131: Color{TRUE_COLOR, color.NRGBA{175, 95, 95, 255}},
		132: Color{TRUE_COLOR, color.NRGBA{175, 95, 135, 255}},
		133: Color{TRUE_COLOR, color.NRGBA{175, 95, 175, 255}},
		134: Color{TRUE_COLOR, color.NRGBA{175, 95, 215, 255}},
		135: Color{TRUE_COLOR, color.NRGBA{175, 95, 255, 255}},
		136: Color{TRUE_COLOR, color.NRGBA{175, 135, 0, 255}},
		137: Color{TRUE_COLOR, color.NRGBA{175, 135, 95, 255}},
		138: Color{TRUE_COLOR, color.NRGBA{175, 135, 135, 255}},
		139: Color{TRUE_COLOR, color.NRGBA{175, 135, 175, 255}},
		140: Color{TRUE_COLOR, color.NRGBA{175, 135, 215, 255}},
		141: Color{TRUE_COLOR, color.NRGBA{175, 135, 255, 255}},
		142: Color{TRUE_COLOR, color.NRGBA{175, 175, 0, 255}},
		143: Color{TRUE_COLOR, color.NRGBA{175, 175, 95, 255}},
		144: Color{TRUE_COLOR, color.NRGBA{175, 175, 135, 255}},
		145: Color{TRUE_COLOR, color.NRGBA{175, 175, 175, 255}},
		146: Color{TRUE_COLOR, color.NRGBA{175, 175, 215, 255}},
		147: Color{TRUE_COLOR, color.NRGBA{175, 175, 255, 255}},
		148: Color{TRUE_COLOR, color.NRGBA{175, 215, 0, 255}},
		149: Color{TRUE_COLOR, color.NRGBA{175, 215, 95, 255}},
		150: Color{TRUE_COLOR, color.NRGBA{175, 215, 135, 255}},
		151: Color{TRUE_COLOR, color.NRGBA{175, 215, 175, 255}},
		152: Color{TRUE_COLOR, color.NRGBA{175, 215, 215, 255}},
		153: Color{TRUE_COLOR, color.NRGBA{175, 215, 255, 255}},
		154: Color{TRUE_COLOR, color.NRGBA{175, 255, 0, 255}},
		155: Color{TRUE_COLOR, color.NRGBA{175, 255, 95, 255}},
		156: Color{TRUE_COLOR, color.NRGBA{175, 255, 135, 255}},
		157: Color{TRUE_COLOR, color.NRGBA{175, 255, 175, 255}},
		158: Color{TRUE_COLOR, color.NRGBA{175, 255, 215, 255}},
		159: Color{TRUE_COLOR, color.NRGBA{175, 255, 255, 255}},
		160: Color{TRUE_COLOR, color.NRGBA{215, 0, 0, 255}},
		161: Color{TRUE_COLOR, color.NRGBA{215, 0, 95, 255}},
		162: Color{TRUE_COLOR, color.NRGBA{215, 0, 135, 255}},
		163: Color{TRUE_COLOR, color.NRGBA{215, 0, 175, 255}},
		164: Color{TRUE_COLOR, color.NRGBA{215, 0, 215, 255}},
		165: Color{TRUE_COLOR, color.NRGBA{215, 0, 255, 255}},
		166: Color{TRUE_COLOR, color.NRGBA{215, 95, 0, 255}},
		167: Color{TRUE_COLOR, color.NRGBA{215, 95, 95, 255}},
		168: Color{TRUE_COLOR, color.NRGBA{215, 95, 135, 255}},
		169: Color{TRUE_COLOR, color.NRGBA{215, 95, 175, 255}},
		170: Color{TRUE_COLOR, color.NRGBA{215, 95, 215, 255}},
		171: Color{TRUE_COLOR, color.NRGBA{215, 95, 255, 255}},
		172: Color{TRUE_COLOR, color.NRGBA{215, 135, 0, 255}},
		173: Color{TRUE_COLOR, color.NRGBA{215, 135, 95, 255}},
		174: Color{TRUE_COLOR, color.NRGBA{215, 135, 135, 255}},
		175: Color{TRUE_COLOR, color.NRGBA{215, 135, 175, 255}},
		176: Color{TRUE_COLOR, color.NRGBA{215, 135, 215, 255}},
		177: Color{TRUE_COLOR, color.NRGBA{215, 135, 255, 255}},
		178: Color{TRUE_COLOR, color.NRGBA{215, 175, 0, 255}},
		179: Color{TRUE_COLOR, color.NRGBA{215, 175, 95, 255}},
		180: Color{TRUE_COLOR, color.NRGBA{215, 175, 135, 255}},
		181: Color{TRUE_COLOR, color.NRGBA{215, 175, 175, 255}},
		182: Color{TRUE_COLOR, color.NRGBA{215, 175, 215, 255}},
		183: Color{TRUE_COLOR, color.NRGBA{215, 175, 255, 255}},
		184: Color{TRUE_COLOR, color.NRGBA{215, 215, 0, 255}},
		185: Color{TRUE_COLOR, color.NRGBA{215, 215, 95, 255}},
		186: Color{TRUE_COLOR, color.NRGBA{215, 215, 135, 255}},
		187: Color{TRUE_COLOR, color.NRGBA{215, 215, 175, 255}},
		188: Color{TRUE_COLOR, color.NRGBA{215, 215, 215, 255}},
		189: Color{TRUE_COLOR, color.NRGBA{215, 215, 255, 255}},
		190: Color{TRUE_COLOR, color.NRGBA{215, 255, 0, 255}},
		191: Color{TRUE_COLOR, color.NRGBA{215, 255, 95, 255}},
		192: Color{TRUE_COLOR, color.NRGBA{215, 255, 135, 255}},
		193: Color{TRUE_COLOR, color.NRGBA{215, 255, 175, 255}},
		194: Color{TRUE_COLOR, color.NRGBA{215, 255, 215, 255}},
		195: Color{TRUE_COLOR, color.NRGBA{215, 255, 255, 255}},
		196: Color{TRUE_COLOR, color.NRGBA{255, 0, 0, 255}},
		197: Color{TRUE_COLOR, color.NRGBA{255, 0, 95, 255}},
		198: Color{TRUE_COLOR, color.NRGBA{255, 0, 135, 255}},
		199: Color{TRUE_COLOR, color.NRGBA{255, 0, 175, 255}},
		200: Color{TRUE_COLOR, color.NRGBA{255, 0, 215, 255}},
		201: Color{TRUE_COLOR, color.NRGBA{255, 0, 255, 255}},
		202: Color{TRUE_COLOR, color.NRGBA{255, 95, 0, 255}},
		203: Color{TRUE_COLOR, color.NRGBA{255, 95, 95, 255}},
		204: Color{TRUE_COLOR, color.NRGBA{255, 95, 135, 255}},
		205: Color{TRUE_COLOR, color.NRGBA{255, 95, 175, 255}},
		206: Color{TRUE_COLOR, color.NRGBA{255, 95, 215, 255}},
		207: Color{TRUE_COLOR, color.NRGBA{255, 95, 255, 255}},
		208: Color{TRUE_COLOR, color.NRGBA{255, 135, 0, 255}},
		209: Color{TRUE_COLOR, color.NRGBA{255, 135, 95, 255}},
		210: Color{TRUE_COLOR, color.NRGBA{255, 135, 135, 255}},
		211: Color{TRUE_COLOR, color.NRGBA{255, 135, 175, 255}},
		212: Color{TRUE_COLOR, color.NRGBA{255, 135, 215, 255}},
		213: Color{TRUE_COLOR, color.NRGBA{255, 135, 255, 255}},
		214: Color{TRUE_COLOR, color.NRGBA{255, 175, 0, 255}},
		215: Color{TRUE_COLOR, color.NRGBA{255, 175, 95, 255}},
		216: Color{TRUE_COLOR, color.NRGBA{255, 175, 135, 255}},
		217: Color{TRUE_COLOR, color.NRGBA{255, 175, 175, 255}},
		218: Color{TRUE_COLOR, color.NRGBA{255, 175, 215, 255}},
		219: Color{TRUE_COLOR, color.NRGBA{255, 175, 255, 255}},
		220: Color{TRUE_COLOR, color.NRGBA{255, 215, 0, 255}},
		221: Color{TRUE_COLOR, color.NRGBA{255, 215, 95, 255}},
		222: Color{TRUE_COLOR, color.NRGBA{255, 215, 135, 255}},
		223: Color{TRUE_COLOR, color.NRGBA{255, 215, 175, 255}},
		224: Color{TRUE_COLOR, color.NRGBA{255, 215, 215, 255}},
		225: Color{TRUE_COLOR, color.NRGBA{255, 215, 255, 255}},
		226: Color{TRUE_COLOR, color.NRGBA{255, 255, 0, 255}},
		227: Color{TRUE_COLOR, color.NRGBA{255, 255, 95, 255}},
		228: Color{TRUE_COLOR, color.NRGBA{255, 255, 135, 255}},
		229: Color{TRUE_COLOR, color.NRGBA{255, 255, 175, 255}},
		230: Color{TRUE_COLOR, color.NRGBA{255, 255, 215, 255}},
		231: Color{TRUE_COLOR, color.NRGBA{255, 255, 255, 255}},
		232: Color{TRUE_COLOR, color.NRGBA{8, 8, 8, 255}},
		233: Color{TRUE_COLOR, color.NRGBA{18, 18, 18, 255}},
		234: Color{TRUE_COLOR, color.NRGBA{28, 28, 28, 255}},
		235: Color{TRUE_COLOR, color.NRGBA{38, 38, 38, 255}},
		236: Color{TRUE_COLOR, color.NRGBA{48, 48, 48, 255}},
		237: Color{TRUE_COLOR, color.NRGBA{58, 58, 58, 255}},
		238: Color{TRUE_COLOR, color.NRGBA{68, 68, 68, 255}},
		239: Color{TRUE_COLOR, color.NRGBA{78, 78, 78, 255}},
		240: Color{TRUE_COLOR, color.NRGBA{88, 88, 88, 255}},
		241: Color{TRUE_COLOR, color.NRGBA{98, 98, 98, 255}},
		242: Color{TRUE_COLOR, color.NRGBA{108, 108, 108, 255}},
		243: Color{TRUE_COLOR, color.NRGBA{118, 118, 118, 255}},
		244: Color{TRUE_COLOR, color.NRGBA{128, 128, 128, 255}},
		245: Color{TRUE_COLOR, color.NRGBA{138, 138, 138, 255}},
		246: Color{TRUE_COLOR, color.NRGBA{148, 148, 148, 255}},
		247: Color{TRUE_COLOR, color.NRGBA{158, 158, 158, 255}},
		248: Color{TRUE_COLOR, color.NRGBA{168, 168, 168, 255}},
		249: Color{TRUE_COLOR, color.NRGBA{178, 178, 178, 255}},
		250: Color{TRUE_COLOR, color.NRGBA{188, 188, 188, 255}},
		251: Color{TRUE_COLOR, color.NRGBA{198, 198, 198, 255}},
		252: Color{TRUE_COLOR, color.NRGBA{208, 208, 208, 255}},
		253: Color{TRUE_COLOR, color.NRGBA{218, 218, 218, 255}},
		254: Color{TRUE_COLOR, color.NRGBA{228, 228, 228, 255}},
		255: Color{TRUE_COLOR, color.NRGBA{238, 238, 238, 255}},
	}
}

func (o Color) String() string {
	if o.IsTrueColor() {
		if o.A != 255 {
			return fmt.Sprintf("#%02X%02X%02X%02X", o.R, o.G, o.B, o.A)
		}
		return fmt.Sprintf("#%02X%02X%02X", o.R, o.G, o.B)
	}
	if o.IsColor256() {
		return fmt.Sprintf("$%d", o.Color256)
	}
	if o.IsTransparent() {
		return "TRANSPARENT"
	}
	return "not a valid color"
}

var findColorName func(Color) (string, bool) = nil

func SetFuncToFindTheNameOfAColor(f func(Color) (string, bool)) {
	findColorName = f
}

func (o Color) StringSmart() string {
	if findColorName != nil {
		if s, ok := findColorName(o); ok {
			return s
		}
	}
	return o.String()
}

func (o Color) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["color"] = o.String()
	return json.Marshal(m)
}
