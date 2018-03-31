package middleware

// Manifest represents the variables to inclued in the templating
var Manifest = map[string]string{
	"main.go":             mainTpl,
	"main_test.go":        mainTestTpl,
	"example-config.yaml": exampleTpl,
	"README.md":           readmeTpl,
	"spec.go":             specTpl,
}
