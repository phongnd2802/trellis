.PHONY: test lint vuln cover check

GOLANGCI_LINT_VERSION := v2.15.0 # keep in sync with .github/workflows/ci.yml

test:
	go test -race -shuffle=on ./...

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run

vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

cover:
	go test -coverprofile=cover.out ./... && go tool cover -func=cover.out

# Everything CI runs.
check:
	go mod tidy -diff
	go vet ./...
	$(MAKE) test lint vuln
