package helperfunctions

import "strings"

// Reduces tetromino footprint, removing unnecesary '.'
func TetroTrim(tet []string) []string {
	var trimmed []string
	var line string
	index := occupyCol(tet) // Holds indices of columns to retain
	row := len(tet)

	// Range through tetromino row-by-row removing empty columns
	for i := 0; i < row; i++ {
		for j, ind := range index {
			if j != len(index)-1 {
				// Populate line with characters that will make retained columns
				line += string(tet[i][ind])
			} else {
				// At the final index...
				// Populate line with last character
				line += string(tet[i][ind])
				trimmed = append(trimmed, line) // Append filled line to trimmed tetromino
				line = ""                       // Empty for next round of population
			}
		}
	}

	// Trim empty rows
	trimmed = removeRow(trimmed)

	return trimmed
}

func removeRow(tet []string) []string {
	var result []string

	// Append rows that contain '#'s
	for _, row := range tet {
		if strings.Contains(row, "#") {
			result = append(result, row)
		}
	}
	return result
}

func occupyCol(tet []string) []int {
	var index []int
	tag := byte('#')
	row, col := len(tet), len(tet[0])

	// Find indices in rows that contain '#'s
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if tet[j][i] == tag {

				// Store index of in index
				index = append(index, i)
				break
			}
		}
	}
	return index
}
