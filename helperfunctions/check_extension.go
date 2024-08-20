package helperfunctions

func Check_extension(str string) bool {
	var status bool
	start := len(str) - 4
	if str[start:] == ".txt" {
		status = true
	}
	return status
}
