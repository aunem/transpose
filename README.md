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
    response:
    - name: hydraAuth
      package: github.com/aunem/transpose-plugins/middleware/hydra
      spec:
        authUrl: my.auth.com
        clientID: transposeClient
        auditUrl: my.audit.com

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
 [http](github.com/aunem/transpose-plugins/tree/master/listener/http)| simple http listener | build data | http

### Middleware
Plugin | Description | Build | Contexts Supported
--- | --- | --- | ---

### Roundtrip
Plugin | Description | Build | Contexts Supported
--- | --- | --- | ---
 [supermux](github.com/aunem/transpose-plugins/tree/master/roundtrip/supermux)| an enhanced router | build data | http, grpc

To develop plugins see [developing_plugins.md](docs/developing_plugins.md)

## Inspiration

* Envoy [github.com/envoyproxy/envoy](github.com/envoyproxy/envoy)
* Traefik [github.com/containous/traefik](github.com/containous/traefik)
* Gentleman [github.com/h2non/gentleman](github.com/h2non/gentleman)
* Istio [github.com/istio/istio](github.com/istio/istio)
* OpenFaas [github.com/openfaas/faas](github.com/openfaas/faas)

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
