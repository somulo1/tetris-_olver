package helperfunction

func CalculateSquareSize(totalArea int) int {
	side := 1
	for side*side < totalArea {
		side++
	}
	return side
}
