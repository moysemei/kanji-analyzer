package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Meta struct {
	Reading   string `json:"reading"`
	Frequency struct {
		DisplayValue string `json:"displayValue"`
	} `json:"frequency"`
}

func main() {
	sourceDir := flag.String("source", "internal/dictionary/source", "directory containing Yomitan term_meta_bank files")
	outputPath := flag.String("output", "internal/dictionary/data/jlpt.json", "output JLPT dictionary path")

	flag.Parse()

	dict := make(map[string]string)

	files, err := filepath.Glob(filepath.Join(*sourceDir, "term_meta_bank_*.json"))
	if err != nil {
		log.Fatalf("failed to find source files: %v", err)
	}

	if len(files) == 0 {
		log.Fatalf("no term_meta_bank_*.json files found in %s", *sourceDir)
	}

	for _, file := range files {
		if err := processFile(file, dict); err != nil {
			log.Fatalf("failed to process %s: %v", file, err)
		}
	}

	output, err := json.MarshalIndent(dict, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal dictionary: %v", err)
	}

	if err := os.WriteFile(*outputPath, output, 0644); err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	fmt.Printf("Generated dictionary with %d entries at %s\n", len(dict), *outputPath)
}

func processFile(path string, dict map[string]string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var entries [][]json.RawMessage
	if err := json.Unmarshal(bytes, &entries); err != nil {
		return err
	}

	for _, entry := range entries {
		if len(entry) < 3 {
			continue
		}

		var word string
		if err := json.Unmarshal(entry[0], &word); err != nil {
			continue
		}

		var meta Meta
		if err := json.Unmarshal(entry[2], &meta); err != nil {
			continue
		}

		level := meta.Frequency.DisplayValue
		if level == "" {
			continue
		}

		addWord(dict, word, level)

		// Also add the reading as a key.
		// This helps when subtitles use kana instead of kanji.
		if meta.Reading != "" {
			addWord(dict, meta.Reading, level)
		}
	}

	return nil
}

func addWord(dict map[string]string, word string, level string) {
	if word == "" {
		return
	}

	currentLevel, exists := dict[word]
	if !exists {
		dict[word] = level
		return
	}

	if isEasierLevel(level, currentLevel) {
		dict[word] = level
	}
}

func isEasierLevel(newLevel string, currentLevel string) bool {
	priority := map[string]int{
		"N5": 5,
		"N4": 4,
		"N3": 3,
		"N2": 2,
		"N1": 1,
	}

	return priority[newLevel] > priority[currentLevel]
}
