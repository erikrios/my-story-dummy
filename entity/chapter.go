package entity

import "time"

type Chapter struct {
	ID        string    `json:"id"`
	StoryID   string    `json:"story_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Views     uint      `json:"views"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
