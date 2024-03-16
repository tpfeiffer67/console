package property

type IId interface {
	Id() string
	SetId(string)
}

type IParent interface {
	Parent() string
	SetParent(string)
}

type PropertyId struct {
	id     string
	parent string
}

func (o *PropertyId) Id() string     { return o.id }
func (o *PropertyId) SetId(s string) { o.id = s }

func (o *PropertyId) Parent() string     { return o.parent }
func (o *PropertyId) SetParent(s string) { o.parent = s }
