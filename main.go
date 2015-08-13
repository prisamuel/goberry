package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/buddhamagnet/goconsul"
	"github.com/joho/godotenv"

	stdlog "log"

	kitlog "github.com/go-kit/kit/log"
)

var (
	port                string
	buildstamp          string
	githash             string
	version             string
	ramlFile            string
	serviceName         string
	serviceRegistration string
	logger              kitlog.Logger
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

	if version != "" {
		logChannel("build", fmt.Sprintf("build date: %s commit: %s", buildstamp, githash))
	}

	logChannel("information", fmt.Sprintf("%s up on port %s", serviceName, port))
	log.Fatal(http.ListenAndServe(port, nil))
}
