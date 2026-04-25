BINARY := harness
PKG    := ./cmd/harness
DIST   := ../dist

.PHONY: all build test tidy run clean

all: build

build:
	go build -o $(DIST)/$(BINARY) $(PKG)

test:
	go test ./...

tidy:
	go mod tidy

run:
	go run $(PKG)

clean:
	rm -f $(DIST)/$(BINARY)
