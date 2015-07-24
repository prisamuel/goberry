package main

import (
	"encoding/json"
	"net/http"
)

var RouteMap = map[string]http.HandlerFunc{

	"Root": Root,
}

// Handler for rest URI / and the action GET
//
func Root(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(map[string]string{
		"message": "RootGET",
	})
	w.Write(json)
}
