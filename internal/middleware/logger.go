package middleware

import (
	"log"
	"net/http"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &ResponseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		handler(rw, r)
		log.Printf("from mw %v %v %v", rw.statusCode, r.Method, r.URL.Path)

	}
}
