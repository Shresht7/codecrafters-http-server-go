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
	for k, v := range h.hashmap {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}
	return "", false
}

// Check if a header is present in the Headers object
func (h *Headers) Contains(key string) bool {
	_, found := h.Get(key)
	return found
}

// Delete a header from the Headers object
func (h *Headers) Delete(key string) {
	for k := range h.hashmap {
		if strings.EqualFold(k, key) {
			delete(h.hashmap, k)
			return
		}
	}
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
