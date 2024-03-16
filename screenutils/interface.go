package screenutils

import "github.com/tpfeiffer67/console/screen"

type ICanvas interface {
	SetDefaults(rune, screen.Style)
	Clear()
	GetRuneCanvas() *screen.RuneCanvas
	GetFColorCanvas() *screen.FColorCanvas
	GetBColorCanvas() *screen.BColorCanvas
	GetEmphasisCanvas() *screen.EmphasisCanvas
	GetStencil() *screen.Stencil
	GetOrigins() screen.Coordinates
}

type IHasRuneCanvas interface {
	GetRuneCanvas() *screen.RuneCanvas
	Clear()
}
