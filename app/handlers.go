package main

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// handleEcho handles the "/echo/{str}" endpoint.
// It extracts the string from the request path and returns it as the response body.
// If the string is not found in the request path, it sets the response status to 404.
func handleEcho(req *http.Request, res *http.Response) {

	// Cut the prefix "/echo/" from the request path
	str, found := strings.CutPrefix(req.Path, "/echo/")

	// If the prefix is not found, set the response status to 404
	if !found {
		res.WithStatus(404)
		return
	}

	// Set the response status to 200, content type to "text/plain",
	// content length to the length of the string, and body to the string
	res.
		WithStatus(200).
		WithHeaders(map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(str)),
		}).
		WithBody(str)
}

// handleUserAgent handles the "/user-agent" endpoint.
// It extracts the User-Agent header from the request headers and returns it as the response body.
// If the User-Agent header is not found, it sets the response status to 404.
func handleUserAgent(req *http.Request, res *http.Response) {
	// Extract the User-Agent header from the request headers
	userAgent, ok := req.Headers["User-Agent"]

	// If the User-Agent header is not found, set the response status to 404
	if !ok {
		res.WithStatus(404)
		return
	}

	// Set the response status to 200, content type to "text/plain",
	// content length to the length of the user agent, and body to the user agent
	res.
		WithStatus(200).
		WithHeaders(map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(userAgent)),
		}).
		WithBody(userAgent)
}
