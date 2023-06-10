LAST_COMMIT := $(shell git rev-parse --short HEAD)
LAST_COMMIT_DATE := $(shell git show -s --format=%ci ${LAST_COMMIT})

BINARY_NAME=tyme

BUILDSTR := (build "\\\#"${LAST_COMMIT} $(shell date '+%Y-%m-%d %H:%M:%S'))
GOPATH ?= $(HOME)/go
TYME ?= $(GOPATH)/bin/TYME

.PHONY: run 
run: 
	CGO_ENABLED=0 go run -ldflags="-s -w -X 'main.buildString=${BUILDSTR}'" ./cmd/tyme

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BINARY_NAME) ./cmd
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)
