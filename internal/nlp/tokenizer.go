// Package nlp handles Natural Language Processing tasks for Japanese text.
package nlp

import (
	"fmt"
	"unicode"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var ignoredWords = map[string]bool{
	"ん": true,
	"の": true,
	"こ": true,
	"っ": true,
	"ー": true,
	"！": true,
	"？": true,
	"〜": true,
}

func ExtractVocabulary(text string) ([]string, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize kagome tokenizer: %w", err)
	}

	tokens := t.Tokenize(text)

	var vocabulary []string
	seen := make(map[string]bool)

	for _, token := range tokens {
		pos := token.POS()
		if len(pos) == 0 {
			continue
		}

		mainPOS := pos[0]

		baseWord, ok := token.BaseForm()
		if !ok || baseWord == "" {
			baseWord = token.Surface
		}

		if !isValidVocabulary(baseWord, mainPOS) {
			continue
		}

		if !seen[baseWord] {
			seen[baseWord] = true
			vocabulary = append(vocabulary, baseWord)
		}
	}

	return vocabulary, nil
}

func isValidVocabulary(word string, mainPOS string) bool {
	if word == "" {
		return false
	}

	if ignoredWords[word] {
		return false
	}

	if mainPOS == "記号" || mainPOS == "助詞" || mainPOS == "助動詞" || mainPOS == "フィラー" {
		return false
	}

	if !containsJapanese(word) {
		return false
	}

	return true
}

func containsJapanese(word string) bool {
	for _, r := range word {
		if unicode.In(r, unicode.Hiragana, unicode.Katakana, unicode.Han) {
			return true
		}
	}

	return false
}
