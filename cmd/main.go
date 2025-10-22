package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/VasySS/project-stas/internal/components"
)

type TimelineDate struct {
	Year   string `json:"year"`
	Month  string `json:"month,omitempty"`
	Day    string `json:"day,omitempty"`
	Hour   string `json:"hour,omitempty"`
	Minute string `json:"minute,omitempty"`
}

type TimelineText struct {
	Headline string `json:"headline"`
	Text     string `json:"text"`
}

type TimelineEvent struct {
	StartDate TimelineDate `json:"start_date"`
	Text      TimelineText `json:"text"`
}

type TimelineTitle struct {
	Text TimelineText `json:"text"`
}

type TimelineData struct {
	Title  TimelineTitle   `json:"title"`
	Events []TimelineEvent `json:"events"`
}

func main() {
	loreData, err := loadTimelineData("static/lore.json")
	if err != nil {
		log.Fatalf("failed to load timeline data: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = components.Layout("Timeline Demo", components.Home()).Render(r.Context(), w)
	})

	mux.HandleFunc("/timeline", func(w http.ResponseWriter, r *http.Request) {
		_ = components.TimelineFragment().Render(r.Context(), w)
	})

	mux.HandleFunc("/api/timeline-data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(loreData)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func loadTimelineData(path string) (TimelineData, error) {
	f, err := os.Open(path)
	if err != nil {
		return TimelineData{}, err
	}
	defer f.Close()

	var data TimelineData
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return TimelineData{}, err
	}

	return data, nil
}
