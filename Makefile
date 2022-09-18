# Compile binaries
build-http-mux:
	go build -o ./build/http-mux ./cmd/http_mux/main.go
build-http-mux-pg:
	go build -o ./build/http-mux-pg ./cmd/http_mux_pg/main.go
clean:
	rm -rf ./build

# Development
test:
	go test ./... -cover
tidy:
	go mod tidy
format:
	go fmt ./...
