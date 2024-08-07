package helperfunction

import "strings"

func ParseTetrominoes(content string) ([]string, error) {
	lines := strings.Split(content, "\n")
	var tetrominoes []string
	var currentTetromino strings.Builder

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			if currentTetromino.Len() > 0 {
				tetrominoes = append(tetrominoes, currentTetromino.String())
				currentTetromino.Reset()
			}
		} else {
			currentTetromino.WriteString(line + "\n")
		}
	}

	// Append last tetromino if there is any
	if currentTetromino.Len() > 0 {
		tetrominoes = append(tetrominoes, currentTetromino.String())
	}

	return tetrominoes, nil
}
