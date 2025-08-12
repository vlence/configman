package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vlence/gossert"
)

type customResponseWriter struct {
        http.ResponseWriter
        statusCode int
}

func (w *customResponseWriter) WriteHeader(statusCode int) {
        w.statusCode = statusCode
        w.ResponseWriter.WriteHeader(statusCode)
}

func (w *customResponseWriter) Flush() {
        flusher, ok := w.ResponseWriter.(http.Flusher)

        if ok {
                flusher.Flush()
        }
}

// logger logs the time it took for the request, the request method and the request url path.
func logger(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                start := time.Now()
                ww := &customResponseWriter{w, -1}
                next.ServeHTTP(ww, r)
                end := time.Now()
                gossert.Ok(ww.statusCode != -1, "logger: response status code is -1")
                log.Printf("%s %d %s %s\n", end.Sub(start).String(), ww.statusCode, r.Method, r.URL.Path)
        })
}
