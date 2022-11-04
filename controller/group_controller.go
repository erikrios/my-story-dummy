package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/util/httperr"
	"github.com/go-chi/chi/v5"
)

type groupController struct{}

func NewGroupController() *groupController {
	return &groupController{}
}

func (g *groupController) Route(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		r.Post("/", g.postCreateGroup)
	})
}

func (g *groupController) postCreateGroup(w http.ResponseWriter, r *http.Request) {
	var p payload.CreateGroup

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		httperr.PayloadParse.Response(w)
	}

	log.Printf("%#v\n", p)

	w.WriteHeader(http.StatusNoContent)
}
