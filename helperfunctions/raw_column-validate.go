package helperfunctions

import (
	"strings"
)

func ValidateTetrominoFormat(filename string) bool {
	lines := strings.Split(strings.TrimSpace(filename), "\n")
	if len(lines) != 4 {
		return false
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != 4 && width != 4 {
			return false
		}
	}
	return true
	
}
