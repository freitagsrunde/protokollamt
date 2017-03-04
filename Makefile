.PHONY: all clean deps build
all: clean deps build

clean:
	go clean -i ./...
	rm -f protokollamt

deps:
	go get -v -t ./...

build:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'
