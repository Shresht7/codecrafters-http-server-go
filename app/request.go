package main

import (
	"strings"
)

// Represents a HTTP Request
type Request struct {
	*HTTP         // Embeds the HTTP message
	method string // HTTP Method (e.g. GET, POST, PATCH, DELETE)
	path   string // Path of the requested resource
}

// Parse the incoming request
func ParseRequest(msg string) *Request {
	// Instantiate the Request
	request := &Request{
		HTTP: createHTTPMessage(),
	}

	// Parse the incoming message
	s := strings.Split(msg, request.separator)
	request.parseRequestLine(s[0])
	// TODO: Parse request.headers
	request.body = s[1]

	return request
}

// Parse the Method and Path from the request line
func (r *Request) parseRequestLine(line string) {
	r.startLine = line
	s := strings.Split(r.startLine, " ")
	r.method = s[0]
	r.path = s[1]
}
