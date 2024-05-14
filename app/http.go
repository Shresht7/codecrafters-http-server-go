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

// Set the start-line of the HTTP Request/Response Message
func (r *HTTP) WithStartLine(startLine string) *HTTP {
	r.startLine = startLine
	return r
}

// Set the headers of the HTTP Request/Response Message
func (r *HTTP) WithHeaders(headers map[string]string) *HTTP {
	for key, value := range headers {
		r.headers[key] = value
	}
	return r
}

// Set the body of the HTTP Request/Response Message
func (r *HTTP) WithBody(b string) *HTTP {
	r.body = b
	return r
}

// Parse the message as HTTP. See https://datatracker.ietf.org/doc/html/rfc9112#section-2.2
func (r *HTTP) ParseMessage(message string) *HTTP {
	// Split the message using the separator
	s := strings.Split(message, r.separator)

	// The first line is the start-line
	r.startLine = s[0]

	// Read each header field line into a hash table by field name until the empty line
	for _, line := range s[1:] {
		if line == "" {
			break // Stop when an empty line is encountered
		}
		parts := strings.Split(line, ": ") // Split the line into field-value pairs
		r.headers[parts[0]] = parts[1]
	}

	// The remainder of the message is the body
	r.body = s[len(s)-1]

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
