package handlers

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Echo handles the "/echo/{str}" endpoint.
// It extracts the string from the request path and returns it as the response body.
// If the string is not found in the request path, it sets the response status to 404.
func Echo(req *http.Request, res *http.Response) {

	// Cut the prefix "/echo/" from the request path
	str, found := strings.CutPrefix(req.Path, "/echo/")

	// If the prefix is not found, set the response status to 404
	if !found {
		res.WithStatus(404)
		return
	}

	// If the request contains the header "Accept-Encoding" with the value "gzip",
	// set the response header "Content-Encoding" to "gzip"
	if req.Headers["Accept-Encoding"] == "gzip" {
		res.WithHeaders(map[string]string{
			"Content-Encoding": "gzip",
		})
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
