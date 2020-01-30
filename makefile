.PHONY :build

build:
	go build -v ./cmd/apiserver

.PHONY :test

test:
	go test -race -timeout 30s ./...

.PHONY :test_log

test_log:
	go test -v -race -timeout 30s ./...

.PHONY :run

run: test build
	.\apiserver.exe
	
.DUFAULT_GOAL := build

.PHONY :pack

pack:
	d:\Apps\upx --ultra-brute .\apiserver.exe

.PHONY :deploy_win

deploy_win: test build pack
	copy .\apiserver.exe build\apiserver.exe /Y
	if not exist build\configs mkdir build\configs 
	copy configs\apiserver.toml build\configs\apiserver.toml /-Y
	if not exist build\logs mkdir build\logs
	copy apiserver.db build\apiserver.db /Y