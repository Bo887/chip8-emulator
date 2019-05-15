GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

.PHONY: test

check:
	$(GOBUILD) -v ./...

clean:
	$(GOCLEAN)

deps:
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint

test:
	golangci-lint run
	$(GOTEST) -v ./test
