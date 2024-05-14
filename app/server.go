package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	// Bind to port 4221
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	// Accept a connection
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// Parse the HTTP Request
	request := ParseRequest(conn)

	fmt.Printf("Request:\n%+v\n", request)

	// Create the HTTP Response
	response := createResponse()

	// Route the request based on the requested path
	route(request, response)

	fmt.Println("Response:\n", response.String())

	// Respond to the connection
	conn.Write(response.Bytes())
}

// Route the request to the correct handler
func route(req *Request, res *Response) {
	switch {
	case req.path == "/":
		res.WithStatus(200)
	case req.path == "/user-agent":
		handleUserAgent(req, res)
	case strings.HasPrefix(req.path, "/echo/"):
		handleEcho(req, res)
	default:
		res.WithStatus(404)
	}
}
