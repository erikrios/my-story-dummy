package payload

type CreateGroup struct {
	Name string `json:"name" validate:"required,alphanum,min=2,max=10"`
}
