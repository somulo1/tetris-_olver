package main

import (
	"fmt"
	"os"

	"Tetris-optimizer/helperfunctions"
)

func main() {
	// Limit CLI arguments to 2  argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <something.txt>")
		return
	}
	filePath := os.Args[1]

	// limit file type to a text file
	if !helperfunctions.Check_extension(filePath) {
		fmt.Println("Only provide a .txt file")
		return
	}
	//  Check if file is empty
	if err := helperfunctions.CheckIfFileIsEmpty(filePath); err != nil {
		fmt.Println("ERROR")
		return
	}

	//  Read file content
	content, err := helperfunctions.ReadFile(filePath)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	//  Parse tetrominoes
	tetrominoes := helperfunctions.ValidateTetrominoFormat(content)
	fmt.Println(tetrominoes)
	// Validate tetromino format
	for _, tetromino := range tetrominoes {
		if !helperfunctions.CheckTetri(tetromino) {
			fmt.Println("ERROR1")
			return
		}
		helperfunctions.CheckTetris
	}
	// trim the tetromino
	// placing the tetris
}
