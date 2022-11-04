package controller

import (
	"net/http"

	"github.com/erikrios/my-story-dummy/util/httperr"
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
	httperr.NotFound.Response(w)
}

func (c customController) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	httperr.MethodNotAllowed.Response(w)
}
