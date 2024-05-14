package http

import "testing"

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
		if http.Headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.Headers[key])
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
		if http.Headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.Headers[key])
		}
	}

	// Check if the previous headers are still set
	for key, value := range headers {
		if http.Headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.Headers[key])
		}
	}

	// Check if the total number of headers is correct
	if len(http.Headers) != 4 {
		t.Errorf("Expected 4 headers, but got %d", len(http.Headers))
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
		if http.Headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.Headers[key])
		}
	}

	// Check if the total number of headers is correct
	if len(http.Headers) != 2 {
		t.Errorf("Expected 2 headers, but got %d", len(http.Headers))
	}

	// Check if the previous headers are removed
	for key := range headers {
		if _, ok := http.Headers[key]; !ok {
			t.Errorf("Expected header %s to be removed, but it is still present", key)
		}
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
	message := "GET / HTTP/1.1\r\nContent-Type: application/json\r\nAuthorization: Bearer token\r\n\r\n{\"key\": \"value\"}"

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET / HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	if http.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected header Content-Type: application/json, but got %s", http.Headers["Content-Type"])
	}
	if http.Headers["Authorization"] != "Bearer token" {
		t.Errorf("Expected header Authorization: Bearer token, but got %s", http.Headers["Authorization"])
	}
	// Check if the total number of headers is correct
	if len(http.Headers) != 2 {
		t.Errorf("Expected 2 headers, but got %d", len(http.Headers))
	}

	// Check if the body is set correctly
	if http.Body != "{\"key\": \"value\"}" {
		t.Errorf("Expected body {\"key\": \"value\"}, but got %s", http.Body)
	}
}

func TestParseMessageText(t *testing.T) {
	message := "GET / HTTP/1.1\r\nContent-Type: text/plain\r\n\r\nHello, World!"

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET / HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	if http.Headers["Content-Type"] != "text/plain" {
		t.Errorf("Expected header Content-Type: text/plain, but got %s", http.Headers["Content-Type"])
	}
	// Check if the total number of headers is correct
	if len(http.Headers) != 1 {
		t.Errorf("Expected 1 header, but got %d", len(http.Headers))
	}

	// Check if the body is set correctly
	if http.Body != "Hello, World!" {
		t.Errorf("Expected body Hello, World!, but got %s", http.Body)
	}
}

func TestParseMessageHTML(t *testing.T) {
	message := "GET /index.html HTTP/1.1\r\nContent-Type: text/html\r\n\r\n<!DOCTYPE html><html><head><title>Hello, World!</title></head><body><h1>Hello, World!</h1></body></html>"

	http := createHTTPMessage().ParseMessage(message)

	// Check if the start line is set correctly
	if http.StartLine != "GET /index.html HTTP/1.1" {
		t.Errorf("Expected start line GET / HTTP/1.1, but got %s", http.StartLine)
	}

	// Check if the headers are set correctly
	if http.Headers["Content-Type"] != "text/html" {
		t.Errorf("Expected header Content-Type: text/html, but got %s", http.Headers["Content-Type"])
	}
	// Check if the total number of headers is correct
	if len(http.Headers) != 1 {
		t.Errorf("Expected 1 header, but got %d", len(http.Headers))
	}

	// Check if the body is set correctly
	if http.Body != "<!DOCTYPE html><html><head><title>Hello, World!</title></head><body><h1>Hello, World!</h1></body></html>" {
		t.Errorf("Expected body <!DOCTYPE html><html><head><title>Hello, World!</title></head><body><h1>Hello, World!</h1></body></html>, but got %s", http.Body)
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
	if http.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected header Content-Type: application/json, but got %s", http.Headers["Content-Type"])
	}
	if http.Headers["Authorization"] != "Bearer token" {
		t.Errorf("Expected header Authorization: Bearer token, but got %s", http.Headers["Authorization"])
	}
	// Check if the total number of headers is correct
	if len(http.Headers) != 2 {
		t.Errorf("Expected 2 headers, but got %d", len(http.Headers))
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
			"Content-Type":  "application/json",
			"Authorization": "Bearer token",
		}).
		WithBody("{\"key\": \"value\"}")

	expected := "GET / HTTP/1.1\r\nContent-Type: application/json\r\nAuthorization: Bearer token\r\n\r\n{\"key\": \"value\"}"

	// Check if the string representation is correct
	if http.String() != expected {
		t.Errorf("Expected string\n%s,\n\nbut got\n%s", expected, http.String())
	}
}
