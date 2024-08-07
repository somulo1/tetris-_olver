package helperfunction

import (
	"fmt"
	"os"
)

// 1.1 Check if the file is empty
func CheckIfFileIsEmpty(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		return fmt.Errorf("file is empty")
	}
	return nil
}
