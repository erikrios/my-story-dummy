package payload

type CreateChapter struct {
	StoryID string `json:"storyID" validate:"required,alphanum,min=2,max=10"`
	Title   string `json:"title" validate:"required,min=2,max=50"`
	Content string `json:"content" validate:"required,min=2"`
}
