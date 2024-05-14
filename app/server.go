package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
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
