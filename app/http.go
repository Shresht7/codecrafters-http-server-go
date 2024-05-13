package main

import (
	"fmt"
	"strings"
)

// Represents a HTTP Request/Response
type HTTP struct {
	protocol string // The protocol version (e.g. `HTTP/1.1`)

	startLine string            // The first line of the HTTP Request/Response
	headers   map[string]string // The heeders of the HTTP Request/Response
	body      string            // The body of the HTTP Request/Response

	separator string // The sequence of characters that separate the startLine, headers and the body
}

// Create a HTTP Request/Response
func createHTTPMessage() *HTTP {
	return &HTTP{
		protocol:  "HTTP/1.1",
		separator: CRLF,
	}
}

// Set the start-line of the HTTP Request/Response
func (r *HTTP) WithStartLine(message string) *HTTP {
	r.startLine = message
	return r
}

// Set the headers of the HTTP Request/Response
func (r *HTTP) WithHeader(headers map[string]string) *HTTP {
	r.headers = headers
	return r
}

// Set the body of the HTTP Request/Response
func (r *HTTP) WithBody(b string) *HTTP {
	r.body = b
	return r
}

// Generate a string representation of the Headers
func (r *HTTP) headersString() string {
	sb := strings.Builder{}
	for k, v := range r.headers {
		sb.WriteString(fmt.Sprintf("%s: %s", k, v))
	}
	return sb.String()
}

// The string representation of the HTTP Request/Response
func (r *HTTP) String() string {
	return strings.Join([]string{
		r.startLine,
		r.headersString(),
		r.body,
	}, r.separator)
}

// The byte-array representation of the HTTP Request/Response
func (r *HTTP) Bytes() []byte {
	return []byte(r.String())
}