package ntt

import (
	"image"

	"github.com/tpfeiffer67/console/screenutils"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/theme"
)

// TODO Implement different display modes (Resize, scroll)
// TODO Add selected and hovered style

type ColorImage struct {
	IEntity
	theme.ITheme
	Background rune
}

func NewImage(id string, row, col int, height, width int, image image.Image, syst ISystem) *ColorImage {
	o := new(ColorImage)
	o.IEntity = NewEntity(id, height, width, syst)
	o.ITheme = theme.NewTheme(theme.STYLE_IMAGE_BACKGROUND, theme.IMAGE_BACKGROUND)
	o.SetPosition(row, col)
	o.SetCanMove(true)
	o.SetFocusable(true)
	o.updateAfterThemeChanged()

	consoleImage := screenutils.NewConsoleImage(image, height, width)
	consoleImage.Resize(height, width)
	consoleImage.Draw(o.GetRuneCanvas(), o.GetFColorCanvas(), o.GetBColorCanvas(), 0, 0)

	o.SetOnThemeChange(func() {
		o.updateAfterThemeChanged()
	})

	o.SetOnDraw(func() {
		style, _ := o.GetStyle(theme.STYLE_IMAGE_BACKGROUND)
		o.SetDefaults(o.Background, style)
		o.Clear()
		consoleImage.Draw(o.GetRuneCanvas(), o.GetFColorCanvas(), o.GetBColorCanvas(), 0, 0)
	})

	o.SetListener(message.MessageIdMouseClick, func(messageParams interface{}) bool {
		o.SendMessage(message.MessageIdZOrderTop, nil, id)
		o.SendMessage(message.MessageIdFocus, nil, id)
		return true
	})

	o.SetListener(message.MessageIdMouseMove, func(messageParams interface{}) bool {
		if params, ok := messageParams.(message.ParamsMouseExt); ok {
			o.Move(params.DeltaRow, params.DeltaCol)
		}
		return true
	})

	return o
}

func (o *ColorImage) updateAfterThemeChanged() {
	o.Background, _ = o.GetChar(theme.IMAGE_BACKGROUND)
}
