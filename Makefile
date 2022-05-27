default: test

test:
	go test -coverprofile coverage.out -v ./uptycs/...
	go tool cover -func=coverage.out
