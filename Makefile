main.zip: main
	zip main.zip main

main: modules
	GOOS=linux go build main.go

.PHONY: modules
modules:
	go mod download
