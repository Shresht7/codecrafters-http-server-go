package http

import (
	"strings"
	"testing"
)

// ----------
// START LINE
// ----------

func TestWithStartLine(t *testing.T) {
	startLine := "GET / HTTP/1.1"

	http := createHTTPMessage().WithStartLine(startLine)

	// Check if the start line is set correctly
	if http.StartLine != startLine {
		t.Errorf("Expected start line %s, but got %s", startLine, http.StartLine)
	}
}

// -------
// HEADERS
// -------

func TestWithHeaders(t *testing.T) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token",
	}

	http := createHTTPMessage().WithHeaders(headers)

	// Check if the headers are set correctly
	for key, value := range headers {
		h, ok := http.Headers.Get(key)
		if !ok || h != value {
			t.Errorf("Expected header value %s, but got %s", value, h)
		}
	}
}

func TestWithHeadersMultiple(t *testing.T) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token",
	}

	http := createHTTPMessage().WithHeaders(headers)

	// Add more headers
	moreHeaders := map[string]string{
		"X-Request-ID": "12345",
		"User-Agent":   "Go-http-client/1.1",
	}

	http.WithHeaders(moreHeaders)

	// Check if the headers are set correctly
	for key, value := range moreHeaders {
		h, ok := http.Headers.Get(key)
		if !ok || h != value {
			t.Errorf("Expected header value %s, but got %s", value, h)
		}
	}

	// Check if the previous headers are still set
	for key, value := range headers {
		h, ok := http.Headers.Get(key)
		if !ok || h != value {
			t.Errorf("Expected header value %s, but got %s", value, h)
		}
	}

	// Check if the total number of headers is correct
	if http.Headers.Len() != 4 {
		t.Errorf("Expected 4 headers, but got %d", http.Headers.Len())
	}
}

func TestWithHeadersOverride(t *testing.T) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token",
	}

	http := createHTTPMessage().WithHeaders(headers)

	// Override a header
	updatedHeaders := map[string]string{
		"Content-Type": "application/xml",
	}

	http.WithHeaders(updatedHeaders)

	// Check if the headers are set correctly
	for key, value := range updatedHeaders {
		h, ok := http.Headers.Get(key)
		if !ok || h != value {
			t.Errorf("Expected header value %s, but got %s", value, h)
		}
	}

	// Check if the total number of headers is correct
	if http.Headers.Len() != 2 {
		t.Errorf("Expected 2 headers, but got %d", http.Headers.Len())
	}

}

// ----
// BODY
// ----

func TestWithBody(t *testing.T) {
	body := "Hello, World!"

	http := createHTTPMessage().WithBody(body)

	// Check if the body is set correctly
	if http.Body != body {
		t.Errorf("Expected body %s, but got %s", body, http.Body)
	}
}

// -------------
// PARSE MESSAGE
// -------------

func TestParseMessage(t *testing.T) {
	message := strings.Join([]string{
		"GET / HTTP/1.1",
		"Content-Type: application/json",
		"Content-Length: 16",
		"Authorization: Bearer token",
		"",
		"{\"key\": \"value\"}",
	}, CRLF)

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET / HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	contentType, _ := http.Headers.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected header Content-Type: application/json, but got %s", contentType)
	}
	authorization, _ := http.Headers.Get("Authorization")
	if authorization != "Bearer token" {
		t.Errorf("Expected header Authorization: Bearer token, but got %s", authorization)
	}
	// Check if the total number of headers is correct
	if http.Headers.Len() != 3 {
		t.Errorf("Expected 2 headers, but got %d", http.Headers.Len())
	}

	// Check if the body is set correctly
	if http.Body != "{\"key\": \"value\"}" {
		t.Errorf("Expected body {\"key\": \"value\"}, but got %s", http.Body)
	}
}

func TestParseMessageText(t *testing.T) {
	// message := "GET / HTTP/1.1\r\nContent-Type: text/plain\r\n\r\nHello, World!"
	message := strings.Join([]string{
		"GET / HTTP/1.1",
		"Content-Type: text/plain",
		"Content-Length: 13",
		"",
		"Hello, World!",
	}, CRLF)

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET / HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	contentType, _ := http.Headers.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("Expected header Content-Type: text/plain, but got %s", contentType)
	}
	// Check if the total number of headers is correct
	if http.Headers.Len() != 2 {
		t.Errorf("Expected 1 header, but got %d", http.Headers.Len())
	}

	// Check if the body is set correctly
	if http.Body != "Hello, World!" {
		t.Errorf("Expected body Hello, World!, but got %s", http.Body)
	}
}

func TestParseMessageHTML(t *testing.T) {
	message := strings.Join([]string{
		"GET /index.html HTTP/1.1",
		"Content-Type: text/html",
		"Content-Length: 96",
		"",
		"<!DOCTYPE html><html><head><title>TITLE</title></head><body><h1>Hello, World!</h1></body></html>",
	}, CRLF)

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET /index.html HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	contentType, _ := http.Headers.Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("Expected header Content-Type: text/html, but got %s", contentType)
	}
	// Check if the total number of headers is correct
	if http.Headers.Len() != 2 {
		t.Errorf("Expected 1 header, but got %d", http.Headers.Len())
	}

	// Check if the body is set correctly
	doc := "<!DOCTYPE html><html><head><title>TITLE</title></head><body><h1>Hello, World!</h1></body></html>"
	if http.Body != doc {
		t.Errorf("Expected body:\n\n"+doc+"\n\n", http.Body)
	}
}

func TestParseMessageWithNoBody(t *testing.T) {
	message := "GET / HTTP/1.1\r\nContent-Type: application/json\r\nAuthorization: Bearer token\r\n"

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET / HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	contentType, _ := http.Headers.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected header Content-Type: application/json, but got %s", contentType)
	}
	authorization, _ := http.Headers.Get("Authorization")
	if authorization != "Bearer token" {
		t.Errorf("Expected header Authorization: Bearer token, but got %s", authorization)
	}
	// Check if the total number of headers is correct
	if http.Headers.Len() != 2 {
		t.Errorf("Expected 2 headers, but got %d", http.Headers.Len())
	}

	// Check if the body is set correctly
	if http.Body != "" {
		t.Errorf("Expected empty body, but got %s", http.Body)
	}

}

// ------
// STRING
// ------

func TestString(t *testing.T) {
	http := createHTTPMessage().
		WithStartLine("GET / HTTP/1.1").
		WithHeaders(map[string]string{
			"Content-Type":   "application/json",
			"Content-Length": "16",
			"Authorization":  "Bearer token",
		}).
		WithBody("{\"key\": \"value\"}")

	expected := strings.Join([]string{
		"GET / HTTP/1.1",
		"Content-Type: application/json",
		"Content-Length: 16",
		"Authorization: Bearer token",
		"",
		"{\"key\": \"value\"}",
	}, CRLF)

	// Check if the string representation is correct
	if http.String() != expected {
		t.Errorf("Expected string\n%s\n\nbut got\n%s", strings.TrimSpace(expected), http.String())
	}
}
