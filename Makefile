PHONY: run
run: 
	go run ./cmd
	
PHONY: build
build: build-macos

PHONY: build-macos
build-macos: 
	GOOS=darwin GOARCH=arm64 go build -trimpath -a -o ./build/snake ./cmd/main.go

PHONY: build-windows
build-windows: 
	GOOS=windows GOARCH=amd64 go build -trimpath -a -o ./build/snake.exe ./cmd/main.go

