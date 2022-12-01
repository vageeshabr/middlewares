package main

import (
	"github.com/vageeshabr/middlewares/middlewares"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})

	http.Handle("/api/orders", middlewares.Logger(middlewares.NewCT(AlwaysOkHandler)))

	http.Handle("/api/something", middlewares.Logger(middlewares.CT(AlwaysOkHandler)))

	http.Handle("/api/users", addMiddlewares(AlwaysOkHandler, middlewares.CT, middlewares.Logger))

	log.Println("listening on :8000")

	http.ListenAndServe(":8000", nil)
}

var AlwaysOkHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
})

func addMiddlewares(h http.Handler, middlewares ...func(handler http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}
