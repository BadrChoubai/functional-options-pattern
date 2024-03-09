all: build

build:
	go build

clean:
	rm -rf ./bin

lint: govulncheck
	golangci-lint run ./...

govulncheck:
	govulncheck

.DEFAULT_GOAL := all
.PHONY: govulncheck bin/vulncheck
