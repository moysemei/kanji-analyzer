package analyzer

import "testing"

func TestAnalyzeReturnsVocabularyAndStats(t *testing.T) {
	rawContent := `1
00:00:01,000 --> 00:00:03,000
俺は海賊王になる男だ
`

	dict := map[string]string{
		"俺":  "N5",
		"海賊": "N3",
		"王":  "N5",
		"なる": "N5",
		"男":  "N3",
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
