package helperfunctions

import "os"

func PrintGrid(size int, tetroSlc [][]string) {
	// Create empty grid of specified size
	grid := make([][]byte, size)

	for i := range grid {
		grid[i] = make([]byte, size)
	}
	placeTetromino(grid, tetroSlc, 0, size) // Place tetrominos

	// Print grid
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				os.Stdout.WriteString(".") // Replace blanks with '.'s
			} else {
				os.Stdout.WriteString(string(grid[i][j])) // Print alphabetic character
			}
		}
		os.Stdout.WriteString("") // Print newline to move on to next line
	}
}
