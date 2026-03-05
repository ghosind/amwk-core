package core

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// Context is the interface that represents the context of a single HTTP request.
type Context interface {
	// Get returns the value associated with the key in the context.
	Get(string) (any, bool)
	// Set sets the value for the key in the context and returns the previous value.
	Set(string, any) any
	// Context returns the context.Context associated with the request.
	Context() context.Context

	// Application returns the application instance associated with the context.
	Application() Application

	// Abort marks the context as aborted, and subsequent handlers will not be executed.
	Abort()
	// IsAbort checks if the context is marked as aborted.
	IsAbort() bool
	// Next calls the next handler in the handlers chain.
	Next() error
	// Use adds handlers to the context, which will be executed in the order they are added.
	Use(...HandlerFunc)

	// Body returns the request body as a readable stream.
	Body() (io.ReadCloser, error)
	// ClientIP returns the IP address of the client making the request.
	ClientIP() string
	// ContentLength returns the length of the request body in bytes.
	ContentLength() int64
	// ContentType returns the Content-Type header of the request.
	ContentType() string
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
	// Resource returns the resource pattern of the request.
	Resource() string
	// Query retrieves a query parameter value by name from the request.
	Query(string) string
	// QueryValues retrieves all values for a query parameter by name from the request.
	QueryValues(string) []string
	// Queries returns all query parameters from the request as a url.Values.
	Queries() url.Values

	// AddHeader adds a header to the response.
	AddHeader(string, string)
	// SetHeader sets a header in the response.
	SetHeader(string, string)
	// GetHeader retrieves a header value by name from the response.
	GetHeader(string) string
	// DelHeader deletes a header from the response.
	DelHeader(string)
	// Status sets the HTTP status code for the response and returns an error if it fails.
	Status(int) error
	// Write writes data to the response body.
	Write([]byte) (int, error)

	// Request returns the wrapped Request object associated with the context.
	Request() Request
	// Response returns the wrapped Response object associated with the context.
	Response() Response
}
