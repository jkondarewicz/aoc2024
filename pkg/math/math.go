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
	X int
	Y int
}

func NewVertex(x int, y int) Vertex {
	return Vertex{X: x, Y: y}
}

func (vert Vertex) Add(vert2 Vertex) Vertex {
	return Vertex{X: vert.X + vert2.X, Y: vert.Y + vert2.Y}
}
func (vert Vertex) Opposite() Vertex {
	return Vertex{X: -vert.X, Y: -vert.Y}
}
func (vert Vertex) DiffBetweenVertexes(vert2 Vertex) Vertex {
	return Vertex{X: vert2.X - vert.X, Y: vert2.Y - vert.Y}
}

func (vert Vertex) Normalize() Vertex {
	return NewVertex(Normalize(vert.X), Normalize(vert.Y))
}
func (vert Vertex) Diff() int {
	return vert.Y - vert.X
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
