package handlers

import (
	"bytes"
	"compress/gzip"
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
	acceptableEncodings, ok := req.Headers.Get("Accept-Encoding")
	if ok && strings.Contains(acceptableEncodings, "gzip") {
		res.WithHeaders(map[string]string{
			"Content-Encoding": "gzip",
		})
	}

	// Encode the string using the gzip algorithm
	// and set the response body to the encoded string
	content, err := GZip(str)
	if err != nil {
		res.WithStatus(500)
		return
	}

	// Set the response status to 200, content type to "text/plain",
	// content length to the length of the string, and body to the string
	res.
		WithStatus(200).
		WithHeaders(map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(content)),
		}).
		WithBody(string(content))

}

// Return the gzip encoded version of the string
func GZip(str string) ([]byte, error) {

	var buf bytes.Buffer       // Create a buffer to store the compressed data
	gz := gzip.NewWriter(&buf) // Create a new gzip writer
	defer gz.Close()           // Close the gzip writer when the function returns

	// Write the string to the gzip writer
	if _, err := gz.Write([]byte(str)); err != nil {
		return nil, err
	}

	// Flush the gzip writer
	return buf.Bytes(), nil
}
