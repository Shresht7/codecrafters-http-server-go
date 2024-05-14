package main

import (
	"fmt"
	"strings"
)

// --------------------------------------------------------
// REFERENCE: https://datatracker.ietf.org/doc/html/rfc9112
// --------------------------------------------------------

// Represents a HTTP Request/Response Message.
// See https://datatracker.ietf.org/doc/html/rfc9112#section-2
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
		headers:   make(map[string]string),
		separator: CRLF,
	}
}

// Set the start-line of the HTTP Request/Response
func (r *HTTP) WithStartLine(message string) *HTTP {
	r.startLine = message
	return r
}

// Set the headers of the HTTP Request/Response Message
func (r *HTTP) WithHeaders(headers map[string]string) *HTTP {
	for key, value := range headers {
		r.headers[key] = value
	}
	return r
}

// Set the body of the HTTP Request/Response
func (r *HTTP) WithBody(b string) *HTTP {
	r.body = b
	return r
}

// Generate a string representation of the Headers
func (r *HTTP) headersString() string {
	fieldLines := make([]string, 0, len(r.headers))
	for key, value := range r.headers {
		fieldLines = append(fieldLines, fmt.Sprintf("%s: %s", key, value))
	}
	return strings.Join(fieldLines, r.separator)
}

// The string representation of the HTTP Request/Response
func (r *HTTP) String() string {
	return strings.Join([]string{
		r.startLine,
		r.headersString(),
		r.separator, // Add an extra CRLF to separate the headers from the body
		r.body,
	}, r.separator)
}

// The byte-array representation of the HTTP Request/Response
func (r *HTTP) Bytes() []byte {
	return []byte(r.String())
}
