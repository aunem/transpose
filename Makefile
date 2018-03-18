
.PHONY: deps
deps:
	docker-compose run --rm deps

.PHONY: build
	docker-compose run -rm build

.PHONY: integration
integration:
	docker-compose -f integration-compose.yaml up

.PHONY: integration-down
integration-down:
	docker-compose -f integration-compose.yaml down

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