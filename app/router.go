package main

import (
	"strings"

	handle "github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Route the request to the correct handler
func route(req *http.Request, res *http.Response) {
	switch {
	case req.Path == "/":
		res.WithStatus(200)
	case strings.HasPrefix(req.Path, "/files/"):
		handle.Files(req, res)
	case req.Path == "/user-agent":
		handle.UserAgent(req, res)
	case strings.HasPrefix(req.Path, "/echo/"):
		handle.Echo(req, res)
	default:
		res.WithStatus(404)
	}
}
