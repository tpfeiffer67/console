package screenutils

import (
	"image"

	"github.com/nfnt/resize"
	"github.com/tpfeiffer67/console/screen"
)

// TODO Add error management, w,h=0 (input or calculated), Image=nil, ...
type ConsoleImage struct {
	originalImage  image.Image
	resizedImage   image.Image
	originalHeight int
	originalWidth  int
}

func NewConsoleImage(img image.Image, height, width int) *ConsoleImage {
	o := new(ConsoleImage)
	o.originalImage = img
	o.originalHeight = o.originalImage.Bounds().Max.Y
	o.originalWidth = o.originalImage.Bounds().Max.X
	o.Resize(height, width)
	return o
}

func (o *ConsoleImage) Resize(wantedHeight, wantedWidth int) (int, int) {
	if wantedHeight != o.originalHeight || wantedWidth != o.originalWidth { // check if resizing the image is required
		height, width := processNewSize(wantedHeight, wantedWidth, o.originalHeight, o.originalWidth)
		o.resizedImage = resize.Resize(uint(width), uint(height), o.originalImage, resize.NearestNeighbor)
		return height, width
	} else {
		o.resizedImage = o.originalImage
		return o.originalHeight, o.originalWidth
	}
}

func processNewSize(wantedHeightInChar, wantedWidth, imageHeight, imageWidth int) (int, int) {
	wantedHeight := wantedHeightInChar * 2 // 1 character = 2 pixels height

	ratio := float32(wantedHeight) / float32(imageHeight)
	width := int(float32(imageWidth) * ratio)
	if width <= wantedWidth {
		return wantedHeight, width
	}

	ratio = float32(wantedWidth) / float32(imageWidth)
	height := int(float32(imageHeight) * ratio)
	return height, wantedWidth
}

func (o *ConsoleImage) Draw(runes screen.RuneSetter, foreground screen.ColorSetter, background screen.ColorSetter, row, col int) {
	if o.resizedImage == nil {
		return
	}
	img := o.resizedImage
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	for j := 0; j < height; j++ {
		if j%2 != 0 {
			for i := 0; i < width; i++ {
				c := col + i
				runes.SetRune(row, c, '▄')
				color := colorFromPixelAt(img, i, j-1)
				background.SetColor(row, c, color)
				color = colorFromPixelAt(img, i, j)
				foreground.SetColor(row, c, color)
			}
			row = row + 1
		}
	}
}

func (o *ConsoleImage) DrawWithAlpha(runes screen.RuneSetter, foreground screen.ColorSetter, background screen.ColorSetter, row, col int, alpha uint8) {
	if o.resizedImage == nil {
		return
	}
	img := o.resizedImage
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	for j := 0; j < height; j++ {
		if j%2 != 0 {
			for i := 0; i < width; i++ {
				c := col + i
				runes.SetRune(row, c, '▄')
				color := colorFromPixelAt(img, i, j-1)
				color.A = alpha
				background.SetColor(row, c, color)
				color = colorFromPixelAt(img, i, j)
				color.A = alpha
				foreground.SetColor(row, c, color)
			}
			row = row + 1
		}
	}
}

func (o *ConsoleImage) DrawOnCells(colorground screen.ColorSetter, row, col int) {
	if o.resizedImage == nil {
		return
	}
	img := o.resizedImage
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			color := colorFromPixelAt(img, i, j)
			colorground.SetColor(row+j, col+i, color)
		}
	}
}

func (o *ConsoleImage) DrawOnCellsWithAlpha(colorground screen.ColorSetter, row, col int, alpha uint8) {
	if o.resizedImage == nil {
		return
	}
	img := o.resizedImage
	height := img.Bounds().Max.Y
	width := img.Bounds().Max.X
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			color := colorFromPixelAt(img, i, j)
			color.A = alpha
			colorground.SetColor(row+j, col+i, color)
		}
	}
}

func colorFromPixelAt(img image.Image, x, y int) screen.Color {
	r, g, b, a := img.At(x, y).RGBA()
	return screen.NewTrueColorRGBA(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}
