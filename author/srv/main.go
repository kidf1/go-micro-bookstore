package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/author/srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	product "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"

	proto "github.com/ahmadnurus/go-micro-bookstore/author/srv/proto/author"
	repo "github.com/ahmadnurus/go-micro-bookstore/author/srv/repository"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.author"),
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

	// NewProductService
	productService := product.NewProductService("go.micro.bookstore.srv.product", service.Client())

	// Register Handler
	proto.RegisterAuthorHandler(service.Server(), &handler.Handler{Session: session, ProductService: productService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
