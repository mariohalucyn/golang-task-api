.PHONY: target build run_client run_server

target: run_client run_server

build:
	@go build .
	@cd client && npm run build

run_server: build
	@./react-golang

run_client: build
	@cd client && npm run preview
