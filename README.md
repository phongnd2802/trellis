# trellis

Provider-agnostic Go building blocks for AI applications.

```sh
go get github.com/phongnd2802/trellis
```

## Layout

| Path | Purpose |
|---|---|
| `/` | Core types and the `Model` interface. No external dependencies. |
| `provider/<name>` | One package per model provider (Anthropic, OpenAI, ...). |
| `internal/` | Shared helpers that are not public API. |
| `examples/` | Runnable `main` packages. |

## Development

```sh
make test   # go test -race ./...
make lint   # golangci-lint run
```

Integration tests that call real APIs use the `integration` build tag:

```sh
go test -tags integration ./...
```

Status: `v0.x` — the API may change.

## License

[MIT](LICENSE)
