package helperfunctions

func hasSuffix(s, suffix string) bool {
	return s[len(s)-4:] == suffix
}

func isValidFile(fileName string) bool {
	return hasSuffix(fileName, ".txt")
}
