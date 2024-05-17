package handlers

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Files handles requests to the /files/{name} endpoint
// It reads the file content from the --directory and returns it as the response body
func Files(req *http.Request, res *http.Response) {

	// Get the --directory from the arguments
	directory := GetDirectoryFromArguments()

	// Cut the prefix "/files/" from the request path
	fileName, found := strings.CutPrefix(req.Path, "/files/")
	if !found {
		res.WithStatus(404)
		return
	}

	// Construct the full file path
	filePath := path.Join(directory, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		res.WithStatus(404)
		return
	}

	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		res.WithStatus(500).WithBody("Internal Server Error: Could not read file")
		return
	}

	// Set the response status to 200, content type to "application/octet-stream",
	// content length to the length of the file content, and body to the file content
	res.WithStatus(200).WithHeaders(map[string]string{
		"Content-Type":   "application/octet-stream",
		"Content-Length": fmt.Sprintf("%d", len(content)),
	}).WithBody(string(content))

}

// ----------------
// HELPER FUNCTIONS
// ----------------

// Extracts the directory from the command line arguments
func GetDirectoryFromArguments() string {
	// Get Command Line Arguments
	args := os.Args[1:]

	// Extract --directory flag
	directory := ""

	i := 0
	for i < len(args) {
		if args[i] == "--directory" {
			if i+1 < len(args) {
				directory = args[i+1]
				i++ // Skip the next argument as we just used it as the value
			}
		}
		i++
	}

	return directory
}
