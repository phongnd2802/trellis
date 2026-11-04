.PHONY: test lint cover

test:
	go test -race ./...

lint:
	golangci-lint run

cover:
	go test -coverprofile=cover.out ./... && go tool cover -func=cover.out
