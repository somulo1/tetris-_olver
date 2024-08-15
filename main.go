package main

import (
	"fmt"
	"os"

	"Tetris-optimizer/helperfunctions"
)

func main() {
	// Limit CLI arguments to 1 (assuming the file path is the only argument)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <something.txt>")
		return
	}
	filePath := os.Args[1]

	// 1. Check if file is empty
	if err := helperfunctions.CheckIfFileIsEmpty(filePath); err != nil {
		fmt.Println(err)
		return
	}

	// 2. Read file content
	content, err := helperfunctions.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3. Parse tetrominoes
	tetrominoes, err := helperfunctions.CheckTetris(content)
	if err != nil {
		fmt.Println("Error parsing tetrominoes")
		return
	}

	// 4. Validate tetromino format
	for _, tetromino := range tetrominoes {
		if !helperfunctions.ValidateTetrominoFormat(tetromino) {
			fmt.Println("ERROR: Invalid tetromino format")
			return
		}
	}

	// Example board dimensions
	boardWidth := 10
	boardHeight := 10

	// 5. Execute board operations for each tetromino and print results
	for _, tetromino := range tetrominoes {
		// Call ExecuteBoardOperations with example coordinates (0, 0) and print the result
		boardRepresentation := helperfunctions.ExecuteBoardOperations(boardWidth, boardHeight, *tetrominoes, 0, 0)
		fmt.Println(boardRepresentation)
	}
}
