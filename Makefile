PHONY: build
build: 
	go build -o ./build/snake ./cmd/main.go

PHONY: run
run: 
	go run ./cmd