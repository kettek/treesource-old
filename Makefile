GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGENERATE=$(GOCMD) generate

TREESOURCE_SRC=cmd/treesource/treesource.go
ASSETS_DIR=./assets
ASSETS_SRC=assets/assets.go

RELEASE_LDFLAGS="-s -w"
DEBUG_LDFLAGS=""

all: debug

release: $(ASSETS_SRC)
	$(GOBUILD) -ldflags=$(RELEASE_LDFLAGS) -v $(TREESOURCE_SRC)

debug: $(ASSETS_SRC)
	$(GOBUILD) -ldflags=$(DEBUG_LDFLAGS) -v $(TREESOURCE_SRC)

$(ASSETS_SRC):
	$(GOGENERATE) $(ASSETS_DIR)

clean:
	rm -f $(ASSETS_SRC)
	rm -f $(ASSETS_DIR)/resource*.syso
	rm -f treesource treesource.exe
