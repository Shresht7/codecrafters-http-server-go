package main

import (
	"strings"
)

func handleEcho(req *Request, res *Response) {
	str, found := strings.CutPrefix(req.path, "/echo/")
	if !found {
		res.WithStatus(404)
		return
	}
	res.
		WithStatus(200).
		WithHeaders(map[string]string{
			"Content-Type": "text/plain",
		}).
		WithBody(str)
}
