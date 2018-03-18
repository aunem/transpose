# Transpose

Transpose is a cloud native composable proxy with a focus on Kubernetes written in Go.

* **NOTE:** Transpose is pre alpha and under heavy development

## Tenants

* Composable
* Observable
* Simple
* Lightwieght
* Developer friendly
* Cloud native

## Getting Started
Example config:   
```yaml
apiVersion: alpha.aunem.com/v1
Kind: Transpose
Metadata:
  name: myProxy
  namespace: default
spec:
  listener:
    name: myHttplistener
    package: github.com/oscea/transpose-plugins/listener/http
    spec: 
      port: 80
      ssl: false

  middleware:
    request:
    - name: hydraAuth
      package: github.com/oscea/transpose-plugins/middleware/hydra
      spec:
        authUrl: my.auth.com
    response:
    - name: hydraAuth
      package: github.com/oscea/transpose-plugins/middleware/hydra
      spec:
        authUrl: my.auth.com

  roundtrip:
    name: myroundtrip
    package: github.com/oscea/transpose-plugins/roundtrip/supermux
    spec:
      http:
      - path: "/"
        backend:
          serviceName: myservice
          servicePort: 80
```

see [start.md](docs/start.md) for more details

## Inspiration

* Envoy [github.com/envoyproxy/envoy](github.com/envoyproxy/envoy)
* Traefik [github.com/containous/traefik](github.com/containous/traefik)
* Gentleman [github.com/h2non/gentleman](github.com/h2non/gentleman)
* Istio [github.com/istio/istio](github.com/istio/istio)
* OpenFaas [github.com/openfaas/faas](github.com/openfaas/faas)

## Libs

* Oxy [github.com/vulcand/oxy](github.com/vulcand/oxy)
* Gorrilla Mux [github.com/gorilla/mux](github.com/gorilla/mux)

## Roadmap

- [ ] Middleware plugins   
- [ ] Roundtrip plugins   
- [ ] Listener plugins
- [ ] HTTP/2   
- [ ] GRPC   
- [ ] Data plane    
- [ ] Egress 

## Developing

Development happens out of the [Makefile](./Makefile). The targets are fairly simple but if you have any issues contact us.

## Contact
