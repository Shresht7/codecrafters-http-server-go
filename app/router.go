package main

import (
	"strings"

	handle "github.com/codecrafters-io/http-server-starter-go/app/handlers"
	"github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Route the request to the correct handler
func route(req *http.Request, res *http.Response) {
	switch {

	// /files/{name}
	case strings.HasPrefix(req.Path, "/files/"):
		handle.Files(req, res)

	// /user-agent
	case req.Path == "/user-agent":
		handle.UserAgent(req, res)

	// /echo/{str}
	case strings.HasPrefix(req.Path, "/echo/"):
		handle.Echo(req, res)

	// /
	case req.Path == "/":
		res.WithStatus(200)

	// Default
	default:
		res.WithStatus(404)
	}

}
