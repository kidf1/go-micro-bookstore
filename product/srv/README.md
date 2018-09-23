# Product Service

This is the Product service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/product/srv --namespace=go.micro.bookstore --alias=product --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.product
- Type: srv
- Alias: product

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
./product-srv
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

Create Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.CreateBook",
	"request": {
	    "book": {
	        "name": "Learning Nodejs",
	        "description": "Learning nodejs from basic",
	        "cover": "cover.png",
	        "author": "Ahmad Nurus",
	        "price": 25.40,
	        "released_at": 1535744105 
	    }
	}
}'
```

Get Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.GetBook",
	"request": {
		"name": "learning-nodejs"
	}
}'
```

Get Multiple Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.GetMultipleBook",
	"request": {
	    "id": ["5ba75de8c493f626ce20773b"]
	}
}'
```

List Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.ListBook",
	"request": {}
}'
```

Update Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.UpdateBook",
	"request": {
	    "book": {
	        "id": "5ba75de8c493f626ce20773b",
	        "name": "Learn Golang",
	        "description": "Learning golang from zero to hero you are superman",
	        "cover": "cover.png",
	        "price": 14.50
	    }
	}
}'
```

Delete Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.DeleteBook",
	"request": {
		"id": "5b9ddc64c493f69a8c557c5c"
	}
}'
```

Search Book

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.product",
	"method": "Product.SearchBook",
	"request": {
	    "name": "Learn"
	}
}'
```