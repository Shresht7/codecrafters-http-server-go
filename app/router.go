package main

import (
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Route the request to the correct handler
func route(req *http.Request, res *http.Response) {
	switch {
	case req.Path == "/":
		res.WithStatus(200)
	case strings.HasPrefix(req.Path, "/files/"):
		handleFile(req, res)
	case req.Path == "/user-agent":
		handleUserAgent(req, res)
	case strings.HasPrefix(req.Path, "/echo/"):
		handleEcho(req, res)
	default:
		res.WithStatus(404)
	}
}
