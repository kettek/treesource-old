GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGENERATE=$(GOCMD) generate

all: treesource

treesource: ./assets/assets.go
	$(GOBUILD) -v cmd/treesource/treesource.go

./assets/assets.go:
	$(GOGENERATE) ./assets

clean:
	$(GOCLEAN)
	rm -rf $(BINDIR)
