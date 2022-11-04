package controller

import (
	"encoding/json"
	"net/http"

	"github.com/erikrios/my-story-dummy/model"
	"github.com/erikrios/my-story-dummy/model/payload"
	"github.com/erikrios/my-story-dummy/service"
	"github.com/erikrios/my-story-dummy/util/customerr"
	"github.com/go-chi/chi/v5"
)

type storyController struct {
	service service.StoryService
}

func NewStoryController(service service.StoryService) *storyController {
	return &storyController{service: service}
}

func (s *storyController) Route(r chi.Router) {
	r.Route("/stories", func(r chi.Router) {
		r.Post("/", s.postCreateStory)
	})
}

func (s *storyController) postCreateStory(w http.ResponseWriter, r *http.Request) {
	var p payload.CreateStory

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		customerr.PayloadParse.HTTPResponse(w)
		return
	}

	id, err := s.service.Create(p)
	if err != nil {
		if customErr, ok := err.(customerr.Error); ok {
			customErr.HTTPResponse(w)
			return
		}
		customerr.Internal.HTTPResponse(w)
		return
	}

	idResp := model.Response[string]{
		Status:  "success",
		Message: "successfully create a story",
		Data:    id,
	}

	idRespBytes, marshalErr := json.Marshal(idResp)
	if marshalErr != nil {
		customerr.Internal.HTTPResponse(w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(idRespBytes)
}
