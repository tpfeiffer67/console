package property

type IZOrder interface {
	ZOrder() int
	SetZOrder(int)
	ZOrderLayer() int
	SetZOrderLayer(int)
	ZOrderSum() int
}

type PropertyZOrder struct {
	zOrder      int
	zOrderLayer int
}

func (o *PropertyZOrder) ZOrder() int     { return o.zOrder }
func (o *PropertyZOrder) SetZOrder(i int) { o.zOrder = i }

func (o *PropertyZOrder) ZOrderLayer() int     { return o.zOrderLayer }
func (o *PropertyZOrder) SetZOrderLayer(i int) { o.zOrderLayer = i }

func (o *PropertyZOrder) ZOrderSum() int { return o.zOrder + o.zOrderLayer }
