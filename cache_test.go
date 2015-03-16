package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheHandlerWithNonNegativeCacheTime(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler := newCacheHandler(http.HandlerFunc(http.NotFound), 0)
	handler.ServeHTTP(w, r)
	assert.Equal(t, "max-age=0", w.Header().Get("Cache-Control"))
}

func TestCacheHandlerWithNegativeCacheTime(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler := newCacheHandler(http.HandlerFunc(http.NotFound), -1)
	handler.ServeHTTP(w, r)
	assert.Equal(t, "private,max-age=0,no-cache", w.Header().Get("Cache-Control"))
}
