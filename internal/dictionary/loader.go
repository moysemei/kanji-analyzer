// Package dictionary handles loading and querying JLPT vocabulary data.
package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(filepath string) (map[string]string, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dictionary file: %w", err)
	}

	var dict map[string]string

	err = json.Unmarshal(bytes, &dict)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return dict, nil
}
