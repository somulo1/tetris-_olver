//the tetrominos should have a total of four row and four columns
// the number of "#" in each cube must be equal to 4
package helperfunctions

import (
	"strings"
)
//validate the tetromin arrangement
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
