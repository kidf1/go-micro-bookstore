# User Service

This is the User service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/user/srv --namespace=go.micro.bookstore --alias=user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.user
- Type: srv
- Alias: user

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
./user-srv
```

or

```
make run
```

Run service on docker

```
make docker_run
```

Service container will running on background, to stop

```
make docker_stop
```

## Example Call Service

### Call service via micro-api

Create Profile

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.user",
	"method": "User.CreateProfile",
	"request": {
		"profile": {
			"username": "ahmadnurus",
			"email": "ahmadnurus.sh@gmail.com"
		}
	}
}'
```

Get Profile

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.user",
	"method": "User.GetProfile",
	"request": {
		"username": "ahmadnurus"
	}
}'
```

Update Profile

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.user",
	"method": "User.UpdateProfile",
	"request": {
		"profile": {
			"id": "5ba77beec493f63b0c057802",
			"fullname": "Ahmad Nurus Shobah",
			"image": "profile.png"
		}
	}
}'
```

Delete Profile

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.user",
	"method": "User.DeleteProfile",
	"request": {
		"id": "5ba77beec493f63b0c057802"
	}
}'
```

Search Profile

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.user",
	"method": "User.SearchProfile",
	"request": {
		"name": "nurus"
	}
}'
```