package main

import (
	"fmt"
	"os"
	"strings"
	"tetris-optimizer/helperfunction"
)	

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <file-path>")
		return
	}

	filePath := os.Args[1]

	// 1. Check if file is empty
	if err := helperfunction.CheckIfFileIsEmpty(filePath); err != nil {
		fmt.Println(err)
		return
	}

	// 2. Read file content
	content, err := helperfunction.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3. Parse tetrominoes
	tetrominoes, err := helperfunction.ParseTetrominoes(content)
	if err != nil {
		fmt.Println("Error parsing tetrominoes")
		return
	}

	// 4. Validate tetromino format
	for _, tetromino := range tetrominoes {
		if !helperfunction.ValidateTetrominoFormat(tetromino) {
			fmt.Println("ERROR")
			return
		}
	}

	// 5. Calculate grid size
	totalArea := 0
	for _, tetromino := range tetrominoes {
		totalArea += strings.Count(tetromino, "#")
	}
	gridSize := helperfunction.CalculateSquareSize(totalArea)

	// 6. Place tetrominoes into the grid
	grid := helperfunction.PlaceTetrominoes(tetrominoes, gridSize)

	// 7. Print the result
	helperfunction.PrintGrid(grid)
}
