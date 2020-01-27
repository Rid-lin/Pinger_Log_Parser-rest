.PHONY :build

build:
	go build -v ./cmd/apiserver

.PHONY :test

test:
	go test -v -race -timeout 30s ./...


.PHONY :run

run: test build
	.\apiserver.exe
	
.DUFAULT_GOAL := build

.PHONY :pack

pack:
	d:\Apps\upx --ultra-brute .\apiserver.exe