package screen

type RuneGetter interface {
	GetRune(int, int) (rune, bool)
}

type RuneSetter interface {
	SetRune(int, int, rune)
}

type RuneModifier interface {
	RuneGetter
	RuneSetter
}

type ColorGetter interface {
	GetColor(int, int) (Color, bool)
}

type ColorSetter interface {
	SetColor(int, int, Color)
}

type ColorModifier interface {
	ColorGetter
	ColorSetter
}

type EmphasisGetter interface {
	GetEmphasis(int, int) (Attributes, bool)
}

type EmphasisSetter interface {
	SetEmphasis(int, int, Attributes)
}

type EmphasisModifier interface {
	EmphasisGetter
	EmphasisSetter
}

type SizeGetter interface {
	Height() int
	Width() int
}

type SizeSetter interface {
	SetHeight(int)
	SetWidth(int)
	Resize(int, int)
	SetOnSizeChange(func(int, int))
	GetOnSizeChange() func(int, int)
	DoOnSizeChange(int, int)
}

type Sizer interface {
	SizeGetter
	SizeSetter
}

type PositionGetter interface {
	Row() int
	Col() int
	GetPosition() Coordinates
}

type PositionSetter interface {
	SetRow(int)
	SetCol(int)
	SetPosition(int, int)
	MoveRow(int)
	MoveCol(int)
	Move(int, int)
	CanMove() bool
	SetCanMove(bool)
}

type Positioner interface {
	PositionGetter
	PositionSetter
}

type SizeAndPositionGetter interface {
	SizeGetter
	PositionGetter
}

type SizeAndPositionSetter interface {
	SizeSetter
	PositionSetter
}

type SizeGetterAndPositionSetter interface {
	SizeGetter
	PositionSetter
}
