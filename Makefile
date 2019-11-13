GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

SERVER_BINARY=bourbon-finder

all: build

build:
	cd cmd/bourbon-finder && $(GOBUILD)

run:
	./cmd/$(SERVER_BINARY)/$(SERVER_BINARY)

clean:
	rm -f cmd/$(SERVER_BINARY)/$(SERVER_BINARY)

fmt:
	cd cmd/$(SERVER_BINARY) && $(GOFMT)
	cd src && $(GOFMT) bourbonfinder

vet:
	cd cmd/$(SERVER_BINARY) && $(GOVET)
	cd src && $(GOVET) bourbonfinder

deps:
	echo no deps





