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
	*HTTP          // Embeds the HTTP message
	statusCode int // The status code for the response
}

// Create a new HTTP Response
func createResponse() *Response {
	return &Response{
		HTTP:       createHTTPMessage(),
		statusCode: http.StatusInternalServerError,
	}
}

// Set the status code of the HTTP Response
func (r *Response) WithStatus(code int) *Response {
	r.statusCode = code
	text := http.StatusText(code)
	codeStr := strconv.Itoa(code)
	statusMsg := strings.Join([]string{r.protocol, codeStr, text}, " ")
	r.WithStartLine(statusMsg)
	return r
}
