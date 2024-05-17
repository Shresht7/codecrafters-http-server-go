package http

import (
	"fmt"
	"strings"
)

// Headers represents the headers of a HTTP Request/Response
type Headers struct {
	hashmap map[string]string
}

// Instantiate a new Headers object with an empty hashmap
func NewHeaders() *Headers {
	return &Headers{
		hashmap: make(map[string]string),
	}
}

// Set a header in the Headers object
func (h *Headers) Set(key, value string) {
	h.hashmap[key] = value
}

// Get a header from the Headers object
func (h *Headers) Get(key string) (string, bool) {
	value, found := h.hashmap[key]
	return value, found
}

// Check if a header is present in the Headers object
func (h *Headers) Contains(key string) bool {
	_, found := h.Get(key)
	return found
}

// Delete a header from the Headers object
func (h *Headers) Delete(key string) {
	delete(h.hashmap, key)
}

// Returns the length of the Headers object
func (h *Headers) Len() int {
	return len(h.hashmap)
}

// Enumerate the Headers object
func (h *Headers) Enumerate() map[string]string {
	return h.hashmap
}

// Convert the Headers object to a string
func (h *Headers) String() string {
	fieldLines := make([]string, 0, h.Len())
	for key, value := range h.Enumerate() {
		fieldLines = append(fieldLines, fmt.Sprintf("%s: %s", key, value))
	}
	// Add an extra CRLF to separate the headers from the body
	return strings.Join(fieldLines, CRLF) + CRLF
}

// Check if a header is present in the Headers object and its value matches any of the checks
func (h *Headers) Check(key string, values ...string) bool {
	// Get the value of the header
	value, found := h.Get(key)
	if !found {
		return false // Return false, if the header is not present
	}

	if len(values) == 0 {
		return true // Return true, if the header is present and no checks are provided
	}

	for _, v := range values {
		if value == v {
			return true // Return true, if the header is present and the value matches any of the checks
		}
	}

	// Return false, if the header is present but the value does not match any of the checks
	return false
}
