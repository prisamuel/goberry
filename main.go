package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	port        string
	ramlFile    string
	serviceName string
)

func init() {
	flag.StringVar(&port, "port", ":9494", "port to listen on")
	flag.StringVar(&ramlFile, "ramlFile", "api.raml", "RAML file to parse")
	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()

	NewRouter(ramlFile)

	log.Printf("%s up on port %s\n", serviceName, port)
	log.Fatal(http.ListenAndServe(port, nil))
}
