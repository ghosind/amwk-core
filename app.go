package core

import "context"

// Application is the interface for a web application.
type Application interface {
	// Start starts the application server and listens for incoming requests. It returns an error if
	// it fails to start.
	Start() error
	// Use adds the given handlers to the application.
	Use(handlers ...HandlerFunc) Application
	// Close closes the application.
	Close() error
	// Shutdown gracefully shuts down the application server.
	Shutdown(ctx context.Context) error
}
