package engine

type OnRefresh struct {
	onRefresh func(int, int)
}

func (o *OnRefresh) SetOnRefresh(f func(int, int)) {
	o.onRefresh = f
}

func (o *OnRefresh) GetOnRefresh() func(int, int) {
	return o.onRefresh
}

func (o *OnRefresh) doOnRefresh(refreshRequests, actualRefreshes int) {
	if o.onRefresh != nil {
		o.onRefresh(refreshRequests, actualRefreshes)
	}
}
