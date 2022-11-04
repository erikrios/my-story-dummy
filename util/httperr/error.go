package httperr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/model"
)

type Error int64

const (
	NotFound Error = iota
	MethodNotAllowed
	PayloadParse
)

func (e Error) Response(w http.ResponseWriter) {
	var message string
	var data string
	var code int

	switch e {
	case NotFound:
		message = "failed to access"
		data = "Route does not exist"
		code = http.StatusNotFound
	case MethodNotAllowed:
		message = "failed to access"
		data = "Method is not allowed in this route"
		code = http.StatusMethodNotAllowed
	case PayloadParse:
		message = "failed to parse payload"
		data = "Invalid request body, please check the API documentation"
		code = http.StatusBadRequest
	default:
		message = "server error happened"
		data = "Something went wrong."
		code = http.StatusInternalServerError
	}

	w.Header().Add("Content-Type", "application/json")
	response := model.Response[string]{
		Status:  "error",
		Message: message,
		Data:    data,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(code)
	w.Write(resp)
}
