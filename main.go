package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/buddhamagnet/goconsul"
	"github.com/joho/godotenv"

	kitlog "github.com/go-kit/kit/log"
	stdlog "log"
)

var (
	port                string
	ramlFile            string
	serviceName         string
	serviceRegistration string
	logger              kitlog.Logger
)

func init() {
	flag.StringVar(&port, "port", ":9494", "port to listen on")
	flag.StringVar(&ramlFile, "ramlFile", "api.raml", "RAML file to parse")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	serviceName = os.Getenv("SERVICE_NAME")
	serviceRegistration = os.Getenv("SERVICE_REGISTRATION")

	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
	}

	// Integrate go-kit logger.
	logger = kitlog.NewJSONLogger(os.Stdout)
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger))
}

func main() {

	if serviceRegistration != "" {
		if err := goconsul.RegisterService(); err != nil {
			log.Fatal(err)
		}
	}

	flag.Parse()

	NewRouter(ramlFile)

	logChannel("information", fmt.Sprintf("%s up on port %s", serviceName, port))
	log.Fatal(http.ListenAndServe(port, nil))
}
