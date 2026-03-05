package core

// HandlerFunc defines a function to serve HTTP requests.
type HandlerFunc func(Context) error

// HandlersChain defines a slice of HandlerFunc.
type HandlersChain []HandlerFunc
