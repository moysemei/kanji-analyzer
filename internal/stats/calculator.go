// Package stats handle mathematical calculations for vocabulary density.
package stats

import "github.com/moysemei/kanji-analyzer/internal/dictionary"

type Report struct {
	TotalWords int                `json:"totalWords"`
	LevelCount map[string]int     `json:"levelCount"`
	Density    map[string]float64 `json:"density"`
}

func CalculateDensity(vocabulary []string, dict map[string]dictionary.Entry) Report {
	report := Report{
		TotalWords: len(vocabulary),
		LevelCount: make(map[string]int),
		Density:    make(map[string]float64),
	}

	if report.TotalWords == 0 {
		return report
	}

	for _, word := range vocabulary {
		entry, exists := dict[word]
		level := "Unknown"
		if exists {
			level = entry.Level
		}

		report.LevelCount[level]++
	}

	for level, count := range report.LevelCount {
		percentage := (float64(count) / float64(report.TotalWords)) * 100

		report.Density[level] = percentage
	}

	return report
}
