package core

import "net/http"

// Response is the interface that represents an HTTP response.
type Response interface {
	// AddHeader adds a header value for the response.
	AddHeader(string, string)
	// SetHeader sets a header value for the response.
	SetHeader(string, string)
	// GetHeader retrieves a header value by name from the response.
	GetHeader(string) string
	// DelHeader deletes a header from the response.
	DelHeader(string)
	// Headers returns all headers from the response.
	Headers() http.Header
	// Write writes data to the response body.
	Write([]byte) (int, error)
	// Status sets the HTTP status code for the response.
	Status(int)
	// StatusCode returns the current HTTP status code of the response.
	StatusCode() int

	// Response returns the underlying response object, which can be of any type depending on the implementation.
	Response() any
}
