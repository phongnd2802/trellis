package trellis

import (
	"context"
	"iter"
)

// Model is implemented by every provider.
type Model interface {
	Generate(ctx context.Context, req *Request) (*Response, error)
	Stream(ctx context.Context, req *Request) iter.Seq2[*Chunk, error]
}

// Chunk is one streamed piece of a response.
type Chunk struct {
	Delta    string
	ToolCall *ToolCall // set when a complete tool call has been streamed
	Done     *Response // set on the final chunk
}
