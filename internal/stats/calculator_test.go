package stats

import (
	"testing"

	"github.com/moysemei/kanji-analyzer/internal/dictionary"
)

func TestCalculateDensity(t *testing.T) {
	vocabulary := []string{"学校", "悪魔", "王", "男", "ナレーション"}

	dict := map[string]dictionary.Entry{
		"学校": {Word: "学校", Reading: "がっこう", Level: "N5"},
		"悪魔": {Word: "悪魔", Reading: "あくま", Level: "N3"},
		"王":  {Word: "王", Reading: "おう", Level: "N3"},
		"男":  {Word: "男", Reading: "おとこ", Level: "N5"},
	}

	report := CalculateDensity(vocabulary, dict)

	if report.TotalWords != 5 {
		t.Fatalf("expected total words 5, got %d", report.TotalWords)
	}

	if report.LevelCount["N5"] != 2 {
		t.Fatalf("expected N5 count 2, got %d", report.LevelCount["N5"])
	}

	if report.LevelCount["N3"] != 2 {
		t.Fatalf("expected N3 count 2, got %d", report.LevelCount["N3"])
	}

	if report.LevelCount["Unknown"] != 1 {
		t.Fatalf("expected Unknown count 1, got %d", report.LevelCount["Unknown"])
	}

	if report.Density["N5"] != 40 {
		t.Fatalf("expected N5 density 40, got %.2f", report.Density["N5"])
	}
}
