package roundtrip

var readmeTpl = `

# {{ .Name }}

// Your overview here

## Spec
` + "```yaml " + `
apiVersion: alpha.aunem.io/v1
Kind: Transpose
Metadata:
  name: myProxy
  namespace: default
spec:
  listener:
    name: mylistener
    package: github.com/aunem/transpose-plugins/listener/http
    spec: 
      port: 80
      ssl: false

  roundtrip:
    name: {{ .Name }}
    package: {{ .Pkg }}
    spec:
      http:
      - path: "/"
        backend:
          serviceName: myservice
          servicePort: 80
` + "```" + `

## Contexts Supported
* Http

## Compatibility
// your compatibility data goes here

## Dependencies
// your dependencies go here

## Test
` + "`go test ./...`"
