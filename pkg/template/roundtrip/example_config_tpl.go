package roundtrip

var exampleTpl = `apiVersion: alpha.aunem.io/v1
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
`
