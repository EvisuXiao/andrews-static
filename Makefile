.PHONY: pre-build build clean help

all: build

pre-build:
	make clean
	mkdir -p build/bin
	mkdir build/runtime
	cp -r conf build/

build:
	make pre-build
	@go build -v -o build/bin/upload-bin

clean:
	rm -rf build/*
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make pre-build: clean project and create directories"
	@echo "make clean: remove object files and cached files"