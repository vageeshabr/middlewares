package middlewares

import (
	"log"
	"net/http"
)

type ContentTypeValidator struct {
	inner http.Handler
}

func NewCT(inner http.Handler) *ContentTypeValidator {
	return &ContentTypeValidator{inner: inner}
}

func (middleware *ContentTypeValidator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(400)
		w.Write([]byte(`invalid content type`))
	}
	middleware.inner.ServeHTTP(w, r)
}

func CT(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Content type middleware called.")
		w.Header().Set("Content-Type", "application/json")
		inner.ServeHTTP(w, r)
	})
}
