.PHONY : all

all: build test

clean:
	go clean -v .

deps:
	@echo [install dependencies]
	@go get -t -v .

fmt:
	@echo [fmt]
	@gofmt -d -w -e ./

build: deps fmt
	@echo [build]
	@go install ./

test:
	@echo [test]
	@go test -v .

