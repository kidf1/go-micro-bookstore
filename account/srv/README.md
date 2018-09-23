# Account Service

This is the Account service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/account/srv --namespace=go.micro.bookstore --alias=account --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.account
- Type: srv
- Alias: account

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
./account-srv
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

### Call service using http via micro-api

Create Account

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.CreateAccount",
	"request": {
	    "account": {
	    	"username": "ahmadnurus",
			"email": "ahmadnurus.sh@gmail.com",
			"password": "secret"
	    }
	}
}'
```

Get Account

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.GetAccount",
	"request": {
		"id": "5ba797dec493f6468a3b353f"
	}
}'
```

Update Account

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.UpdateAccount",
	"request": {
	    "account": {
	        "id": "5ba797dec493f6468a3b353f",
			"is_verified": true,
			"author_id": "",
			"last_login": 0
	    }
	}
}'
```

Delete Account

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.DeleteAccount",
	"request": {
		"id": "5ba797dec493f6468a3b353f"
	}
}'
```

Login

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.Login",
	"request": {
		"email": "ahmadnurus.sh@gmail.com",
		"password": "secret"
	}
}'
```

Logout

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.account",
	"method": "Account.Logout",
	"request": {}
}'
```