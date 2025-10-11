package core

import (
	"net/http"
	"net/url"
)

// Request is the interface that represents an HTTP request.
type Request interface {
	// BasicAuth returns the username and password from the request's Basic Authentication header.
	// If the header is not present, it returns empty strings and false.
	BasicAuth() (string, string, bool)
	// Body returns the request body as a byte slice.
	Body() ([]byte, error)
	// ClientIP returns the IP address of the client making the request.
	ClientIP() string
	// ContentLength returns the length of the request body in bytes.
	ContentLength() int64
	// Cookie retrieves a cookie by name from the request.
	Cookie(string) (*http.Cookie, error)
	// Cookies returns all cookies from the request.
	Cookies() []*http.Cookie
	// Header retrieves a header value by name from the request.
	Header(string) string
	// HeaderValues retrieves all values for a header by name from the request.
	HeaderValues(string) []string
	// Headers returns all headers from the request.
	Headers() http.Header
	// Method returns the HTTP method of the request (e.g., GET, POST).
	Method() string
	// Protocol returns the HTTP protocol version of the request (e.g., HTTP/1.1).
	Protocol() string
	// Path returns the request path.
	Path() string
	// PathValue retrieves a path parameter value by name from the request.
	PathValue(string) string
	// SetPathValue sets a path parameter value by name in the request.
	SetPathValue(string, string)
	// Resource returns the resource pattern of the request.
	Resource() string
	// SetResource sets the resource pattern of the request.
	SetResource(string)
	// Query retrieves a query parameter value by name from the request.
	Query(string) string
	// QueryValues retrieves all values for a query parameter by name from the request.
	QueryValues(string) []string
	// Queries returns all query parameters from the request as a url.Values.
	Queries() url.Values
}
