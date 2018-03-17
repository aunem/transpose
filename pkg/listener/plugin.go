package listener

// Plugin is the interface that listener plugins must implement
type Plugin interface {
	// Listen
	Listen(spec interface{}) error
	// Stats can return arbitrary json stats
	Stats() ([]byte, error)
}
