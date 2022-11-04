package model

type ListedData interface {
	string
}

type Response[T ListedData] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
