package handlers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Handles the `/echo/{str}` endpoint.
// Extracts the string from the request path and returns it as the response body.
func Echo(req *http.Request, res *http.Response) {

	// Remove the prefix `/echo/` from the request path (e.g. `/echo/hello` -> `hello`)
	str, found := strings.CutPrefix(req.Path, "/echo/")

	// If the prefix is not found, set the response status to 404
	if !found {
		res.WithStatus(404)
		return
	}

	// If the request contains the `Accept-Encoding` header with the value "gzip"...
	acceptEncoding, ok := req.Headers.Get("Accept-Encoding")
	if ok && strings.Contains(acceptEncoding, "gzip") {

		// Encode the string using the gzip algorithm
		gzStr, err := GZip(str)
		if err != nil {
			res.WithStatus(500)
			return
		}
		str = gzStr // Set the contents to the compressed string

		// Set the response header "Content-Encoding" to "gzip"
		res.WithHeaders(map[string]string{
			"Content-Encoding": "gzip",
		})

	}

	// Respond with the contents
	res.
		WithStatus(200).
		WithHeaders(map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(str)),
		}).
		WithBody(str)

}

// Compress a string using the gzip algorithm.
func GZip(str string) (string, error) {

	var buf bytes.Buffer       // Create a buffer to store the compressed data
	gz := gzip.NewWriter(&buf) // Create a new gzip writer

	// Write the string to the gzip writer
	if _, err := gz.Write([]byte(str)); err != nil {
		return "", err
	}

	// Close the gzip writer
	if err := gz.Close(); err != nil {
		return "", err
	}

	// Flush the gzip writer
	return buf.String(), nil

}
