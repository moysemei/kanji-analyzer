package analyzer

import (
	"testing"

	"github.com/moysemei/kanji-analyzer/internal/dictionary"
)

func TestAnalyzeReturnsVocabularyAndStats(t *testing.T) {
	rawContent := `1
00:00:01,000 --> 00:00:03,000
俺は海賊王になる男だ
`

	dict := map[string]dictionary.Entry{
		"俺":  {Word: "俺", Reading: "おれ", Level: "N1"},
		"海賊": {Word: "海賊", Reading: "かいぞく", Level: "N3"},
		"王":  {Word: "王", Reading: "おう", Level: "N3"},
		"なる": {Word: "なる", Reading: "なる", Level: "N5"},
		"男":  {Word: "男", Reading: "おとこ", Level: "N5"},
	}

	result, err := Analyze(rawContent, dict)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Stats.TotalWords == 0 {
		t.Fatal("expected total words to be greater than 0")
	}

	if len(result.Vocabulary) == 0 {
		t.Fatal("expected vocabulary to not be empty")
	}
}
