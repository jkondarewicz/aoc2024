package utils

func CharToInt(char rune) (int, bool) {
	if char >= '0' && char <= '9' {
		return int(char - '0'), true
	}
	return 0, false
}
