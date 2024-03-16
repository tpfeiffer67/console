package engine

import (
	"errors"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/tpfeiffer67/console/screen"
	"github.com/tpfeiffer67/console/ui/message"
	"github.com/tpfeiffer67/console/ui/ntt"
	"github.com/tpfeiffer67/console/ui/property"
	"github.com/tpfeiffer67/console/ui/theme"
)

type Logger interface {
	Debug(msg string, args ...any)
}

type NilLogger struct{}

func (o NilLogger) Debug(msg string, args ...any) {}

type Engine struct {
	*Screen
	Entities
	Actions
	theme.ITheme
	FocusedEntity any
	EditMode      bool
	screenHeight  int
	screenWidth   int
	renderFunc    func(o *screen.Buffer) int
	mouseControl
	inputMessageChannel     chan message.InputMessage
	uiMessageChannel        chan message.Message
	refreshCount            int
	refreshRequests         int
	refreshRequestsFollower int
	renderingMutex          sync.Mutex // renderingMutex is necessary to prevent resizing while the output process is running
	RefreshInterval         time.Duration
	isRunning               bool
	OnModeChange
	OnRefresh
	OnKey
	OnMouse
	property.OnScreenResize
	MinimizeVobj any
	Logger
	beforeInputMessageProcessing func(message.InputMessage)
}

func New(inputMessageChannel chan message.InputMessage, uiMessageChannel chan message.Message, renderFunc func(o *screen.Buffer) int, refreshIntervalInMilliseconds int) (*Engine, error) {
	if refreshIntervalInMilliseconds < 1 || refreshIntervalInMilliseconds > 1000 {
		return nil, errors.New("the value of refreshIntervalInMilliseconds must be between 1 and 1000")
	}
	o := new(Engine)
	o.ITheme = theme.NewTheme()
	o.Logger = &NilLogger{}
	o.Actions = NewActions()
	o.Entities = make(map[string]any)
	o.pointer = ntt.NewPointer()
	o.initScreen()

	o.RefreshInterval = time.Millisecond * time.Duration(refreshIntervalInMilliseconds)
	o.renderFunc = renderFunc
	o.inputMessageChannel = inputMessageChannel
	o.uiMessageChannel = uiMessageChannel

	return o, nil
}

func (o *Engine) SetLogger(logger Logger) {
	o.Logger = logger
}

func (o *Engine) SetFuncBeforeInputMessageProcessing(f func(message.InputMessage)) {
	o.beforeInputMessageProcessing = f
}

func (o *Engine) AddEntity(a any, parent string) any {
	id := a.(property.IId).Id() // An entity must at least have an Id, otherwize this make the program crash.  #directId
	o.Entities[id] = a

	if e, ok := a.(property.IParent); ok {
		e.SetParent(parent)
	}

	if e, ok := a.(message.IMessageSender); ok {
		e.SetSendMessage(func(messageId message.MessageId, messageParams interface{}, objId string) {
			o.SendMessage(messageId, messageParams, objId)
		})
	}

	order := o.nextOrder(parent)

	if e, ok := a.(property.IOrder); ok {
		e.SetOrder(order)
	}

	if e, ok := a.(property.IFocus); ok {
		e.SetFocusOrder(order)
	}

	if e, ok := a.(property.IZOrder); ok {
		if e.ZOrder() == 0 {
			e.SetZOrder(order)
		}
	}
	return a
}

func (o *Engine) GetRefreshCounters() (int, int) {
	return o.refreshRequestsFollower, o.refreshCount
}

func (o *Engine) Start() {
	o.isRunning = true
	go o.processInputMessages()
	go o.processUiMessages()
	go o.refresh()
	o.doOnModeChange(o.EditMode)
}

func (o *Engine) Run() {
	o.Start()
	for {
		time.Sleep(1000 * time.Millisecond)
		if !o.IsRunning() {
			break
		}
	}
}

func (o *Engine) IsRunning() bool { return o.isRunning }

func (o *Engine) Stop() {
	o.isRunning = false
	_ = keyboard.Close()
}

func (o *Engine) Refresh() {
	o.refreshRequests = o.refreshRequests + 1
}

func (o *Engine) refresh() {
	for {
		if o.refreshRequests != o.refreshRequestsFollower {
			// TODO Add a method OnBeforeRefresh
			// TODO Add a method or property to Enable/Disable Refresh
			o.renderingMutex.Lock()
			o.refreshRequestsFollower = o.refreshRequests
			o.refreshCount = o.refreshCount + 1
			o.doOnRefresh(o.refreshRequests, o.refreshCount)
			o.clearScreenBuffer()
			o.renderEntitiesToScreenBuffer()
			o.renderScreenBufferToPhysicalOutput()
			o.renderingMutex.Unlock()
		}
		time.Sleep(o.RefreshInterval)
	}
}

func (o *Engine) processInputMessages() {
	for {
		m := <-o.inputMessageChannel
		if o.beforeInputMessageProcessing != nil {
			o.beforeInputMessageProcessing(m)
		}
		o.Debug("INPMES", "Msg", m)

		switch mes := m.(type) {
		case message.InputMessageScreenSize:
			o.updateScreenSize(mes.Height, mes.Width)
			o.Refresh()
		case message.InputMessageMouse:
			o.processMouse(mes.ParamsMouse)
			o.Refresh()
		case message.InputMessageKey:
			o.processKey(mes.ParamsKey)
			o.Refresh()
		}
	}
}

func (o *Engine) processUiMessages() {
	for loop := true; loop; {
		m := <-o.uiMessageChannel
		o.Debug("MESSAGE", "Msg", m)

		switch m.MessageId {

		case message.MessageIdExit:
			loop = false

		case message.MessageIdSetTheme:
			if params, ok := m.Params.(message.ParamsSetTheme); ok {
				err := theme.SetTheme(params.Name)
				if err == nil {
					o.UpdateEntitiesTheme()
					o.Refresh()
				}
			}

		case message.MessageIdMove:
			if params, ok := m.Params.(message.ParamsMove); ok {
				if a, ok := o.GetEntityById(m.ObjId); ok {
					if e, ok := a.(screen.Positioner); ok {
						if e.CanMove() {
							currentPosition := e.GetPosition()
							e.SetPosition(currentPosition.Row+params.Rows, currentPosition.Col+params.Cols)
							o.Refresh()
						}
					}
				}
			}

		case message.MessageIdResize:
			if params, ok := m.Params.(message.ParamsResize); ok {
				if a, ok := o.GetEntityById(m.ObjId); ok {
					if e, ok := a.(screen.SizeSetter); ok {
						e.Resize(params.Rows, params.Cols)
						o.Refresh()
					}
				}
			}

		case message.MessageIdFocus:
			o.FocusById(m.ObjId)
			o.Refresh()

		case message.MessageIdFocusNext:
			o.FocusNext()

		case message.MessageIdFocusPrevious:
			o.FocusPrevious()

		case message.MessageIdFocusParent:
			o.FocusParent()

		case message.MessageIdFocusChild:
			o.FocusChild()

		case message.MessageIdZOrderUp:
			o.ZOrderUp(m.ObjId)
			o.Refresh()

		case message.MessageIdZOrderDown:
			o.ZOrderDown(m.ObjId)
			o.Refresh()

		case message.MessageIdZOrderTop:
			o.ZOrderToTop(m.ObjId)
			o.Refresh()

		case message.MessageIdZOrderBottom:
			o.ZOrderToBottom(m.ObjId)
			o.Refresh()

		/*
			case message.MessageIdMinimize:
				if e, ok := o.GetEntityById(m.ObjId); ok {
					if p, ok := e.(ntt.Paneler); ok {
						p.Minimize(o.MinimizeVobj)
					}
				}

			case message.MessageIdRestoreSize:
				if e, ok := o.GetEntityById(m.ObjId); ok {
					if p, ok := e.(ntt.Paneler); ok {
						p.RestoreSize()
					}
				}
		*/ // #minimizer
		default:
			if a, ok := o.GetEntityById(m.ObjId); ok {
				if e, ok := a.(message.IMessageListener); ok {
					e.ProcessMessage(m.MessageId, m.Params)
				}
			}
		}
	}
	o.Stop()
}

func (o *Engine) GetTheme() theme.ITheme {
	return o.ITheme
}

/*
   A REMETTRE
   		case message.MessageIdMove:
   			focusedObject := o.focusedObject() // can be nil
   			if focusedObject != nil {
   				if params, ok := m.Params.(message.ParamsMove); ok {
   					log.P rintln("[MESSAGE]", m.MessageId, params, "->", focusedObject)
   					focusedObject.(Mover).Move(params.Rows, params.Cols)
   				}
   				o.Refresh()
   			}

   		case message.MessageIdSetPosition:
   			focusedObject := o.focusedObject() // can be nil
   			if focusedObject != nil {
   				if params, ok := m.Params.(message.ParamsSetPosition); ok {
   					log.P rintln("[MESSAGE]", m.MessageId, params)
   					focusedObject.(screen.Positioner).SetPosition(params.Row, params.Col)
   					o.Refresh()
   				}
   			}

   		case message.MessageIdResize:
   			focusedObject := o.focusedObject() // can be nil
   			if focusedObject != nil {
   				if params, ok := m.Params.(message.ParamsResize); ok {
   					newHeight := focusedObject.(screen.Sizer).Height() + params.Rows
   					newWidth := focusedObject.(screen.Sizer).Width() + params.Cols
   					log.P rintln("[MESSAGE]", m.MessageId, params, "->", focusedObject, newHeight, newWidth)
   					focusedObject.(screen.Sizer).Resize(newHeight, newWidth)
   					focusedObject.(property.IOnDraw).DoOnDraw()
   				}
   				o.Refresh()
   			}
*/
