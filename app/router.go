package main

import "strings"

// Route the request to the correct handler
func route(req *Request, res *Response) {
	switch {
	case req.path == "/":
		res.WithStatus(200)
	case req.path == "/user-agent":
		handleUserAgent(req, res)
	case strings.HasPrefix(req.path, "/echo/"):
		handleEcho(req, res)
	default:
		res.WithStatus(404)
	}
}
