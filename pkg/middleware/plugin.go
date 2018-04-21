package middleware

import "github.com/aunem/transpose/pkg/context"

// Plugin is the interface that middleware plugins must implement
type Plugin interface {
	// ProcessRequest will be called on incomming requests
	ProcessRequest(req context.Request) (context.Request, error)
	// ProcessResponse will be called on incoming responses
	ProcessResponse(resp context.Response) (context.Response, error)
	// Init is called on initialization and config updates
	Init(spec interface{}) error
	// Stats can return arbitrary json stats
	Stats() (error, []byte)
}
