default: build test-integration lint

build:
	go build ./...

test:
	go test -short ./...

test-integration:
	go test -count=1 ./...

lint: fmt vet

fmt:
	go fmt ./...

vet:
	go vet ./...
