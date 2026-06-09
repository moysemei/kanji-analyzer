package stats

import "testing"

func TestCalculateDensity(t *testing.T) {
	vocabulary := []string{"学校", "悪魔", "王", "男", "ナレーション"}

	dict := map[string]string{
		"学校": "N5",
		"悪魔": "N3",
		"王":  "N3",
		"男":  "N5",
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
