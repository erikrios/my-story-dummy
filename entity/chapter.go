package entity

type Chapter struct {
	ID        string `json:"id"`
	StoryID   string `json:"storyID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Views     uint   `json:"views"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
