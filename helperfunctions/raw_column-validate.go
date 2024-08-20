// the tetrominos should have a total of four row and four columns
// the number of "#" in each cube must be equal to 4
package helperfunctions

import (
	"fmt"
	"strings"
)

// validate the tetromin arrangement
func ValidateTetrominoFormat(tetromino string) bool {
	lines := strings.Split(strings.TrimSpace(tetromino), "\n")
	fmt.Println(lines)
	if len(lines) != 4 {
		return false
	}
	var countDots int
	var countHash int
	// check for every line

	for i, line := range lines {
		//fmt.Println(line)
		if len(line) != 4 {
			return false
		}
		// for _ , char := range line {
		// 	fmt.Print(char)
		// }
		if string(line[i]) == "." {
			countDots++
		}
		if string(line[i]) == "#" {
			countHash++
		}

	}
	if countDots+countHash != 16 || countDots != 12 || countHash != 4 {
		fmt.Errorf("\nERROR1.1")
		return false

	}
	return true
}
