package helperfunction

import (
	"strings"
)

// PlaceTetrominoes places tetrominoes in the grid and returns the grid.
func PlaceTetrominoes(tetrominoes []string, size int) [][]rune {
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
	}

	// Initialize grid with '.' to indicate empty spaces
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	// Place each tetromino in the grid
	for i, tetromino := range tetrominoes {
		lines := strings.Split(strings.TrimSpace(tetromino), "\n")

		for y, line := range lines {
			for x, ch := range line {
				// Ensure tetromino fits within grid boundaries and is valid
				if y < size && x < size && ch == '#' {
					grid[y][x] = rune('A' + i) // Place tetromino with its corresponding letter
				}
			}
		}
	}

	return grid
}
