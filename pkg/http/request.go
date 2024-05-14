package http

import (
	"fmt"
	"net"
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

	// Read the incoming connection and retrieve the request message
	msg := readConnection(conn)

	// Parse the incoming message
	request.ParseMessage(msg)

	// Parse the Method and Path from the request line
	request.parseRequestLine()

	return request
}

// Parse the Method and Path from the request line. See https://datatracker.ietf.org/doc/html/rfc9112#section-3
func (r *Request) parseRequestLine() {
	s := strings.Fields(r.StartLine)
	r.Method = s[0]
	r.Path = s[1]
	r.protocol = s[2]
}

// ----------------
// HELPER FUNCTIONS
// ----------------

// Read the incoming connection and return the request message
func readConnection(conn net.Conn) string {
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
	return msg
}
