GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGENERATE=$(GOCMD) generate
BINDIR=bin/
BINARY_NAME=treesource
BINARY=$(BINDIR)$(BINARY_NAME)

all: treesource

treesource: ./assets/assets.go
	$(GOBUILD) -o $(BINARY) -v cmd/treesource/main.go

./assets/assets.go:
	$(GOGENERATE) ./assets

clean:
	$(GOCLEAN)
	rm -rf $(BINDIR)
