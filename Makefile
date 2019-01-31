.PHONY: run build clean

run: build
	@./log_http

build:
	@go build .

clean:
	@go clean .
