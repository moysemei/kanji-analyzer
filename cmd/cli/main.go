package main

import (
	"fmt"
	"log"

	"github.com/moysemei/kanji-analyzer/internal/nlp"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

func main() {
	filepath := "../../test_anime.srt"

	fmt.Println("Initializing Kanji Analyzer CLI...")
	fmt.Printf("Reading file: %s\n\n", filepath)

	rawContent, err := subtitle.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
	}

	cleanedDialogue := subtitle.CleanSRT(rawContent)

	pureJapanese := subtitle.RemoveNonJapanese(cleanedDialogue)

	vocabulary, err := nlp.ExtractVocabulary(pureJapanese)
	if err != nil {
		log.Fatalf("Fatal error in NLP engine: %v", err)
	}

	fmt.Printf("--- EXTRACTED VOCABULARY (%d words) ---\n", len(vocabulary))
	for i, word := range vocabulary {
		fmt.Printf("%d: %s\n", i+1, word)
	}
	fmt.Println("---------------------------------------")
}
