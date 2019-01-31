.PHONY: run build clean install

run: build
	./log_http

build:
	go build .

clean:
	go clean .

install:
	go install .
