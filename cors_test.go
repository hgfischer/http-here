package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSHandler(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler := newCORSHandler(http.HandlerFunc(http.NotFound))
	handler.ServeHTTP(w, r)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
}
