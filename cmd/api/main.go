package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/moysemei/kanji-analyzer/internal/analyzer"
	"github.com/moysemei/kanji-analyzer/internal/dictionary"
	"github.com/moysemei/kanji-analyzer/internal/stats"
)

type APIResponse struct {
	Stats      stats.Report `json:"stats"`
	Vocabulary []string     `json:"vocabulary"`
}

func main() {
	port := flag.String("port", "8080", "HTTP server port")
	dictPath := flag.String("dict", "internal/dictionary/data/jlpt.json", "path to the JLPT dictionary JSON")

	flag.Parse()

	jlptDict, err := dictionary.Load(*dictPath)
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
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, `{"error": "Failed to parse multipart form"}`, http.StatusBadRequest)
			return
		}

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

		result, err := analyzer.Analyze(rawContent, jlptDict)
		if err != nil {
			http.Error(w, `{"error": "Failed to process subtitle"}`, http.StatusInternalServerError)
			return
		}

		response := APIResponse{
			Stats:      result.Stats,
			Vocabulary: result.Vocabulary,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	addr := ":" + *port

	fmt.Printf("Starting HTTP server on http://localhost%s\n", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
