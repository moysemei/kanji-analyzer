package main

import (
	"fmt"
	"log"

	"github.com/moysemei/kanji-analyzer/internal/dictionary"
	"github.com/moysemei/kanji-analyzer/internal/exporter"
	"github.com/moysemei/kanji-analyzer/internal/nlp"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

func main() {
	subtitlePath := "../../test_anime.srt"
	dictPath := "../../internal/dictionary/data/jlpt.json"

	fmt.Println("Initializing Kanji Analyzer CLI...")
	fmt.Printf("Reading file: %s\n\n", subtitlePath)

	rawContent, err := subtitle.ReadFile(subtitlePath)
	if err != nil {
		log.Fatalf("Fatal error reading file: %v", err)
	}

	cleanedDialogue := subtitle.CleanSRT(rawContent)
	pureJapanese := subtitle.RemoveNonJapanese(cleanedDialogue)

	vocabulary, err := nlp.ExtractVocabulary(pureJapanese)
	if err != nil {
		log.Fatalf("Fatal error in NLP engine: %v", err)
	}

	jlptDict, err := dictionary.Load(dictPath)
	if err != nil {
		log.Fatalf("Fatal error loading dictionary: %v", err)
	}

	outputFile := "../../deck_test_anime.csv"
	fmt.Printf("Generating Anki deck at: %s...\n", outputFile)

	err = exporter.ToCSV(vocabulary, jlptDict, outputFile)
	if err != nil {
		log.Fatalf("Fatal error generating CSV: %v", err)
	}

	fmt.Println("Success! Your Anki deck is ready.")

	fmt.Println("--- JLPT VOCABULARY ANALYSIS ---")

	for _, word := range vocabulary {
		level, exists := jlptDict[word]

		if exists {
			fmt.Printf("Word: %-10s | Level: %s\n", word, level)
		} else {
			fmt.Printf("Word: %-10s | Level: Unknown\n", word)
		}
	}

	fmt.Println("---------------------------------")
}
