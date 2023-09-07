package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Task struct {
	Slack_name      string    `json:"slack_name"`
	Current_day     string    `json:"current_day"`
	Utc_time        time.Time `json:"utc_time"`
	Track           string    `json:"track"`
	Github_file_url string    `json:"github_file_url"`
	Github_repo_url string    `json:"github_repo_url"`
	Status_code     int       `json:"status_code"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "John's hngx task 1")
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		slack_name := r.URL.Query().Get("slack_name")
		track := r.URL.Query().Get("track")

		task1 := Task{
			Slack_name:      slack_name,
			Current_day:     time.Now().Weekday().String(),
			Utc_time:        time.Now(),
			Track:           track,
			Github_file_url: "github.com/jenweani/",
			Github_repo_url: "github.com/jenweani/",
			Status_code:     http.StatusOK,
		}

		json.NewEncoder(w).Encode(task1)
	})

	http.ListenAndServe(":8080", nil)
}
