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
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) golang.org/x/tools/cmd/cover
	$(GOGET) github.com/gdamore/tcell

test:
	$(GOTEST) -v -coverpkg ./... ./test
