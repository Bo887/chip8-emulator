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
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) code.google.com/p/go.tools/cmd/cover

test:
	golangci-lint run
	$(GOTEST) -v -coverpkg ./... ./test
