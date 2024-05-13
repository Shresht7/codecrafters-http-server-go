package main

import (
	"fmt"

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

	// Read the incoming HTTP request
	// reqMsg, err := io.ReadAll(conn)
	// Note: io.ReadAll() expects an EOF to stop reading
	// The HTTP request here does not have an EOF, so this was throwing an error
	// and causing the tests to fail.
	// The following is a makeshift solution.
	buf := make([]byte, 512)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request from connection: ", err.Error())
	}
	reqMsg := string(buf)

	// Parse the HTTP Request
	request := ParseRequest(string(reqMsg))

	fmt.Printf("Request: %s --- %s\n", request.method, request.path)

	// Create the HTTP Response
	response := createResponse()

	// Route the request based on the requested path
	_, response = route(request, response)

	fmt.Println("Response:", response)

	// Respond to the connection
	conn.Write(response.Bytes())
}

// Route the request to the correct handler
func route(req *Request, res *Response) (*Request, *Response) {
	switch req.path {
	case "/":
		res.WithStatus(200)
	default:
		res.WithStatus(404)
	}
	return req, res
}
