# Verification Service

This is the Verification service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/verification/srv --namespace=go.micro.bookstore --alias=verification --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.verification
- Type: srv
- Alias: verification

## Usage

### Build

Build binary

```
make build
```

Build docker image

```
make docker
```

### Run service

Run service locally

```
./verification-srv
```

or

```
make run
```

Run service on docker

```
make docker

make docker_run
```

Service container will running on background, to stop

```
make docker_stop
```

## Example Call Service

### Call service via micro-api

Verify

*account_id and token are displayed in the log this service when account created*

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.verification",
	"method": "Verification.VerifyAccount",
	"request": {
		"account_id": "5ba7a3bbc493f64ce96ca0bf",
		"token": "JDJhJDEwJHBEd2dKRW5CdEdaYWMwLlZHcUlHdWVHVHh0ZjEwS2p6Y2EzYzFiNlZjdkxjYUhScXY3U1Rx"
	}
}'
```