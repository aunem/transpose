# Transpose

Transpose is a cloud native composable proxy with a focus on kubernetes written in go.

* **NOTE:** Transpose is pre alpha and under heavy development

## Tenants

* Composable
* Observable
* Simple
* Lightwieght
* Developer friendly
* Cloud native

## Getting Started

##### Install
`go get -u github.com/aunem/transpose`   

Transpose depends on [dep](github.com/golang/dep)

##### Example
```yaml
apiVersion: alpha.aunem.com/v1
Kind: Transpose
Metadata:
  name: myProxy
  namespace: default
spec:
  listener:
    name: myHttplistener
    package: github.com/aunem/transpose-plugins/listener/http
    spec: 
      port: 80
      ssl: false

  middleware:
    request:
    - name: hydraAuth
      package: github.com/aunem/transpose-plugins/middleware/hydra
      spec:
        authUrl: my.auth.com
        clientID: transposeClient

  roundtrip:
    name: myroundtrip
    package: github.com/aunem/transpose-plugins/roundtrip/supermux
    spec:
      http:
      - path: "/"
        backend:
          serviceName: myservice
          servicePort: 80
```

see [start.md](docs/start.md) for more details

## Plugins

### Listener
Plugin | Description | Build | Contexts Supported
--- | --- | --- | ---
[http](github.com/aunem/transpose-plugins/listener/http)| simple http listener | build data | http

### Middleware
Plugin | Description | Build | Contexts Supported
--- | --- | --- | ---
[hydra](github.com/aunem/transpose-plugins/middleware/hydra)| hydra auth middleware | build data | http

### Roundtrip
Plugin | Description | Build | Contexts Supported
--- | --- | --- | ---
[supermux](github.com/aunem/transpose-plugins/roundtrip/supermux)| an enhanced router | build data | http, grpc

    
       
To develop plugins see [developing_plugins.md](docs/developing_plugins.md)

## Inspiration

* Envoy [github.com/envoyproxy/envoy](github.com/envoyproxy/envoy)
* Traefik [github.com/containous/traefik](github.com/containous/traefik)
* Gentleman [github.com/h2non/gentleman](github.com/h2non/gentleman)
* Istio [github.com/istio/istio](github.com/istio/istio)
* OpenFaas [github.com/openfaas/faas](github.com/openfaas/faas)
* Fluentd [github.com/fluent/fluentd](github.com/fluent/fluentd)

## Roadmap

- [ ] Middleware plugins   
- [x] Roundtrip plugins   
- [x] Listener plugins
- [ ] HTTP/2   
- [ ] GRPC   
- [ ] Data plane    
- [ ] Egress 

## Developing

Development happens out of the [Makefile](./Makefile). The targets are fairly simple but if you have any issues contact us.

## Contact
