package engine

import "github.com/tpfeiffer67/console/ui/message"

type Action struct {
	hotkey string
	action func()
}

type Actions struct {
	actions map[string]Action
}

func NewActions() Actions {
	a := Actions{}
	a.actions = make(map[string]Action)
	return a
}

func (a *Actions) Add(id string, hotkey string, f func()) {
	na := Action{hotkey: hotkey, action: f}
	a.actions[id] = na
}

func (a *Actions) Do(id string) {
	if a, ok := a.actions[id]; ok {
		a.action()
	}
}

func (a *Actions) DoOnHotKey(keyParams message.ParamsKey) {
	for _, v := range a.actions {
		if v.hotkey == keyParams.Name {
			v.action()
			return
		}
	}
}
