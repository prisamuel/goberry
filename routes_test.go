package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func setup(t *testing.T) (*mux.Router, *httptest.ResponseRecorder) {
	router := NewRouter("api.raml")
	res := httptest.NewRecorder()
	return router, res
}
