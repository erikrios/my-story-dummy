package controller

import (
	"encoding/json"
	"net/http"

	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/service"
	"github.com/erikrios/my-story-dummy/util/customerr"
	"github.com/go-chi/chi/v5"
)

type groupController struct {
	service service.GroupService
}

func NewGroupController(service service.GroupService) *groupController {
	return &groupController{service: service}
}

func (g *groupController) Route(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		r.Post("/", g.postCreateGroup)
	})
}

func (g *groupController) postCreateGroup(w http.ResponseWriter, r *http.Request) {
	var p payload.CreateGroup

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		customerr.PayloadParse.HTTPResponse(w)
		return
	}

	if err := g.service.Create(p); err != nil {
		if customErr, ok := err.(customerr.Error); ok {
			customErr.HTTPResponse(w)
			return
		}
		customerr.Internal.HTTPResponse(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
