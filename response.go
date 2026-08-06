package core

import "net/http"

// Response is the interface that represents an HTTP response.
type Response interface {
	// Application returns the application instance associated with the response.
	Application() Application
	// Body returns the response body as a readable stream.
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
	// Size returns the size of the response body in bytes.
	Size() int
	// Status sets the HTTP status code for the response.
	Status(int)
	// StatusCode returns the current HTTP status code of the response.
	StatusCode() int
	// Write writes data to the response body.
	Write([]byte) (int, error)
	// WriteString writes a string to the response body.
	WriteString(string) (int, error)
	// Written returns the data that has been written to the response body so far.
	Written() []byte

	// Response returns the underlying response object, which can be of any type depending on the implementation.
	Response() any
}
