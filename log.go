package main

import (
	"log"
	"net/http"
)

type logHandler struct {
	handler http.Handler
}

func newLogHandler(h http.Handler) http.Handler {
	return &logHandler{h}
}

func (h *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
	log.Printf(`- %s - %s "%s %s" "%s"`,
		r.RemoteAddr, r.Proto, r.Method, r.RequestURI, r.Header.Get("User-Agent"))
}
