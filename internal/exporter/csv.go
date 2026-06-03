// Package exporter handles formatting and saving data to external files.
package exporter

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ToCSV(vocabulary []string, dict map[string]string, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create csv file: %w", err)
	}

	defer file.Close()

	file.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, word := range vocabulary {
		level, exists := dict[word]
		if !exists {
			level = "Unknown"
		}

		row := []string{word, level}

		err := writer.Write(row)
		if err != nil {
			return fmt.Errorf("failed to write row to csv: %w", err)
		}
	}

	return nil
}
