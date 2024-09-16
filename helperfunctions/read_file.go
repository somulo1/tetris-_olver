package helperfunctions

import (
	"bufio"
	"errors"
	"os"
)

var (
	ErrInvalidTetroSize = errors.New("ERROR")
	ErrInvalidTetroFile = errors.New("ERROR")
	ErrInvalidTetroType = errors.New("ERROR")
)

// reads input file and returns pointer to a Tetrominos and an error
func InputFileReader(fileName string) (*Tetrominos, error) {
	if !isValidFile(fileName) {
		return nil, ErrInvalidTetroFile
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		tetrominos     [][]string //2 D terominos
		tetromino      []string   // 1 D tetromino
		linesRead      int
		tetrominoLabel = 'A' // letter to assign
	)

	for scanner.Scan() {
		line := scanner.Text()

		// Check for end of tetromino block (expecting a blank line after every 4th line)
		if linesRead == 4 {
			if line != "" {
				return nil, ErrInvalidTetroSize
			}
			tetrominos = append(tetrominos, tetromino)
			tetromino = nil
			linesRead = 0
			tetrominoLabel++
			continue
		}

		// Ensure tetromino does not exceed 26 pieces
		if tetrominoLabel > 'Z' {
			return nil, ErrInvalidTetroFile
		}

		// Validate the line contains only '.' or '#'
		for _, ch := range line {
			if ch != '.' && ch != '#' {
				return nil, ErrInvalidTetroType
			}
		}

		// Replace '#' with the corresponding label
		newLine := ""
		for _, ch := range line {
			if ch == '#' {
				newLine += string(tetrominoLabel)
			} else {
				newLine += string(ch)
			}
		}

		// Add the line to the current tetromino
		tetromino = append(tetromino, newLine)
		linesRead++
	}

	// Final check for unprocessed tetrominos or incomplete file
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if linesRead != 0 {
		return nil, ErrInvalidTetroSize
	}
	if len(tetrominos) == 0 {
		return nil, ErrInvalidTetroFile
	}

	return &Tetrominos{Tet: tetrominos}, nil
}
