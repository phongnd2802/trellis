package trellis

import "errors"

// Sentinel errors providers wrap so callers can use errors.Is.
var (
	ErrRateLimited   = errors.New("trellis: rate limited")
	ErrContextLength = errors.New("trellis: context length exceeded")
	ErrUnauthorized  = errors.New("trellis: unauthorized")
)
