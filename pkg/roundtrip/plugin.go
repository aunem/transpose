package roundtrip

import "github.com/aunem/transpose/pkg/context"

// Plugin is the interface that roundtrip plugins must implement
type Plugin interface {
	// Roundtrip handles interactions with the downstream service/s
	Roundtrip(req context.Request) (context.Response, error)
	// Init is ran once on plugin initialization
	Init() error
	// LoadSpec loads the spec on initialization and config updates
	LoadSpec(spec interface{}) error
	// Stats can return arbitrary json stats
	Stats() ([]byte, error)
}
