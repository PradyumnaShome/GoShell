#/bin/bash

.PHONY: all fmt

all:
	go build -o build/GoShell
	./build/GoShell

fmt:
	go fmt ./...
