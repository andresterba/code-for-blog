GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOFORMAT=$(GOCMD) fmt
BINARY_NAME=gocker

all: build
run: build
	./$(BINARY_NAME)
build: 
	$(GOBUILD) -o $(BINARY_NAME)
format:
	$(GOFORMAT) ./...
test:
	$(GOTEST) ./... -cover