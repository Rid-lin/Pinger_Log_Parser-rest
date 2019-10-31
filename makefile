.PHONY build
build:
	go build -v ./cmd/apiserver

.DUFAULT_GOAL := build