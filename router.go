package main

import (
	"log"
	"net/http"

	"github.com/EconomistDigitalSolutions/ramlapi"
	"github.com/gorilla/mux"
)

var router *mux.Router

// NewRouter creates a mux router, sets up
// a static handler and registers the dynamic
// routes and middleware handlers with the mux.
func NewRouter(ramlFile string) *mux.Router {
	router = mux.NewRouter().StrictSlash(true)
	// Assemble middleware as required.
	// assembleMiddleware(router)
	assembleRoutes(router, ramlFile)
	return router
}

// assembleMiddleware sets up the middleware stack for gref.
func assembleMiddleware(r *mux.Router) {
	http.Handle("/",
		JSONMiddleware(
			LoggingMiddleware(
				RecoverMiddleware(r))))
}

func assembleRoutes(r *mux.Router, f string) {
	// Parse the RAML API specification.
	api, err := ramlapi.ProcessRAML(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing API spec for", api.Title)
	log.Println("Base URI at", api.BaseUri)
	ramlapi.Build(api, routerFunc)
}

func routerFunc(data map[string]string) {
	router.
		Methods(data["verb"]).
		Path(data["path"]).
		Handler(RouteMap[data["handler"]])
}
