
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/verification/verification.proto

run: proto

	bash -c "trap 'go run *.go --registry_address=127.0.0.1:8500' EXIT"

build: proto

	go build -o verification-srv main.go plugin.go

test:
	go test -v ./... -cover

clean:
	rm -f verification-srv

docker:

	CGO_ENABLED=0 GOOS=linux go build -o verification-srv -a -installsuffix cgo .

	docker build . -t verification-srv:latest

	rm -f verification-srv

docker_run:
	docker run --rm -e MICRO_REGISTRY_ADDRESS=consul:8500 -e REDIS_HOST=redis:6379 --network book-store --name verification-srv -d verification-srv

docker_stop:
	docker stop verification-srv
