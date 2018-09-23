
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/product/product.proto

run: proto

	bash -c "trap 'go run *.go --registry_address=127.0.0.1:8500' EXIT"

build: proto

	go build -o product-srv main.go plugin.go

test:
	go test -v ./... -cover

clean:
	rm -f product-srv

docker:

	CGO_ENABLED=0 GOOS=linux go build -o product-srv -a -installsuffix cgo .

	docker build . -t product-srv:latest

	rm -f product-srv

docker_run:
	docker run --rm -e MICRO_REGISTRY_ADDRESS=consul:8500 -e MONGO_HOST=mongo:27017 --network book-store --name product-srv -d product-srv

docker_stop:
	docker stop product-srv