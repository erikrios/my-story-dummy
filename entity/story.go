package entity

type Story struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Poster    string    `json:"poster"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	Chapters  []Chapter `json:"chapters"`
}
