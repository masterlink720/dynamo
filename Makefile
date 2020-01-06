default: install 

# This how we want to name the binary output
BINARY=dynamo

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

fmt:
	go fmt ./internal/...
	go fmt ./pkg/...
	go fmt ./cmd/...

vet:
	go vet ./internal/...
	go vet ./pkg/...
	go vet ./cmd/...

dev: fmt vet
	go build -race ./cmd/...

test:
	go test -v ./cmd/...

# Builds the project
build: fmt vet
	go build ${LDFLAGS} -o ${BINARY}

# Installs our project: copies binaries
install: fmt vet
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install fmt vet dev build
