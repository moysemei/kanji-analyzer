// Package stats handle mathematical calculations for vocabulary density.
package stats

type Report struct {
	TotalWords int
	LevelCount map[string]int
	Density    map[string]float64
}

func CalculateDensity(vocabulary []string, dict map[string]string) Report {
	report := Report{
		TotalWords: len(vocabulary),
		LevelCount: make(map[string]int),
		Density:    make(map[string]float64),
	}

	if report.TotalWords == 0 {
		return report
	}

	for _, word := range vocabulary {
		level, exists := dict[word]
		if !exists {
			level = "Unknown"
		}

		report.LevelCount[level]++
	}

	for level, count := range report.LevelCount {
		percentage := (float64(count) / float64(report.TotalWords)) * 100

		report.Density[level] = percentage
	}

	return report
}
