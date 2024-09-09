package main

import "fmt"

// Constants for Tetrimino data processing
const (
	blockSize = 20   // Size of each Tetrimino block in bytes
	dot       = '.'  // Character representing an empty cell
	hashTag   = '#'  // Character representing a filled cell in a Tetrimino
	newLine   = '\n' // Character representing a newline
)

// isValidBlock checks if a block of data (20 bytes) is a valid Tetrimino block.
// Returns true if the block contains exactly 12 dots and 4 hashtags,
// with new lines in the correct positions.
func isValidBlock(data []byte) bool {
	numDots, numHashtags := 0, 0

	for i, v := range data {
		switch v {
		case dot:
			numDots++
		case hashTag:
			numHashtags++
		case newLine:
			if i%5 != 4 {
				return false
			}
		default:
			return false
		}
	}
	return numDots == 12 && numHashtags == 4
}

// validateTetris processes the entire data slice to extract and validate Tetriminos.
// Returns a boolean indicating if the data is valid and a slice of valid Tetriminos.
func validateTetris(data []byte) (ok bool, result []tetrimino) {
	for i := 0; i < len(data); {
		if i+blockSize > len(data) {
			break
		}
		if !isValidBlock(data[i : i+blockSize]) {
			fmt.Println("invalid block")
			break
		}
		t := blockToTetrimino(data[i : i+blockSize])
		if !t.isValid() {
			fmt.Println("invalid Tetrimino")
			break
		}
		result = append(result, t)
		i += blockSize
		if i == len(data) {
			ok = true
			break
		}
		if data[i] != '\n' {
			break
		}
		i++
	}
	return ok, result
}
