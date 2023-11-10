package screen

func verifyIfRowAndColAreInsideLimits(row, col int, height, width int) bool {
	if row < 0 {
		return false
	}
	if row >= height {
		return false
	}
	if col < 0 {
		return false
	}
	if col >= width {
		return false
	}
	return true
}
