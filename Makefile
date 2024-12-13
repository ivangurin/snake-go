PHONY: build
build: 
	go build -o ./build ./cmd/main.go

PHONY: run
run: 
	go run ./cmd