package property

type Selecter interface {
	Select()
	Unselect()
	Selected() bool
	SetSelectGroup(int)
	GetSelectGroup() int
	SetOnSelect(func())
	GetOnSelect() func()
	DoOnSelect()
}

type OnSelect struct {
	selected   bool
	group      int
	onSelect   func()
	onUnselect func()
}

func (o *OnSelect) Select() {
	o.selected = true
	o.DoOnSelect()
}
func (o *OnSelect) Unselect() {
	o.selected = false
	o.DoOnUnselect()
}
func (o *OnSelect) Selected() bool { return o.selected }

func (o *OnSelect) SetSelectGroup(g int) { o.group = g }
func (o *OnSelect) GetSelectGroup() int  { return o.group }

func (o *OnSelect) SetOnSelect(f func()) { o.onSelect = f }

func (o *OnSelect) GetOnSelect() func() { return o.onSelect }

func (o *OnSelect) DoOnSelect() {
	if o.onSelect != nil {
		o.onSelect()
	}
}

func (o *OnSelect) SetOnUnselect(f func()) { o.onUnselect = f }

func (o *OnSelect) GetOnUnselect() func() { return o.onUnselect }

func (o *OnSelect) DoOnUnselect() {
	if o.onUnselect != nil {
		o.onUnselect()
	}
}
