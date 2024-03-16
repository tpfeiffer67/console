package ntt

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/theme"
)

// Shadow
type Shadow struct {
	screen.Position
	*screenutils.ChromaticFilter
	Visible bool
}

func NewShadow(id string, row, col int, stencil *screen.Stencil) *Shadow {
	o := new(Shadow)
	shadowDensity := uint8(120)
	shadowColor := screen.NewTrueColorRGBA(0, 0, 0, shadowDensity)
	o.ChromaticFilter = screenutils.NewChromaticFilter(stencil, shadowColor)
	o.SetPosition(row, col)
	o.Visible = true
	return o
}

func (o *Shadow) Render(sb *screen.Buffer, pos screen.Coordinates) {
	if o.Visible {
		o.ChromaticFilter.Render(sb, o.Stencil, pos, screen.Coordinates{})
	}
}

// TODO Replace shadowValue by color. So it will be possible to have colored shadow.
func DrawShadowAccordingToTheTheme(screenbuffer *screen.Buffer, stencil *screen.Stencil, position screen.Coordinates, ntt theme.ITheme) {
	/* #minimizer
	if p, ok := ntt.(Paneler); ok {
		if p.SizeStatus() != SizeStatusNormal {
			return
		}
	}
	*/
	b, _ := ntt.GetBool(theme.SHADOW_ENABLED)
	if b {
		shadowVOffset, _ := ntt.GetInt(theme.SHADOW_VERTICAL_OFFSET)
		shadowHOffset, _ := ntt.GetInt(theme.SHADOW_HORIZONTAL_OFFSET)
		shadowValue, _ := ntt.GetUInt8(theme.SHADOW_VALUE)
		screenutils.DrawShadow(screenbuffer, stencil, position, shadowVOffset, shadowHOffset, shadowValue)
	}
}
