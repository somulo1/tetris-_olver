package validity

//check if it has four hashes
func FourHashes(tetro []string) bool {
	var count int
	row := len(tetro)
	col := len(tetro[0])

	// j ranges though lines of tetromino
	// k ranges through character in each line
	for j := 0; j < row; j++ {
		for k := 0; k < col; k++ {
			// Count number of '#' in tetromino
			if string(tetro[j][k]) == "#" {
				count++
			}
		}
	}

	return count == 4
}

// check for the valid tetromino representations in ech of the four chuncks
func ValidConnections(tet []string) bool {
	var count int
	tag := byte('#')
	row, col := len(tet), len(tet[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if tet[i][j] == tag {
				// Check above
				if i > 0 && tet[i-1][j] == tag {
					count++
				}
				// Check below
				if i < row-1 && tet[i+1][j] == tag {
					count++
				}
				// Check left
				if j > 0 && tet[i][j-1] == tag {
					count++
				}
				// Check right
				if j < col-1 && tet[i][j+1] == tag {
					count++
				}

			}
		}
	}
	// Return true should 'count' be 6 or 8
	return count == 6 || count == 8
}

 //check the tetromino representations
 func FourByFour(tetro []string) bool {
	valid := true

	// Determine if tetromino is stacked four lines high
	if len(tetro) != 4 {
		valid = false
	} else {
		// Determine if each line has four characters
		for _, line := range tetro {
			if len(line) != 4 {
				valid = false
				break
			}
		}
	}
	return valid
}
