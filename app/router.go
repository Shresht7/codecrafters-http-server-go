package main

import (
	"net/http"
	"strings"

	handle "github.com/codecrafters-io/http-server-starter-go/app/handlers"
	httpMessage "github.com/codecrafters-io/http-server-starter-go/pkg/http"
)

// Route the request to the correct handler
func route(req *httpMessage.Request, res *httpMessage.Response) {
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
		res.WithStatus(http.StatusOK)

	// Default
	default:
		res.WithStatus(http.StatusNotFound)
	}

}
