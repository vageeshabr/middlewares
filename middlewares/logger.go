package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *StatusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func Logger(inner http.Handler) http.Handler {
	logger := log.New(os.Stdout, "[REQ] ", 0)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logger middleware called.")

		start := time.Now()
		srw := &StatusResponseWriter{ResponseWriter: w}

		defer func(res *StatusResponseWriter, req *http.Request) {
			line, _ := json.Marshal(map[string]interface{}{
				"ResponseTime": time.Since(start).Nanoseconds() / 1000,
				"Method":       req.Method,
				"UserAgent":    req.UserAgent(),
				"URI":          req.RequestURI,
				"Code":         srw.status,
			})
			logger.Printf("%s", string(line))
		}(srw, r)

		inner.ServeHTTP(srw, r)

	})
}
