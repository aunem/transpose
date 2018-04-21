package roundtrip

// Manifest is a map of file name to template string
var Manifest = map[string]string{
	"main.go":             mainTpl,
	"main_test.go":        mainTestTpl,
	"example-config.yaml": exampleTpl,
	"README.md":           readmeTpl,
	"spec.go":             specTpl,
}
