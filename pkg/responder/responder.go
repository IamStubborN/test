package responder

import "net/http"

type Responder interface {
	ResponseWithObject(w http.ResponseWriter, object interface{}, code int)
	ResponseOK(w http.ResponseWriter, code int)
	ResponseWithError(w http.ResponseWriter, err error, code int)
}
