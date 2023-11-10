package ntt

type SizeManager interface {
	SizeStatus() SizeStatus
	SetSizeStatus(SizeStatus)
	MemorizeNormalStatus(int, int, int, int, string)
	GetMemorizedNormalStatus() (int, int, int, int, string)
}

type SizeManagement struct {
	memoRow    int
	memoCol    int
	memoHeight int
	memoWidth  int
	memoParent string
	sizeStatus SizeStatus
}

func (o *SizeManagement) SizeStatus() SizeStatus     { return o.sizeStatus }
func (o *SizeManagement) SetSizeStatus(s SizeStatus) { o.sizeStatus = s }

func (o *SizeManagement) MemorizeNormalStatus(r, c, h, w int, p string) {
	o.memoRow = r
	o.memoCol = c
	o.memoHeight = h
	o.memoWidth = w
	o.memoParent = p
}

func (o *SizeManagement) GetMemorizedNormalStatus() (int, int, int, int, string) {
	return o.memoRow, o.memoCol, o.memoHeight, o.memoWidth, o.memoParent
}
