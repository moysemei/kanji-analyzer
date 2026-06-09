package nlp

import "testing"

func TestIsValidVocabularyFiltersIgnoredWords(t *testing.T) {
	ignored := []string{"ん", "の", "こ", "っ", "ー", "！", "？", "〜"}

	for _, word := range ignored {
		t.Run(word, func(t *testing.T) {
			if isValidVocabulary(word, "名詞") {
				t.Fatalf("expected %q to be invalid", word)
			}
		})
	}
}

func TestIsValidVocabularyKeepsValidShortWords(t *testing.T) {
	validWords := []string{"王", "人", "犬", "猫", "夢", "心"}

	for _, word := range validWords {
		t.Run(word, func(t *testing.T) {
			if !isValidVocabulary(word, "名詞") {
				t.Fatalf("expected %q to be valid", word)
			}
		})
	}
}

func TestContainsJapanese(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "kanji", word: "悪魔", want: true},
		{name: "hiragana", word: "なる", want: true},
		{name: "katakana", word: "ナレーション", want: true},
		{name: "latin", word: "hello", want: false},
		{name: "number", word: "123", want: false},
		{name: "punctuation", word: "!", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsJapanese(tt.word)
			if got != tt.want {
				t.Fatalf("containsJapanese(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
