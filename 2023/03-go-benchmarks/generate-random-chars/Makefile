GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test
test:
	$(GOTEST) ./...
bench:
	$(GOTEST) -bench=. -benchmem -benchtime=100x ./...
profile:
	$(GOTEST) -cpuprofile cpu.prof -memprofile mem.prof -bench .
	# go tool pprof -http 127.0.0.1:8080 cpu.prof
testcoverage:
	$(GOTEST) -coverprofile coverage.out ./... && go tool cover -html=coverage.out && rm coverage.out
bench-diff:
	$(GOTEST) -bench=. -benchmem -benchtime=100x -count=10 ./... >> old.txt
	$(GOTEST) -bench=. -benchmem -benchtime=100x -count=10 ./... >> new.txt
	benchstat old.txt new.txt
	rm old.txt && rm new.txt