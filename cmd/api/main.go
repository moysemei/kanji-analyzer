package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/moysemei/kanji-analyzer/internal/dictionary"
	"github.com/moysemei/kanji-analyzer/internal/nlp"
	"github.com/moysemei/kanji-analyzer/internal/stats"
	"github.com/moysemei/kanji-analyzer/internal/subtitle"
)

type APIResponse struct {
	Stats      stats.Report `json:"stats"`
	Vocabulary []string     `json:"vocabulary"`
}

func main() {
	port := ":8080"
	dictPath := "../../internal/dictionary/data/jlpt.json"

	jlptDict, err := dictionary.Load(dictPath)
	if err != nil {
		log.Fatalf("Fatal error loading dictionary: %v", err)
	}
	fmt.Println("Dictionary loaded into memory.")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	http.HandleFunc("/api/analyze", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		r.ParseMultipartForm(10 << 20)

		file, _, err := r.FormFile("subtitle")
		if err != nil {
			http.Error(w, `{"error": "Failed to receive subtitle file"}`, http.StatusBadRequest)
			return
		}
		defer file.Close()

		bytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, `{"error": "Failed to read file bytes"}`, http.StatusInternalServerError)
			return
		}

		rawContent := string(bytes)

		cleanedDialogue := subtitle.CleanSRT(rawContent)
		pureJapanese := subtitle.RemoveNonJapanese(cleanedDialogue)

		vocab, err := nlp.ExtractVocabulary(pureJapanese)
		if err != nil {
			http.Error(w, `{"error": "Failed to process NLP engine"}`, http.StatusInternalServerError)
			return
		}

		report := stats.CalculateDensity(vocab, jlptDict)

		response := APIResponse{
			Stats:      report,
			Vocabulary: vocab,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	fmt.Printf("Starting HTTP server on http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
