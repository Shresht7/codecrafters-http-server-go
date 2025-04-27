package http

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

// Represents a HTTP Request
type Request struct {
	*HTTPMessage        // Embeds the HTTP message
	Method       string // HTTP Method (e.g. GET, POST, PATCH, DELETE)
	Path         string // Path of the requested resource
}

// Parse the incoming request
func ParseRequest(conn net.Conn) *Request {
	// Instantiate the Request
	request := &Request{
		HTTPMessage: createHTTPMessage(),
	}

	// Create a buffered reader to read the connection
	reader := bufio.NewReader(conn)

	// Read and parse the request line
	startLine, err := reader.ReadString('\n')
	if err == io.EOF {
		return nil // End of connection stream. Connection closed
	}
	if err != nil {
		fmt.Println("Error reading request line: ", err.Error())
		return nil
	}
	request.StartLine = startLine
	request.parseRequestLine()

	// Read each header field into a hash table by field name until we hit an empty line
	for {
		line, err := reader.ReadString('\n')
		if line == "" {
			break // Empty line is the delimiter between header and body
		}
		if err != nil {
			fmt.Println("Error reading header line: ", err.Error())
			return nil
		}
		parts := strings.Split(line, ": ")      // Split the line into field-value pairs
		request.Headers.Set(parts[0], parts[1]) // Add the field-value pair to the headers
	}

	// Get the Content-Length header
	contentLengthStr, ok := request.Headers.Get("Content-Length")
	if !ok {
		return request
	}
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return request
	}

	// The body is the last part of the message and is of Content-Length
	buf := make([]byte, contentLength)
	reader.Read(buf)
	request.Body = string(buf)

	return request
}

// Parse the Method and Path from the request line. See https://datatracker.ietf.org/doc/html/rfc9112#section-3
func (r *Request) parseRequestLine() {
	s := strings.Fields(r.StartLine)
	r.Method = s[0]
	r.Path = s[1]
	r.protocol = s[2]
}
