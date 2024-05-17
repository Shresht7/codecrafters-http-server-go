package handlers

import (
	"fmt"
	"net/http"

	httpMessage "github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Handles the `/user-agent` endpoint.
// Extracts the `User-Agent` header from the request headers and returns it as the response body.
// If the `User-Agent` header is not found, sets the response status to 404.
func UserAgent(req *httpMessage.Request, res *httpMessage.Response) {

	// Extract the `User-Agent` header from the request headers
	userAgent, ok := req.Headers.Get("User-Agent")

	// If the `User-Agent` header is not found, set the response status to 404
	if !ok {
		res.WithStatus(http.StatusNotFound)
		return
	}

	// Set the response status to 200, content type to "text/plain",
	// content length to the length of the user agent, and body to the user agent
	res.
		WithStatus(http.StatusOK).
		WithHeaders(map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(userAgent)),
		}).
		WithBody(userAgent)

}
