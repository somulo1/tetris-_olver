package helperfunction

import "strings"

func ValidateTetrominoFormat(tetromino string) bool {
	lines := strings.Split(strings.TrimSpace(tetromino), "\n")
	if len(lines) == 0 {
		return false
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			return false
		}
	}
	return true
}
