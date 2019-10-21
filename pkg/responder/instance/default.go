package instance

import (
	"encoding/json"
	"net/http"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/responder"
)

type responseError struct {
	Error string `json:"error"`
}

type resp struct {
	Logger logger.Logger
}

func NewJSONResponder(l logger.Logger) responder.Responder {
	return &resp{
		Logger: l,
	}
}

func (r *resp) ResponseGETWithObject(w http.ResponseWriter, object interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&object)
	if err != nil {
		r.Logger.Warn(err)
	}
}

func (r *resp) ResponsePOSTWithObject(w http.ResponseWriter, object interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(&object)
	if err != nil {
		r.Logger.Warn(err)
	}
}

func (r *resp) ResponseOK(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	response := responseError{
		Error: "",
	}

	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		r.Logger.Warn(err)
	}
}

func (r *resp) ResponseWithError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	response := responseError{
		Error: err.Error(),
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		r.Logger.Warn(err)
	}
}
