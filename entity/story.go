package entity

import "time"

type Story struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Poster    string    `json:"poster"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Chapters  []Chapter `json:"chapters"`
}
