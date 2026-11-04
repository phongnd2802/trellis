// Package trellis provides provider-agnostic building blocks for AI applications.
//
// The root package holds only core types and interfaces and has no external
// dependencies. Concrete model providers live under provider/.
package trellis

import "encoding/json"

// Role identifies who authored a message.
type Role string

// Roles a Message can have.
const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

// Message is one turn in a conversation.
type Message struct {
	Role    Role
	Content string

	// ToolCalls is set on assistant messages that request tool execution.
	ToolCalls []ToolCall
	// ToolCallID links a RoleTool message to the call it answers.
	ToolCallID string
}

// Request is the input to a Model.
type Request struct {
	Model       string
	Messages    []Message
	Tools       []Tool
	MaxTokens   int
	Temperature *float64 // nil means provider default
}

// Response is a complete model reply.
type Response struct {
	Message    Message
	StopReason StopReason
	Usage      Usage
}

// StopReason explains why generation ended.
type StopReason string

// Stop reasons a provider can report.
const (
	StopEnd       StopReason = "end"
	StopMaxTokens StopReason = "max_tokens"
	StopToolUse   StopReason = "tool_use"
)

// Usage reports token consumption.
type Usage struct {
	InputTokens  int
	OutputTokens int
}

// ToolCall is a model's request to run a tool.
type ToolCall struct {
	ID    string
	Name  string
	Input json.RawMessage
}
