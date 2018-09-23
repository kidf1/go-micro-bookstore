# Catalog Service

This is the Catalog service

Generated with

```
micro new github.com/ahmadnurus/go-micro-bookstore/catalog/srv --namespace=go.micro.bookstore --alias=catalog --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Usage](#usage)
- [Example Call Service](#example-call-service)

## Configuration

- FQDN: go.micro.bookstore.srv.catalog
- Type: srv
- Alias: catalog

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
./catalog-srv
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

Create Category

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.CreateCategory",
	"request": {
	    "category": {
	        "name": "Programming"
	    }
	}
}'
```

Get Category

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.GetCategory",
	"request": {
		"name": "programming"
	}
}'
```

List Category

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.ListCategory",
	"request": {}
}'
```

Update Category

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.UpdateCategory",
	"request": {
	    "category": {
	        "id": "5ba76c7cc493f63114a56b48",
	        "name": "Programming Go"
	    }
	}
}'
```

Delete Category

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.DeleteCategory",
	"request": {
		"id": "5ba76c7cc493f63114a56b48"
	}
}'
```

AddBook

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.AddBook",
	"request": {
		"category": "programming-go",
		"book_id": "5ba75de8c493f626ce20773b"
	}
}'
```

GetByBook

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.GetByBook",
	"request": {
		"book_id": "5b9de9d2c493f6a5bdb8f951"
	}
}'
```

ListBook

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.ListBook",
	"request": {
		"category": "programming-go"
	}
}'
```

RemoveBook

```
curl -X POST \
  http://127.0.0.1:8080/rpc \
  -H 'Content-Type: application/json' \
  -d '{
	"service": "go.micro.bookstore.srv.catalog",
	"method": "Catalog.RemoveBook",
	"request": {
		"category": "programming-go",
		"book_id": "5ba75de8c493f626ce20773b"
	}
}'
```