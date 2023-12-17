package file

import (
	"fmt"
	"os"
)

// ReadFile reads the content from the specified file
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, create it
			if err := os.WriteFile(filePath, []byte(""), 0644); err != nil {
				return "", fmt.Errorf("error creating file: %v", err)
			}
			return "", nil
		}
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(data), nil
}

// WriteFile writes the given content to the specified file
func WriteFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}
