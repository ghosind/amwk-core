package core

import "context"

// App is the interface for a web application.
type App interface {
	// Run runs the application server.
	Run() error
	// RunTLS runs the application server with TLS.
	RunTLS() error
	// Use adds the given handlers to the application.
	Use(handlers ...HandlerFunc) App
	// Close closes the application.
	Close() error
	// Shutdown gracefully shuts down the application server.
	Shutdown(ctx context.Context) error
}
