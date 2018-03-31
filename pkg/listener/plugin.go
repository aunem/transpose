package listener

import (
	"github.com/aunem/transpose/pkg/middleware"
	"github.com/aunem/transpose/pkg/roundtrip"
)

// Plugin is the interface that listener plugins must implement
type Plugin interface {
	// Listen
	Listen(mw *middleware.Manager, rt *roundtrip.Manager) error
	// Init is called on initialization and config updates
	Init(spec interface{}) error
	// Stats can return arbitrary json stats
	Stats() ([]byte, error)
}
