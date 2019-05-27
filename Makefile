GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOINSTALL = $(GOCMD) install
GOFMT = $(GOCMD) fmt

.PHONY: test

check:
	$(GOBUILD) -v ./...

clean:
	$(GOCLEAN)

deps:
	$(GOGET) github.com/stretchr/testify/assert
	$(GOGET) golang.org/x/tools/cmd/cover
	$(GOGET) github.com/gdamore/tcell
	$(GOGET) github.com/urfave/cli

install:
	$(GOINSTALL)

format:
	$(GOFMT) ./...

test:
	$(GOTEST) -v -coverpkg ./chip8 ./test --race
