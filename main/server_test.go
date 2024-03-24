package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	server := createPlayerServer()

	t.Run("returns Pepper's score", func(t *testing.T) {
		request, response := createRequestResponse("/players/Pepper")

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Steve's score", func(t *testing.T) {
		request, response := createRequestResponse("/players/Steve")

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func createRequestResponse(playerPath string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(http.MethodGet, playerPath, nil)
	response := httptest.NewRecorder()
	return request, response
}

func createPlayerServer() *PlayerServer {
	return &PlayerServer{}
}

func assertResponseBody(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
