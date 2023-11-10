package engine

import (
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/ntt"
)

type Screen struct {
	*screen.Buffer
}

func (o *Engine) initScreen() {
	o.Screen = newScreen(SCREEN_ID, o.screenHeight, o.screenWidth)
}

func newScreen(id string, height, width int) *Screen {
	o := new(Screen)
	o.Buffer = screen.NewBuffer(height, width)
	return o
}

func (o *Engine) clearScreenBuffer() {
	o.Screen.Clear()
}

func (o *Engine) renderScreenBufferToPhysicalOutput() {
	o.renderFunc(o.Screen.Buffer)
}

func (o *Engine) updateScreenSize(height, width int) {
	if o.screenHeight == height && o.screenWidth == width {
		return
	}
	o.screenHeight = height
	o.screenWidth = width
	o.renderingMutex.Lock()
	o.Screen.Resize(height, width)
	o.Screen.Clear()
	o.DoOnScreenResize(height, width)
	o.processObjectsInTreeAfterScreenResized(height, width)
	o.Refresh()
	o.renderingMutex.Unlock()
}

func (o *Engine) processObjectsInTreeAfterScreenResized(height, width int) {
	f := func(e ntt.IEntity) {
		e.DoOnScreenResize(height, width)
	}
	o.callFuncWithAllEntities(f)
}
