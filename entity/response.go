package entity

type Response struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []Story `json:"data"`
}
