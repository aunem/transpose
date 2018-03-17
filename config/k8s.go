package config

import (
	"github.com/ericchiang/k8s"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// Transpose represnts a configuration for an transpose instance or instances
type Transpose struct {
	Metadata   *metav1.ObjectMeta `json:"metadata"`
	Debug      bool               `json:"debug" yaml:"debug"`
	SSL        bool               `json:"ssl" yaml:"ssl"`
	Port       string             `json:"port" yaml:"port"`
	Listener   ListenerPlugin     `json:"listener" yaml:"listener"`
	Middleware MiddlewarePlugins  `json:"middleware" yaml:"middleware"`
	Roundtrip  RoundtripPlugin    `json:"roundtrip" yaml:"roundtrip"`
}

// ListenerPlugin holds the configuration for the listener plugin
type ListenerPlugin struct {
	Name    string      `json:"name" yaml:"name"`
	Package string      `json:"package" yaml:"package"`
	Spec    interface{} `json:"spec" yaml:"spec"`
}

// MiddlewarePlugins holds the middleware plugin information
type MiddlewarePlugins struct {
	Request  []MiddlewarePlugin `json:"request" yaml:"request"`
	Response []MiddlewarePlugin `json:"response" yaml:"response"`
}

// MiddlewarePlugin holds the configuration for middleware plugins
type MiddlewarePlugin struct {
	Name    string      `json:"name" yaml:"name"`
	Package string      `json:"package" yaml:"package"`
	Spec    interface{} `json:"spec" yaml:"spec"`
}

// RoundtripPlugin holds the configuratin for the server plugin
type RoundtripPlugin struct {
	Name    string      `json:"name" yaml:"name"`
	Package string      `json:"package" yaml:"package"`
	Spec    interface{} `json:"spec" yaml:"spec"`
}

// GetMetadata required for Transpose to implement k8s.Resource
func (o *Transpose) GetMetadata() *metav1.ObjectMeta {
	return o.Metadata
}

// TransposeList is a list of transposes
type TransposeList struct {
	Metadata *metav1.ListMeta `json:"metadata"`
	Items    []Transpose      `json:"items"`
}

// GetMetadata required for TransposeList to implement k8s.ResourceList
func (m *TransposeList) GetMetadata() *metav1.ListMeta {
	return m.Metadata
}

func init() {
	// Register resources with the k8s package.
	k8s.Register("alpha.transpose.com", "v1", "transpose", true, &Transpose{})
	k8s.RegisterList("alpha.transpose.com", "v1", "transpose", true, &TransposeList{})
}
