package main

import (
	"strings"
)

// Represents a HTTP Request
type Request struct {
	*HTTP // Embeds the HTTP message
}

func ParseRequest(msg string) *Request {
	// Instantiate the Request
	request := &Request{
		HTTP: createHTTPMessage(),
	}

	// Parse the incoming message
	s := strings.Split(msg, request.separator)
	request.startLine = s[0]
	// TODO: Parse request.headers
	request.body = s[1]

	return request
}
