package main

import (
	"fmt"
	"net"
	"strings"
)

// Represents a HTTP Request
type Request struct {
	*HTTP         // Embeds the HTTP message
	method string // HTTP Method (e.g. GET, POST, PATCH, DELETE)
	path   string // Path of the requested resource
}

// Parse the incoming request
func ParseRequest(conn net.Conn) *Request {
	// Instantiate the Request
	request := &Request{
		HTTP: createHTTPMessage(),
	}

	// Read the incoming HTTP request
	// reqMsg, err := io.ReadAll(conn)
	// Note: io.ReadAll() expects an EOF to stop reading
	// The HTTP request here does not have an EOF, so this was throwing an error
	// and causing the tests to fail.
	// The following is a makeshift solution.
	// TODO: Find a better way to do this
	buf := make([]byte, 512)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request from connection: ", err.Error())
	}
	msg := string(buf)

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
