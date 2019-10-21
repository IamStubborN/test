package mware

import "net/http"

type MWare interface {
	AuthMiddleware(next http.Handler) http.Handler
	RequestLogger(next http.Handler) http.Handler
}
