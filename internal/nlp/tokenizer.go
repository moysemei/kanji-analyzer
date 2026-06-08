// Package nlp handles Natural Language Processing tasks for Japanese text.
package nlp

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

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

		if mainPOS == "記号" || mainPOS == "助詞" || mainPOS == "助動詞" || mainPOS == "フィラー" {
			continue
		}

		baseWord, ok := token.BaseForm()
		if !ok || baseWord == "" {
			baseWord = token.Surface
		}

		if baseWord != "" && !seen[baseWord] {
			seen[baseWord] = true
			vocabulary = append(vocabulary, baseWord)
		}
	}

	return vocabulary, nil
}
