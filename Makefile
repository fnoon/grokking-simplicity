all: shopping

shopping: shopping.go
	go fmt ./...
	go build

.PHONY: run
run: shopping
	go run shopping

.PHONY: clean
clean:
	go clean
