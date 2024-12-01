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
