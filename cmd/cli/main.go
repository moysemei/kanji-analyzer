package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/moysemei/kanji-analyzer/internal/analyzer"
	"github.com/moysemei/kanji-analyzer/internal/dictionary"
	"github.com/moysemei/kanji-analyzer/internal/exporter"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

func main() {
	inputPath := flag.String("input", "", "path to the subtitle .srt file")
	dictPath := flag.String("dict", "internal/dictionary/data/jlpt.json", "path to the JLPT dictionary JSON")
	outputPath := flag.String("output", "anki_deck.csv", "path to the output CSV file")

	flag.Parse()

	if *inputPath == "" {
		log.Fatal("Missing required flag: -input")
	}

	fmt.Println("Initializing Kanji Analyzer CLI...")
	fmt.Printf("Reading file: %s\n\n", *inputPath)

	rawContent, err := subtitle.ReadFile(*inputPath)
	if err != nil {
		log.Fatalf("Fatal error reading file: %v", err)
	}

	jlptDict, err := dictionary.Load(*dictPath)
	if err != nil {
		log.Fatalf("Fatal error loading dictionary: %v", err)
	}

	result, err := analyzer.Analyze(rawContent, jlptDict)
	if err != nil {
		log.Fatalf("Fatal error analyzing subtitle: %v", err)
	}

	fmt.Println("=== DENSITY REPORT ===")
	fmt.Printf("Total Valid Words: %d\n", result.Stats.TotalWords)

	for level, percentage := range result.Stats.Density {
		fmt.Printf("- %-7s: %5.2f%%\n", level, percentage)
	}
	fmt.Println("======================")
	fmt.Println()

	fmt.Printf("Generating Anki deck at: %s...\n", *outputPath)

	err = exporter.ToCSV(result.Vocabulary, jlptDict, *outputPath)
	if err != nil {
		log.Fatalf("Fatal error generating CSV: %v", err)
	}

	fmt.Println("Success! Your Anki deck is ready.")

	fmt.Println("--- JLPT VOCABULARY ANALYSIS ---")

	for _, word := range result.Vocabulary {
		level, exists := jlptDict[word]

		if exists {
			fmt.Printf("Word: %-10s | Level: %s\n", word, level)
		} else {
			fmt.Printf("Word: %-10s | Level: Unknown\n", word)
		}
	}

	fmt.Println("---------------------------------")
}
