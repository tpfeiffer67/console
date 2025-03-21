package extwidget

import (
	"fmt"
	"strings"

	"github.com/tpfeiffer67/console/ui/engine"
	"github.com/tpfeiffer67/console/ui/entity/widget"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/property"
)

func AddDebugStatusBar(ngn *engine.Engine) {
	statusbardbg := widget.NewStatusBar("statusbardbg", ngn)
	label1 := widget.NewLabel("labelSB1", 0, 0, "", ngn)
	label2 := widget.NewLabel("labelSB2", 0, 0, "", ngn)
	label3 := widget.NewLabel("labelSB3", 0, 0, "", ngn)
	label4 := widget.NewLabel("labelSB4", 0, 0, "", ngn)
	label5 := widget.NewLabel("labelSB5", 0, 0, "", ngn)
	ngn.AddEntity(statusbardbg, "")
	ngn.AddEntity(label1, "statusbardbg")
	ngn.AddEntity(label2, "statusbardbg")
	ngn.AddEntity(label3, "statusbardbg")
	ngn.AddEntity(label4, "statusbardbg")
	ngn.AddEntity(label5, "statusbardbg")

	ngn.SetOnScreenResize(func(height, width int) {
		label2.SetText(fmt.Sprintf("🖥️  🠇 %d 🠆 %d", height, width))
	})

	ngn.SetOnKey(func(params *message.ParamsKey) bool {
		label4.SetText(fmt.Sprintf("[%s] Key %d %c %s%s%s ", params.Name, params.Key, params.Rune, renderCASKey(params.Ctrl, "c"), renderCASKey(params.Alt, "a"), renderCASKey(params.Shift, "s")))
		return false
	})

	label5.SetText("none")
	ngn.SetOnMouse(func(params *message.ParamsMouse, memoParams message.ParamsMouse, a any) bool {
		// Displays mouse coordinates
		label3.SetText(fmt.Sprintf("🖱️  🠇 %d 🠆 %d %s%s%s", params.Row, params.Col, renderButton(params.ButtonPrimary), renderButton(params.ButtonMiddle), renderButton(params.ButtonSecondary)))
		// Displays the current hovered entity Id
		if a != nil {
			label5.SetText(a.(property.IId).Id()) // #directId
		} else {
			label5.SetText("none")
		}
		return false
	})

	ngn.SetOnRefresh(func(refreshRequests, refreshCount int) {
		switch refreshCount % 4 {
		case 0:
			label1.SetText(" ◐")
		case 1:
			label1.SetText(" ◓")
		case 2:
			label1.SetText(" ◑")
		case 3:
			label1.SetText(" ◒")
		}
	})
}

func renderCASKey(b bool, s string) string {
	if b {
		return strings.ToUpper(s)
	}
	return s
}

func renderButton(b bool) string {
	if b {
		return "▮"
	}
	return "▯"
}
