package main

import "testing"

func TestParseRequestLine(t *testing.T) {
	testCases := []struct {
		startLine string
		method    string
		path      string
		protocol  string
	}{
		{
			startLine: "GET / HTTP/1.1",
			method:    "GET",
			path:      "/",
			protocol:  "HTTP/1.1",
		},
		{
			startLine: "POST /api/v1/users HTTP/1.1",
			method:    "POST",
			path:      "/api/v1/users",
			protocol:  "HTTP/1.1",
		},
		{
			startLine: "PATCH /api/v1/users/1 HTTP/1.1",
			method:    "PATCH",
			path:      "/api/v1/users/1",
			protocol:  "HTTP/1.1",
		},
	}

	for _, tc := range testCases {
		req := &Request{
			HTTPMessage: createHTTPMessage().WithStartLine(tc.startLine),
		}

		req.parseRequestLine()

		if req.method != tc.method {
			t.Errorf("Expected method %s, but got %s", tc.method, req.method)
		}

		if req.path != tc.path {
			t.Errorf("Expected path %s, but got %s", tc.path, req.path)
		}

		if req.protocol != tc.protocol {
			t.Errorf("Expected protocol %s, but got %s", tc.protocol, req.protocol)
		}
	}
}
