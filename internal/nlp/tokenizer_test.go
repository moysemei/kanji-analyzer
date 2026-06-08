package nlp

import "testing"

func TestContainsJapanese(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{name: "kanji", word: "海賊", want: true},
		{name: "hiragana", word: "なる", want: true},
		{name: "katakana", word: "ナレーション", want: true},
		{name: "number", word: "2", want: false},
		{name: "punctuation", word: "！", want: false},
		{name: "latin", word: "hello", want: false},
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
