# Compile binaries
build-http-mux:
	go build -o ./build/http-mux ./cmd/http_mux/main.go
clean:
	rm -rf ./build

# Development
test:
	go test ./... -cover
tidy:
	go mod tidy
format:
	go fmt ./...
