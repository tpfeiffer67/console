package property

type IOrder interface {
	Order() int
	SetOrder(int)
}

type PropertyOrder struct {
	order int
}

func (o *PropertyOrder) Order() int     { return o.order }
func (o *PropertyOrder) SetOrder(i int) { o.order = i }
