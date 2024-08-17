package helperfunctions

import "math"

func Solve(tetroSlc [][]string) int {
	// Find size of smallest possible grid to fit all tetrominoes
	size := math.Sqrt(float64(len(tetroSlc) * 4))
	return int(math.Ceil(size))
}
