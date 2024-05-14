package main

import "testing"

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
		if http.headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.headers[key])
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
		if http.headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.headers[key])
		}
	}

	// Check if the previous headers are still set
	for key, value := range headers {
		if http.headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.headers[key])
		}
	}

	// Check if the total number of headers is correct
	if len(http.headers) != 4 {
		t.Errorf("Expected 4 headers, but got %d", len(http.headers))
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
		if http.headers[key] != value {
			t.Errorf("Expected header value %s, but got %s", value, http.headers[key])
		}
	}

	// Check if the total number of headers is correct
	if len(http.headers) != 2 {
		t.Errorf("Expected 2 headers, but got %d", len(http.headers))
	}

	// Check if the previous headers are removed
	for key := range headers {
		if _, ok := http.headers[key]; !ok {
			t.Errorf("Expected header %s to be removed, but it is still present", key)
		}
	}
}
