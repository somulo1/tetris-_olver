package helperfunctions

import (
	"math"
	"strings"
)

type Tetrominos struct {
	Tet [][]string
}

// Estimates the initial grid size for the tetrominoes
func calculateInitialGridSize(tet *Tetrominos) (int, int) {
	var maxWidth, maxHeight int
	for _, tetromino := range tet.Tet {
		width := len(tetromino[0])
		height := len(tetromino)
		if height > maxHeight {
			maxHeight = height
		}
		if width > maxWidth {
			maxWidth = width
		}
	}
	gridSize := int(math.Sqrt(float64(maxWidth * maxHeight)))
	return gridSize, gridSize
}

// Creates a grid initialized with '.' characters
func createGrid(width, height int) [][]string {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return grid
}

// Trims unnecessary empty lines from tetrominoes
func TrimTetrominos(tet *Tetrominos) (*Tetrominos, error) {
	var trimmedTetrominos [][]string
	for _, tetromino := range tet.Tet {
		if !isValidTetromino(tetromino) {
			return nil, ErrInvalidTetroType
		}
		trimmedTetrominos = append(trimmedTetrominos, trimEmptyLines(tetromino))
	}
	return &Tetrominos{Tet: trimmedTetrominos}, nil
}

// Removes empty rows and columns from a tetromino
func trimEmptyLines(tetromino []string) []string {
	tetromino = removeEmptyRows(tetromino)
	return removeEmptyColumns(tetromino)
}

// Removes rows that are entirely '.'
func removeEmptyRows(tetromino []string) []string {
	var nonEmptyRows []string
	for _, row := range tetromino {
		if strings.Trim(row, ".") != "" {
			nonEmptyRows = append(nonEmptyRows, row)
		}
	}
	return nonEmptyRows
}

// Removes columns that are entirely '.'
func removeEmptyColumns(tetromino []string) []string {
	width := len(tetromino[0])
	for x := 0; x < width; x++ {
		isEmptyColumn := true
		for _, row := range tetromino {
			if row[x] != '.' {
				isEmptyColumn = false
				break
			}
		}
		if isEmptyColumn {
			for i := range tetromino {
				tetromino[i] = tetromino[i][:x] + tetromino[i][x+1:]
			}
			x--     // Recheck the current column index after removal
			width-- // Adjust the width
		}
	}
	return tetromino
}

// Validates if a tetromino is a 4x4 grid with exactly 4 blocks connected
func isValidTetromino(tetromino []string) bool {
	if len(tetromino) != 4 || len(tetromino[0]) != 4 {
		return false
	}

	blocks := 0
	connections := 0

	for y := range tetromino {
		for x := range tetromino[y] {
			if tetromino[y][x] != '.' {
				blocks++
				connections += countConnections(tetromino, x, y)
			}
		}
	}

	return blocks == 4 && connections >= 6 && connections <= 8
}

// Counts the number of connections of a block at (x, y) with its neighbors
func countConnections(tetromino []string, x, y int) int {
	connections := 0
	if y > 0 && tetromino[y-1][x] != '.' {
		connections++
	}
	if y < len(tetromino)-1 && tetromino[y+1][x] != '.' {
		connections++
	}
	if x > 0 && tetromino[y][x-1] != '.' {
		connections++
	}
	if x < len(tetromino[y])-1 && tetromino[y][x+1] != '.' {
		connections++
	}
	return connections
}
