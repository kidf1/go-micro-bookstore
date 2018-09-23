# Book Store Microservices

This repository contains my learning project how to implement microservices architecture.

## Microservices

### API

### Service

## How To Run

### Under development

#### Requirement

* [Golang](https://golang.org/dl/)
* [Micro](https://github.com/micro/micro)
* [Docker](https://docs.docker.com/install/) & [docker-compose](https://docs.docker.com/compose/install/)

#### Run docker container

```
docker-compose up -d
```

docker-compose contains:

* consul - service discorvery - binding port :8500
* go.micro.api - micro api gateway - binding port :8080

make sure any port above is not being used

#### Run microservices

```
make run
```

in each [microservices](#microservices)