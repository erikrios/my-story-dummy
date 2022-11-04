package payload

type CreateStory struct {
	GroupID string `json:"groupID" validate:"required,alphanum,min=2,max=10"`
	Title   string `json:"title" validate:"required,min=2,max=50"`
	Poster  string `json:"poster" validate:"required,min=1"`
}
