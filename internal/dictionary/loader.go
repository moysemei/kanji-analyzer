// Package dictionary handles loading and querying JLPT vocabulary data.
package dictionary

import (
	"encoding/json"
	"fmt"
	"os"
)

type Entry struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
	Level   string `json:"level"`
}

func Load(filepath string) (map[string]Entry, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dictionary file: %w", err)
	}

	var dict map[string]Entry

	err = json.Unmarshal(bytes, &dict)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return dict, nil
}
