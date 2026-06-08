package analyzer

import (
	"github.com/moysemei/kanji-analyzer/internal/nlp"
	"github.com/moysemei/kanji-analyzer/internal/stats"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

type Result struct {
	Stats      stats.Report `json:"stats"`
	Vocabulary []string     `json:"vocabulary"`
}

func Analyze(rawContent string, dict map[string]string) (Result, error) {
	cleanedDialogue := subtitle.CleanSRT(rawContent)
	pureJapanese := subtitle.RemoveNonJapanese(cleanedDialogue)

	vocabulary, err := nlp.ExtractVocabulary(pureJapanese)
	if err != nil {
		return Result{}, err
	}

	report := stats.CalculateDensity(vocabulary, dict)

	return Result{
		Stats:      report,
		Vocabulary: vocabulary,
	}, nil
}
