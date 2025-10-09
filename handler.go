package core

// HandlerFunc defines a function to serve HTTP requests.
type HandlerFunc func(Context)

// HandlersChain defines a slice of HandlerFunc.
type HandlersChain []HandlerFunc
