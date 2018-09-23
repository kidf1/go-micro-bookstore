# Auth Service

This is the Auth service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/auth/srv --namespace=go.micro.bookstore --alias=auth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.auth
- Type: srv
- Alias: auth

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
./auth-srv
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

Authorize

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.auth",
	"method": "Auth.Authorize",
	"request": {
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc3MjQ2MjgsImp0aSI6IjIxYTgxODhmLTUzM2QtNDlhMS1iMjUyLWM5ZmNjYmUwYTBmNCIsImlhdCI6MTUzNzcxNzQyOCwiaXNzIjoiZ28ubWljcm8uYm9va3N0b3JlLnNydi5hY2NvdW50Iiwic3ViIjoiNWJhN2FiODNjNDkzZjY1NWNhNDdhODgwIn0.3VvvDjXq-o2RCfkzBatlsAuglc1mGUaOQcgaWi-rJPI"
	}
}'
```

Token

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.auth",
	"method": "Auth.Token",
	"request": {
		"client_id": "12345",
		"issuer": "go.micro.bookstore.srv.auth"
	}
}'
```

Revoke

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.auth",
	"method": "Auth.Revoke",
	"request": {
		"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc3MjQ2MjgsImp0aSI6IjIxYTgxODhmLTUzM2QtNDlhMS1iMjUyLWM5ZmNjYmUwYTBmNCIsImlhdCI6MTUzNzcxNzQyOCwiaXNzIjoiZ28ubWljcm8uYm9va3N0b3JlLnNydi5hY2NvdW50Iiwic3ViIjoiNWJhN2FiODNjNDkzZjY1NWNhNDdhODgwIn0.3VvvDjXq-o2RCfkzBatlsAuglc1mGUaOQcgaWi-rJPI"
	}
}'
```