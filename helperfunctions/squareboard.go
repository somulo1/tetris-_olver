package helperfunctions

import (
	"strings"
)

// Define a type for the board and tetrominoes
type Board struct {
	width  int
	height int
	grid   [][]rune
}

type Tetromino struct {
	width  int
	height int
	shape  [][]rune
}

// Create a new board with a given width and height
func NewBoard(width, height int) *Board {
	board := &Board{width: width, height: height}
	board.grid = make([][]rune, height)
	for i := range board.grid {
		board.grid[i] = make([]rune, width)
		for j := range board.grid[i] {
			board.grid[i][j] = '.'
		}
	}
	return board
}

// Print the board and return it as a string
func (b *Board) Print() string {
	var builder strings.Builder
	for _, row := range b.grid {
		for _, cell := range row {
			builder.WriteString(string(cell) + " ")
		}
		builder.WriteString("\n")
	}
	builder.WriteString("\n")
	return builder.String()
}

// Adjust the board size if a tetromino does not fit
func (b *Board) AdjustSizeIfNeeded(tetromino *Tetromino) {
	if tetromino.width > b.width || tetromino.height > b.height {
		newWidth := max(b.width, tetromino.width)
		newHeight := max(b.height, tetromino.height)
		newBoard := NewBoard(newWidth, newHeight)
		for i := 0; i < b.height; i++ {
			for j := 0; j < b.width; j++ {
				newBoard.grid[i][j] = b.grid[i][j]
			}
		}
		b.width = newWidth
		b.height = newHeight
		b.grid = newBoard.grid
	}
}

// Utility function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Place a tetromino on the board (simplified for demonstration)
func (b *Board) PlaceTetromino(tetromino *Tetromino, x, y int) {
	for i := 0; i < tetromino.height; i++ {
		for j := 0; j < tetromino.width; j++ {
			if x+i < b.height && y+j < b.width {
				b.grid[x+i][y+j] = tetromino.shape[i][j]
			}
		}
	}
}

// Manage the entire process of board creation, adjustment, and tetromino placement
func ExecuteBoardOperations(boardWidth, boardHeight int, tetromino *Tetromino, x, y int) string {
	// Create the board
	board := NewBoard(boardWidth, boardHeight)

	// Adjust the board size if needed
	board.AdjustSizeIfNeeded(tetromino)

	// Place the tetromino on the board
	board.PlaceTetromino(tetromino, x, y)

	// Return the board as a string
	return board.Print()
}
