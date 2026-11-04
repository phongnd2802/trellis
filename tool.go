package trellis

import (
	"context"
	"encoding/json"
)

// Tool is a function the model may call.
type Tool struct {
	Name        string
	Description string
	InputSchema json.RawMessage // JSON Schema object
	Handler     func(ctx context.Context, input json.RawMessage) (string, error)
}
