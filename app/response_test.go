package main

import "testing"

func TestResponse_WithStatus(t *testing.T) {
	testCases := []struct {
		name     string
		code     int
		expected string
	}{
		{
			name:     "Status OK",
			code:     200,
			expected: "HTTP/1.1 200 OK",
		},
		{
			name:     "Status Not Found",
			code:     404,
			expected: "HTTP/1.1 404 Not Found",
		},
		{
			name:     "Status Internal Server Error",
			code:     500,
			expected: "HTTP/1.1 500 Internal Server Error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := createResponse().WithStatus(tc.code)
			if r.startLine != tc.expected {
				t.Errorf("Expected start line to be %q, but got %q", tc.expected, r.startLine)
			}
		})
	}
}
