package subtitle

import (
	"regexp"
	"strings"
)

func CleanSRT(rawContent string) string {
	pattern := `(?m)^\d+\r?\n^\d{2}:\d{2}:\d{2},\d{3} --> \d{2}:\d{2}:\d{2},\d{3}\r?\n`
	re := regexp.MustCompile(pattern)

	cleaned := re.ReplaceAllString(rawContent, "")

	return strings.TrimSpace(cleaned)
}

func RemoveNonJapanese(text string) string {
	pattern := `[a-zA-ZÀ-ÿ0-9\s()!?,.\-]+`
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllString(text, "")
}
