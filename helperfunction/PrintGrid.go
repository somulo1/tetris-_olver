package helperfunction

import (
	"fmt"
	"strings"
)

// 3.1 Format the grid for output
func FormatGrid(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, cell := range row {
			sb.WriteRune(cell)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// 3.2 Print the result
func PrintGrid(grid [][]rune) {
	fmt.Print(FormatGrid(grid))
}
