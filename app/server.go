package main

import (
	"fmt"
	"io"
	"net/http"

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
	reqMsg, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Error reading request from connection: ", err.Error())
	}

	// Parse the HTTP Request
	request := ParseRequest(string(reqMsg))

	fmt.Println("Request Method:", request.method)
	fmt.Println("Request Path:", request.path)

	// Create a HTTP Response
	response := createResponse().
		WithStatus(http.StatusOK)

	// Respond to the connection
	conn.Write(response.Bytes())
}
