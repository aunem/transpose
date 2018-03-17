package middleware

import "github.com/aunem/transpose/pkg/context"

// Plugin is the interface that middleware plugins must implement
type Plugin interface {
	// ProcessRequest will be called on incomming requests
	ProcessRequest(req context.Request, spec interface{}) (context.Request, error)
	// ProcessResponse will be called on incoming responses
	ProcessResponse(resp context.Response, spec interface{}) (context.Response, error)
	// Init is ran once on plugin initialization
	Init() error
	// Stats can return arbitrary json stats
	Stats() (error, []byte)
}
