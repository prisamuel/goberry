package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/EconomistDigitalSolutions/watchman/journal"
	_ "github.com/EconomistDigitalSolutions/watchman/meter"

	"github.com/buddhamagnet/goconsul"
	"github.com/joho/godotenv"
)

var (
	port                string
	buildstamp          string
	githash             string
	version             string
	ramlFile            string
	serviceName         string
	serviceRegistration string
)

func init() {
	flag.StringVar(&port, "port", ":9494", "port to listen on")
	flag.StringVar(&version, "version", "", "output build date and commit data")
	flag.StringVar(&ramlFile, "ramlFile", "api.raml", "RAML file to parse")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	serviceName = os.Getenv("SERVICE_NAME")
	serviceRegistration = os.Getenv("SERVICE_REGISTRATION")
	journal.Service = serviceName

	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
	}
}

func main() {

	if serviceRegistration != "" {
		if err := goconsul.RegisterService(); err != nil {
			log.Fatal(err)
		}
	}

	flag.Parse()

	NewRouter(ramlFile)

	if version != "" {
		journal.LogChannel("build", fmt.Sprintf("build date: %s commit: %s", buildstamp, githash))
	}

	journal.LogChannel("information", fmt.Sprintf("%s up on port %s", serviceName, port))
	log.Fatal(http.ListenAndServe(port, nil))
}
