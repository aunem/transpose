
.PHONY: deps
deps:
	go get -u github.com/golang/dep/cmd/dep
	docker-compose run --rm deps

.PHONY: build
build:
	docker-compose run --rm build

.PHONY: plugins
plugins:
	docker-compose run --rm plugins

.PHONY: integration
integration: plugins
	docker-compose run --rm integration

.PHONY: integration-down
integration-down:
	docker-compose down

.PHONY: clean
clean:
	go run main.go plugin clean

.PHONY: proto-gateway
proto-gateway: 
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true:. \
	api/transpose.proto

.PHONY: proto-swagger
proto-swagger:
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--swagger_out=logtostderr=true:. \
	api/transpose.proto

.PHONY: proto-client
proto-client:
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:. \
	api/transpose.proto

.PHONY: proto
proto:
	protoc -I/usr/local/include -I. \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:$(GOPATH)/src \
	api/transpose.proto

gen-proto: proto proto-gateway proto-swagger proto-client