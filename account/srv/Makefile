
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/account/account.proto

run: proto

	bash -c "trap 'go run *.go --registry_address=127.0.0.1:8500' EXIT"

build: proto

	go build -o account-srv main.go plugin.go

test:
	go test -v ./... -cover

clean:
	rm -f account-srv

docker:

	CGO_ENABLED=0 GOOS=linux go build -o account-srv -a -installsuffix cgo .

	docker build . -t account-srv:latest

	rm -f account-srv

docker_run:
	docker run --rm -e MICRO_REGISTRY_ADDRESS=consul:8500 -e MONGO_HOST=mongo:27017 --network book-store --name account-srv -d account-srv

docker_stop:
	docker stop account-srv
