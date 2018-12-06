package main

import (
	"encoding/json"
	"time"
)

// Payload contains data received from GitHub
type Payload struct {
	Release    PayloadRelease    `json:"release"`
	Repository PayloadRepository `json:"repository"`
	Sender     PayloadSender     `json:"sender"`
}

// PayloadSender contains data received from GitHub about the user
type PayloadSender struct {
	Name string `json:"login"`
}

// PayloadRepository contains data received from GitHub about the repository
type PayloadRepository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
}

// PayloadRelease contains data received from GitHub about the release
type PayloadRelease struct {
	Name       string    `json:"name"`
	Body       string    `json:"body"`
	Date       time.Time `json:"created_at"`
	Draft      bool      `json:"draft"`
	Prerelease bool      `json:"prerelease"`
}

func parsePayload(data []byte) (*Payload, error) {
	var payload Payload

	err := json.Unmarshal(data, &payload)

	if err != nil {
		return nil, err
	}

	return &payload, nil
}
