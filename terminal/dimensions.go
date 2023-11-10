package terminal

import (
	"time"

	"github.com/nathan-fiscaletti/consolesize-go"
)

type Dimensions struct {
	height int
	width  int
}

func (o *Dimensions) Width() int  { return o.width }
func (o *Dimensions) Height() int { return o.height }

func GetDimensionsUpdater(intervalMilliseconds int) func(chan<- Dimensions) {
	var terminalDimensions Dimensions
	return func(terminalSizeChanges chan<- Dimensions) {
		for {
			width, height := consolesize.GetConsoleSize()
			if width != terminalDimensions.width || height != terminalDimensions.height {
				terminalDimensions.height = height
				terminalDimensions.width = width
				terminalSizeChanges <- terminalDimensions
			}
			time.Sleep(time.Millisecond * time.Duration(intervalMilliseconds))
		}
	}
}

func GetDimensions() *Dimensions {
	o := new(Dimensions)
	o.width, o.height = consolesize.GetConsoleSize()
	return o
}
