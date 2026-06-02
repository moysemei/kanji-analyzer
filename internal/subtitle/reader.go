// Package subtitle handles the reading, parsing, and cleaning of subtitle files.
package subtitle

import (
	"fmt"
	"os"
)

func ReadFile(filepath string) (string, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	return string(bytes), nil
}
