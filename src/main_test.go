package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Index(response, request)

	t.Run("returns Hello World", func(t *testing.T) {

		got := response.Body.String()
		want := "\n  <h1>Hello, World!</h1>\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
