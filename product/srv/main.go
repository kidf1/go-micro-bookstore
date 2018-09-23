package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/product/srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	proto "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"
	repo "github.com/ahmadnurus/go-micro-bookstore/product/srv/repository"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.product"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:   "mongo_host",
				EnvVar: "MONGO_HOST",
			},
		),
		micro.Action(func(c *cli.Context) {
			if len(c.String("mongo_host")) > 0 {
				repo.Host = c.String("mongo_host")
			}
		}),
	)

	// Initialise service
	service.Init()

	// Call CreateSession function
	session, err := repo.CreateSession()

	defer session.Close()

	if err != nil {
		log.Fatalf("Could not connect to database %s - %v", repo.Host, err)
	}

	// Register Product Handler
	proto.RegisterProductHandler(service.Server(), &handler.Handler{Session: session})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
