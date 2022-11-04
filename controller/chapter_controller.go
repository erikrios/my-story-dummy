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

type chapterController struct {
	service service.ChapterService
}

func NewChapterController(service service.ChapterService) *chapterController {
	return &chapterController{service: service}
}

func (c *chapterController) Route(r chi.Router) {
	r.Route("/chapters", func(r chi.Router) {
		r.Post("/", c.postCreateChapter)
	})
}

func (c *chapterController) postCreateChapter(w http.ResponseWriter, r *http.Request) {
	var p payload.CreateChapter

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		customerr.PayloadParse.HTTPResponse(w)
		return
	}

	id, err := c.service.Create(p)
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
		Message: "successfully create a chapter",
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
