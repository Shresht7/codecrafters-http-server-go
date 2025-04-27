package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

func main() {
	// Bind to port 4221
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	// Accept connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		// Handle the connection in a new goroutine
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently
		go handleConnection(conn)
	}

}

// handleConnection handles an incoming connection.
// It parses the HTTP request, creates an HTTP response, routes the request,
// and responds to the connection.
func handleConnection(conn net.Conn) {
	// Close the connection when the function returns
	defer conn.Close()

	// Setup a persistent connection until we get a "connection-close"
	for {
		// Parse the HTTP Request from the connection
		request := http.ParseRequest(conn)
		if request == nil {
			break // Break out of the persistent connection if request is nil
		}

		// Print the request
		fmt.Printf("Request:\n%+v\n", request)

		// Create the HTTP Response
		response := http.CreateResponse()

		// Route the request based on the requested path
		route(request, response)

		// Print the response
		fmt.Println("Response:\n", response.String())

		// Respond to the connection
		conn.Write(response.Bytes())

		// Close the connection if the `Connection: close` header was passed in
		if val, ok := request.Headers.Get("Connection"); ok && val == "close" {
			break // Close the connection
		}
	}
}
