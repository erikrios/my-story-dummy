package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/model"
	"github.com/go-chi/chi/v5"
)

type helloController struct{}

func NewHelloController() *helloController {
	return &helloController{}
}

func (h *helloController) Route(r chi.Router) {
	r.Get("/", h.getHelloHandler)
}

func (h *helloController) getHelloHandler(w http.ResponseWriter, r *http.Request) {
	message := model.Response[string]{
		Status:  "success",
		Message: "Successfully get the data.",
		Data:    "Hello, World!",
	}

	resp, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
