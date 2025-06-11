package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const filePath = "./server.log" // TODO:: read from config

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func writeLog(pattern, file string) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(pattern); err != nil {
		fmt.Print(err)
	}
	return nil
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
		log.Printf("%v %v %v", rw.statusCode, r.Method, r.URL.Path)
		pattern := fmt.Sprintf("%v %v %v %v\n", time.Now().Format(time.ANSIC), rw.statusCode, r.Method, r.URL.Path)
		writeLog(pattern, filePath)

	}
}
