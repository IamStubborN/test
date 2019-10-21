package instance

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/IamStubborN/test/pkg/logger"
	"github.com/IamStubborN/test/pkg/mware"
	"github.com/IamStubborN/test/pkg/responder"
)

type creds struct {
	Token string `json:"token"`
}

type middleWare struct {
	logger    logger.Logger
	responder responder.Responder
}

func NewMiddleWare(l logger.Logger, r responder.Responder) mware.MWare {
	return middleWare{
		logger:    l,
		responder: r,
	}
}

func (mw middleWare) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			mw.responder.ResponseWithError(w, ErrEmptyBody, http.StatusBadRequest)
		}

		body := ioutil.NopCloser(bytes.NewBuffer(buf))
		copyBody := ioutil.NopCloser(bytes.NewBuffer(buf))
		r.Body = copyBody

		var auth creds
		err = json.NewDecoder(body).Decode(&auth)
		if err != nil {
			mw.responder.ResponseWithError(w, ErrInvalidBody, http.StatusBadRequest)
		}

		if auth.Token != "testtask" {
			mw.responder.ResponseWithError(w, ErrInvalidToken, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (mw middleWare) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			mw.logger.WithFields([]interface{}{
				"path", r.URL.Path,
				"method", r.Method,
				"status", ww.Status(),
				"size", ww.BytesWritten(),
			}).Info("served")
		}()
		next.ServeHTTP(ww, r)
	})
}
