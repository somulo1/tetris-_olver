package helperfunctions

import (
	"fmt"
	"strings"
)

func CheckTetris(input string) []string {
	var tetromino []string // This will store multiple tetrominos as single strings
	var tetro [][]string   // 2D slice for the current tetromino
	var letter rune = 'A'  // Start with letter 'A'
	var new string         // Holds the processed character

	// Split the file content by newline characters
	newinput := strings.Split(string(input), "\n")

	for _, line := range newinput {
		// If the line is empty, process the current tetromino
		if line == "" {
			if len(tetro) > 0 {
				// Convert the current tetro (2D slice) to a single string
				var tetroString strings.Builder
				for _, row := range tetro {
					tetroString.WriteString(strings.Join(row, ""))
					tetroString.WriteString("\n")
				}
				tetromino = append(tetromino, tetroString.String())
				// Reset tetro for the next tetromino
				tetro = nil
			}
			continue
		}

		// Process each line
		var processedLine []string
		for _, char := range line {
			if char == '#' {
				// Replace '#' with the current letter
				new = string(letter)
				// Increment the letter; wrap around if needed
				letter++
				if letter > 'Z' {
					letter = 'A'
				}
			} else if char == '.' {
				// Keep '.' as is
				new = "."
			} else {
				// Handle invalid characters
				fmt.Println("ERROR")
				return nil

			}
			// Append the processed character to the current line
			processedLine = append(processedLine, new)
		}

		// Append the processed line to the current tetro (2D slice)
		tetro = append(tetro, processedLine)
	}

	// Handle the last tetromino if no trailing empty line is present
	if len(tetro) > 0 {
		var tetroString strings.Builder
		for _, row := range tetro {
			tetroString.WriteString(strings.Join(row, ""))
			tetroString.WriteString("\n")
		}
		tetromino = append(tetromino, tetroString.String())
	}

	// For demonstration purposes: Print the result
	return tetromino
}
