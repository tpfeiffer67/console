package property

type IParentVisible interface {
	IParent
	IVisible
}

type IParentVisibleFocusable interface {
	IParent
	IVisible
	IFocus
}

type IParentVisibleZOrder interface {
	IParent
	IVisible
	IZOrder
}
