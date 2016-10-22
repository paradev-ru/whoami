.PHONY: all

all: build

.PHONY: build build-binaries

build:
	@go build -o whoami *.go

build-binaries:
	@scripts/build-binaries.sh

.PHONY: upload-release upload-pre-release upload-draft-release

upload-release:
	@scripts/upload-binaries.sh --pre-release

upload-pre-release: upload-release

upload-draft-release:
	@scripts/upload-binaries.sh --draft
