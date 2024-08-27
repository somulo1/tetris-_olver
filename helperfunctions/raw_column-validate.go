package helperfunctions

import "Tetris-optimizer/helperfunctions/validity"

// ValidateTetromino validates a tetromino by calling the necessary functions and checks.
func ValidateTetromino(tet []string) bool {
	// Trim the tetromino to remove unnecessary columns and rows
	trimmedTetromino := TetroTrim(tet)

	// Check if the trimmed tetromino is a valid 4x4 representation
	if !validity.FourByFour(trimmedTetromino) {
		return false
	}

	// Check if the trimmed tetromino contains exactly four hashes
	if !validity.FourHashes(trimmedTetromino) {
		return false
	}

	// Check if the hashes in the tetromino are connected properly
	if !validity.ValidConnections(trimmedTetromino) {
		return false
	}

	// If all checks pass, return true
	return true
}
