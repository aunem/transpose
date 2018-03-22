package listener

import "github.com/aunem/transpose/config"

// Plugin is the interface that listener plugins must implement
type Plugin interface {
	// Listen
	Listen(conf config.TransposeSpec) error
	// Stats can return arbitrary json stats
	Stats() ([]byte, error)
}
