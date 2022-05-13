TEST?=$$(go list ./... | grep -v 'vendor')
NAME=uptycs
BINARY=${NAME}-client-go
VERSION=0.0.4
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

test:
	go test -coverprofile coverage.out -v ./...
	go tool cover -func=coverage.out
