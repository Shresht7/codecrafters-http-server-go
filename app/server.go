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

	// Accept a connection
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// HTTP Response is made up of three parts, each separated by a [CRLF](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) (`\r\n`):
	// 1. Status line: `HTTP/1.1 200 OK`
	// 2. One or more Headers: `Content-Type: text/html`
	// 3. (Optional) Body: `<!DOCTYPE html><html><body><h1>Hello, World!</h1></body></html>`

	// Create the HTTP response
	statusLine := "HTTP/1.1 200 OK"
	headers := ""
	body := ""
	CRLF := "\r\n"
	response := strings.Join([]string{statusLine, headers, body}, CRLF)

	// Respond to the connection
	conn.Write([]byte(response))
}
