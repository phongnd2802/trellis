package trellis_test

import (
	"context"
	"fmt"
	"iter"

	"github.com/phongnd2802/trellis"
)

// echoModel is a minimal Model used to check the interface is implementable.
type echoModel struct{}

func (echoModel) Generate(_ context.Context, req *trellis.Request) (*trellis.Response, error) {
	last := req.Messages[len(req.Messages)-1]
	return &trellis.Response{
		Message:    trellis.Message{Role: trellis.RoleAssistant, Content: last.Content},
		StopReason: trellis.StopEnd,
	}, nil
}

func (m echoModel) Stream(ctx context.Context, req *trellis.Request) iter.Seq2[*trellis.Chunk, error] {
	return func(yield func(*trellis.Chunk, error) bool) {
		resp, err := m.Generate(ctx, req)
		if err != nil {
			yield(nil, err)
			return
		}
		if !yield(&trellis.Chunk{Delta: resp.Message.Content}, nil) {
			return
		}
		yield(&trellis.Chunk{Done: resp}, nil)
	}
}

var _ trellis.Model = echoModel{}

func ExampleModel() {
	var m trellis.Model = echoModel{}
	req := &trellis.Request{Messages: []trellis.Message{{Role: trellis.RoleUser, Content: "hello"}}}

	for chunk, err := range m.Stream(context.Background(), req) {
		if err != nil {
			panic(err)
		}
		fmt.Print(chunk.Delta)
	}
	// Output: hello
}
