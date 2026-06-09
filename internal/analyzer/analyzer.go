package analyzer

import (
	"github.com/moysemei/kanji-analyzer/internal/nlp"
	"github.com/moysemei/kanji-analyzer/internal/stats"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

type Result struct {
	Stats      stats.Report     `json:"stats"`
	Vocabulary []VocabularyItem `json:"vocabulary"`
}

type VocabularyItem struct {
	Word  string `json:"word"`
	Level string `json:"level"`
}

func Analyze(rawContent string, dict map[string]string) (Result, error) {
	cleanedDialogue := subtitle.CleanSRT(rawContent)
	pureJapanese := subtitle.RemoveNonJapanese(cleanedDialogue)

	words, err := nlp.ExtractVocabulary(pureJapanese)
	if err != nil {
		return Result{}, err
	}

	report := stats.CalculateDensity(words, dict)
	vocabulary := buildVocabularyItems(words, dict)

	return Result{
		Stats:      report,
		Vocabulary: vocabulary,
	}, nil
}

func buildVocabularyItems(words []string, dict map[string]string) []VocabularyItem {
	vocabulary := make([]VocabularyItem, 0, len(words))

	for _, word := range words {
		level, exists := dict[word]
		if !exists {
			level = "Unknown"
		}

		vocabulary = append(vocabulary, VocabularyItem{
			Word:  word,
			Level: level,
		})
	}

	return vocabulary
}
