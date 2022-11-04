package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/model"
	"github.com/go-chi/chi/v5"
)

type customController struct{}

func NewCustomController() *customController {
	return &customController{}
}

func (c *customController) Route(r chi.Router) {
	r.NotFound(c.notFoundHandler)
	r.MethodNotAllowed(c.methodNotAllowed)
}

func (c customController) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := model.Response[string]{
		Status:  "error",
		Message: "failed to access.",
		Data:    "Route does not exist",
	}

	resp, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(resp)
}

func (c customController) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := model.Response[string]{
		Status:  "error",
		Message: "failed to access.",
		Data:    "Method is not valid",
	}

	resp, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(resp)
}
