package context

// Request is a generic incomming request interface
type Request interface {
	GetID() string
	GetMeta() map[string]string
	GetRequest() interface{}
}

// Response is a generic returned response interface
type Response interface {
	GetID() string
	GetMeta() map[string]string
	GetRequest() interface{}
	GetResponse() interface{}
}
