package listener

var exampleTpl = `apiVersion: alpha.aunem.io/v1
Kind: Transpose
Metadata:
  name: myProxy
  namespace: default
spec:
  listener:
    name: {{ .Name }}
    package: {{ .Repo }}
    spec: 
      port: 80
      ssl: false

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
