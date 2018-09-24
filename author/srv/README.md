# Author Service

This is the Author service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/author/srv --namespace=go.micro.bookstore --alias=author --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.author
- Type: srv
- Alias: author

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
./author-srv
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

Create Author

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.CreateAuthor",
	"request": {
	    "author": {
	        "username": "ahmadnurus",
	        "fullname": "Ahmad Nurus S."
	    }
	}
}'
```

Get Author

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.GetAuthor",
	"request": {
		"username": "ahmadnurus"
	}
}'
```

Get Author By ID

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.GetAuthorByID",
	"request": {
		"id": "5ba92defc493f603b0863e21"
	}
}'
```

Update Author

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.UpdateAuthor",
	"request": {
		"author": {
			"id": "5ba85cc5c493f68000e4f8fc",
			"username": "prksu",
			"fullname": "Ahmad Nurus"
		}
	}
}'
```

Delete Author

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.DeleteAuthor",
	"request": {
		"id": "5ba85cc5c493f68000e4f8fc"
	}
}'
```

Add Author Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.AddAuthorBook",
	"request": {
		"author": "ahmadnurus",
		"book_id": "5ba85e41c493f680c53ac0c1"
	}
}'
```

Get Author Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.GetAuthorBook",
	"request": {
		"author": "ahmadnurus"
	}
}'
```

Remove Author Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.author",
	"method": "Author.RemoveAuthorBook",
	"request": {
		"author": "ahmadnurus",
		"book_id": "5ba85e41c493f680c53ac0c1"
	}
}'
```