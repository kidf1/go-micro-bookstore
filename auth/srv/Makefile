
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/auth/auth.proto

run: proto

	bash -c "trap 'go run *.go --registry_address=127.0.0.1:8500' EXIT"

build: proto

	go build -o auth-srv main.go plugin.go

test:
	go test -v ./... -cover

clean:
	rm -f auth-srv

docker:

	CGO_ENABLED=0 GOOS=linux go build -o auth-srv -a -installsuffix cgo .

	docker build . -t auth-srv:latest

	rm -f auth-srv

docker_run:
	docker run --rm -e MICRO_REGISTRY_ADDRESS=consul:8500 -e REDIS_HOST=redis:6379 --network book-store --name auth-srv -d auth-srv

docker_stop:
	docker stop auth-srv
