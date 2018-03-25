package main

type {{ .Name }}Plugin struct {}

// MiddlewarePlugin is an interface to transpose
var MiddlewarePlugin {{ .Name }}Plugin

// Spec holds the spec data
var Spec {{ .Name }}Spec

func main() {}

func (m *{{ .Name }}Plugin) ProcessRequest(req context.Request) (context.Request, error) {
    return nil, nil
}

func (m *{{ .Name }}Plugin) ProcessResponse(resp context.Response) (context.Response, error) {
	return nil, nil
}

func (m *{{ .Name }}Plugin) LoadSpec(spec interface{}) error {
	b, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &Spec)
	if err != nil {
		return err
	}
	log.Debugf("loaded spec: %+v", Spec)
	return nil
}

func (m *{{ .Name }}Plugin) Init() error {
	return nil
}

func (m *{{ .Name }}Plugin) Stats() ([]byte, error) {
	return nil, nil
}