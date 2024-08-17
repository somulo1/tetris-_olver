package helperfunctions

func placeTetromino(grid [][]byte, tetroSlc [][]string, index, size int) bool {
	// Terminate solve function when all tetrominoes are placed on grid
	if index == len(tetroSlc) {
		return true
	}

	// Loop through every cell of the grid...
	// Checking if they are a viable starting point for placing a tetromino
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if canPlace(grid, tetroSlc[index], row, col, size) { // Check if specific tetromino can be placed
				place(grid, tetroSlc[index], row, col)
				if placeTetromino(grid, tetroSlc, index+1, size) {
					return true
				}
				remove(grid, tetroSlc[index], row, col) // Backtrack should program fail to place a tetromino
			}
		}
	}

	return false
}

// Check possibility of placing a specific tetromino on grid
func canPlace(grid [][]byte, tetromino []string, row, col, size int) bool {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] != '.' {
				// Triggered when block of tetromino falls out of bounds in grid
				if row+i >= size || col+j >= size || grid[row+i][col+j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

// Place tetromino in grid
func place(grid [][]byte, tetromino []string, row, col int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			if tetromino[i][j] != '.' {
				// Replace 0s on grid with Alpabetic letter of tetromino
				grid[row+i][col+j] = tetromino[i][j]
			}
		}
	}
}

// Remove tetromino from grid
func remove(grid [][]byte, tetromino []string, row, col int) {
	for i := range tetromino {
		for j := range tetromino[i] {
			// Replace Alphabetic blocks with original 0s
			if tetromino[i][j] != '.' {
				grid[row+i][col+j] = 0
			}
		}
	}
}
