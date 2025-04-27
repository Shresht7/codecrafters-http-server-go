package http

import (
	"strconv"
	"strings"
)

// --------------------------------------------------------
// REFERENCE: https://datatracker.ietf.org/doc/html/rfc9112
// --------------------------------------------------------

// CRLF represents a Carriage Return and Line Feed sequence.
// It is used to separate the different parts of the HTTP message
const CRLF = "\r\n"

// Represents a HTTPMessage Request/Response Message.
// See https://datatracker.ietf.org/doc/html/rfc9112#section-2
type HTTPMessage struct {
	protocol string // The protocol version (e.g. `HTTP/1.1`)

	StartLine string   // The first line of the HTTP Request/Response
	Headers   *Headers // The heeders of the HTTP Request/Response
	Body      string   // The body of the HTTP Request/Response

	separator string // The sequence of characters that separate the startLine, headers and the body
}

// Create a HTTP Request/Response
func createHTTPMessage() *HTTPMessage {
	return &HTTPMessage{
		protocol:  "HTTP/1.1",
		Headers:   NewHeaders(),
		separator: CRLF,
	}
}

// Set the start-line of the HTTP Request/Response Message
func (r *HTTPMessage) WithStartLine(startLine string) *HTTPMessage {
	r.StartLine = startLine
	return r
}

// Set the headers of the HTTP Request/Response Message
func (r *HTTPMessage) WithHeaders(headers map[string]string) *HTTPMessage {
	for key, value := range headers {
		r.Headers.Set(key, value)
	}
	return r
}

// Set the body of the HTTP Request/Response Message
func (r *HTTPMessage) WithBody(b string) *HTTPMessage {
	r.Body = b
	return r
}

// Parse the message as HTTP. See https://datatracker.ietf.org/doc/html/rfc9112#section-2.2
func (r *HTTPMessage) ParseMessage(message string) *HTTPMessage {
	// Split the message using the separator
	s := strings.Split(message, r.separator)

	// The first line is the start-line
	r.StartLine = s[0]

	// Read each header field line into a hash table by field name until the empty line
	for _, line := range s[1:] {
		if line == "" {
			break // Stop when an empty line is encountered
		}
		parts := strings.Split(line, ": ") // Split the line into field-value pairs
		r.Headers.Set(parts[0], parts[1])  // Add the field-value pair to the headers
	}

	// Get the Content-Length header
	contentLengthStr, ok := r.Headers.Get("Content-Length")
	if !ok {
		return r
	}
	// Parse the Content-Length header as an integer
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return r
	}

	// The body is the last part of the message and is of length Content-Length
	r.Body = s[len(s)-1][:contentLength]

	return r
}

// The string representation of the HTTP Request/Response
func (r *HTTPMessage) String() string {
	// If the body is empty, return the startLine and headers only
	if r.Body == "" || r.Body == CRLF {
		return strings.Join([]string{
			r.StartLine,
			r.Headers.String(),
		}, r.separator)
	}
	return strings.Join([]string{
		r.StartLine,
		r.Headers.String(),
		r.Body,
	}, r.separator)
}

// The byte-array representation of the HTTP Request/Response
func (r *HTTPMessage) Bytes() []byte {
	return []byte(r.String())
}
