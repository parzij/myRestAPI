package models

// Message represents a user message
// @Description A simple message object with text content
type Message struct {
	// Text content of the message
	// required: true
	// example: "Hello, World!"
	Text string `json:"text"`
}

// Response represents API response
type Response struct {
	// Status of the operation
	// example: Success
	Status string `json:"status"`
	// Human-readable message
	// example: Operation completed successfully
	Message string `json:"message"`
}

// Messages is a global slice for storing messages (for demo purposes)
var Messages []Message