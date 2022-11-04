package payload

type CreateChapter struct {
	GroupID string `json:"groupID" validate:"required,alphanum,min=2,max=10"`
	StoryID string `json:"storyID" validate:"required,ascii,min=2,max=10"`
	Title   string `json:"title" validate:"required,min=2,max=50"`
	Content string `json:"content" validate:"required,min=2"`
}
