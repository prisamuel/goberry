### GOBERRY - Economist vanilla microservice template.

#### SETUP

1. Clone this repository into your GOPATH.
2. Run ```godep restore``` to get the minimal dependencies.
3. Drop the .git folder.
4. Design your [RAML](http://raml.org) API interface.
5. Run ```raml-gen``` which will generate HTTP handlers for your service.
6. Copy the generated ```handlers_gen.go``` file to ```handlers.go```.
7. Run ```git init```.
8. Run ```go install``` to build the binary, run it in the app root.
9. The ramlapi package will wire up your endpoints to the handlers.
10. Now build out your service.

### 12-FACTOR GOODNESS

We are aiming to make our microservices [12 factor](http://12factor.net/)

* Create a .env file for environment variables and drop into the project root. The [godotenv](http://github.com/joho/godotenv) package will then parse this file and set environment
variables for everything contained in this file.

### MAKEFILE

The Makefile provides the following:

* `gomkbuild`: build the application binary.
* `gomkinstall`: install the application binary.
* `buildstamp`: build the appication binary, and when the binary is run
  with the --version flag, log build date and build commit hash (default stdout).
  * `installstamp`: install the appication binary, and when the binary is run
  with the --version flag, log build date and build commit hash (default stdout).
* `gomkxbuild`: build all cross-platform binaries, using `gox`.
* `gomkclean`: clean the project directory.
* `vet`: run `go tool vet` on each source file.
* `lint`: run `golint` on each source file.
* `fmt`: run `go fmt` on the entire project.
* `test`: run `go test` for all packages in the project.
* `race`: run `go test` with race detection in all packages in the project.
* `cover`: run tests with coverage report in all pkgs in the project.
* `printvars`: print all variables defined in the Makefile.

### BUILD INFORMATION

To build the service and drop build date information, build as follows:

```go build -ldflags "-X main.buildstamp `date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash `git rev-parse HEAD`"```

Then run <binary> --version=yes

### TESTS

Run ```make test``` for boring old black and white test output.

Run ```pride``` to get nicely colorized test output.

### SERVICE DISCOVERY

The goconsul.json file is present to hook up to a package being built to plug into [consul](https://www.consul.io)