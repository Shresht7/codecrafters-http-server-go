package main

import (
	"net/http"
	"strconv"
	"strings"
)

// HTTP Response is made up of three parts, each separated by a [CRLF](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) (`\r\n`):
// 1. Status line: `HTTP/1.1 200 OK`
// 2. One or more Headers: `Content-Type: text/html`
// 3. (Optional) Body: `<!DOCTYPE html><html><body><h1>Hello, World!</h1></body></html>`

// Represents an HTTP response
type Response struct {
	protocol   string            // The protocol version (e.g. HTTP/1.1)
	statusCode int               // The status code for the response
	headers    map[string]string // The headers of the response
	body       string            // The body of the response
	separator  string            // The CRLF separator
}

// Create a new HTTP Response
func createResponse() *Response {
	return &Response{
		protocol:  "HTTP/1.1",
		separator: CRLF,
	}
}

// Set the status code of the HTTP Response
func (r *Response) WithStatus(code int) *Response {
	r.statusCode = code
	return r
}

// Set the headers of the HTTP Response
func (r *Response) WithHeader(headers map[string]string) *Response {
	r.headers = headers
	return r
}

// Set the body of the HTTP Response
func (r *Response) WithBody(b string) *Response {
	r.body = b
	return r
}

// Generate the string representation for the status line
func (r *Response) statusLineStr() string {
	code := strconv.Itoa(r.statusCode)
	text := http.StatusText(r.statusCode)
	return strings.Join([]string{r.protocol, code, text}, " ")
}

// Generate the string representation for the headers
func (r *Response) headersStr() string {
	// TODO: Implementation
	return ""
}

// Generate the string representation of the body
func (r *Response) bodyStr() string {
	// TODO: Implementation
	return ""
}

// The string representation of the HTTP Response
func (r *Response) String() string {
	statusLine := r.statusLineStr()
	headers := r.headersStr()
	body := r.bodyStr()
	return (strings.Join([]string{statusLine, headers, body}, r.separator))
}

// The byte-array representation of the HTTP Response
func (r *Response) Bytes() []byte {
	return []byte(r.String())
}
