package aocmath

func Min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Normalize(a int) int {
	if a == 0 {
		return 0
	}
	return a / Abs(a)
}
type Vertex struct {
	A int
	B int
}
func (vert Vertex) Diff() int {
	return vert.B - vert.A
}
func (vert Vertex) DiffNormalized() int {
	diff := vert.Diff()
	if diff == 0 {
		return 0
	}
	diff = diff / Abs(diff)
	return diff
}
func (vert Vertex) DiffAbs() int {
	return Abs(vert.Diff())
}
