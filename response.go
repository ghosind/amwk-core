package core

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
	// Write writes data to the response body.
	Write([]byte) (int, error)
	// Status sets the HTTP status code for the response and returns an error if it fails.
	Status(int) error
}
