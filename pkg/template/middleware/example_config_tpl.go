package middleware

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
    
  middleware:
    name: {{ .Name }}
    package: {{ .Pkg }}
    # Add your spec data here
    spec:
      my: spec

  roundtrip:
    name: myroundtrip
    package: github.com/aunem/transpose-plugins/roundtrip/supermux
    spec:
      http:
      - path: "/"
        backend:
          serviceName: myservice
          servicePort: 80
`
