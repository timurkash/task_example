package models

import "time"

type AddModel struct {
	Guid string `json:"guid"`
}

type TaskModelSQL struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type TaskModel struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}
