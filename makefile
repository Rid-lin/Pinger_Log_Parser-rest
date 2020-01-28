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

.PHONY :deploy_win

deploy_win:
	copy .\apiserver.exe bin\apiserver.exe /Y
	if not exist bin\configs mkdir bin\configs 
	copy configs\apiserver.toml bin\configs\apiserver.toml /-Y
	if not exist bin\logs mkdir bin\logs
	copy apiserver.db bin\apiserver.db /Y