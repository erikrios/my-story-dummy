package customerr

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erikrios/my-story-dummy/model"
)

type Error int64

func (e Error) String() string {
	var message string
	switch e {
	case NotFound:
		message = "NotFound"
	case MethodNotAllowed:
		message = "MethodNotAllowed"
	case PayloadParse:
		message = "PayloadParse"
	case InvalidPayload:
		message = "InvalidPayload"
	case ResourceAlreadyExists:
		message = "ResourceAlreadyExists"
	case Internal:
		message = "Internal"
	default:
		message = "UnexpectedError"
	}
	return message
}

func (e Error) Error() string {
	return e.String()
}

const (
	NotFound Error = iota
	MethodNotAllowed
	PayloadParse
	InvalidPayload
	ResourceAlreadyExists
	Internal
)

func (e Error) HTTPResponse(w http.ResponseWriter) {
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
	case InvalidPayload:
		message = "invalid payload"
		data = "Invalid request body, please check the API documentation"
		code = http.StatusBadRequest
	case ResourceAlreadyExists:
		message = "failed to create"
		data = "Resource already exists"
		code = http.StatusConflict
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
