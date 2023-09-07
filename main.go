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
	Utc_time        string `json:"utc_time"`
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
			Utc_time:        time.Now().Format("2006-01-02T15:04:05Z"),
			Track:           track,
			Github_file_url: "https://github.com/jenweani/task1/blob/main/main.go",
			Github_repo_url: "https://github.com/jenweani/task1",
			Status_code:     http.StatusOK,
		}

		json.NewEncoder(w).Encode(task1)
	})

	http.ListenAndServe(":8080", nil)
}
