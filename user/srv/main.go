package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/user/srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	proto "github.com/ahmadnurus/go-micro-bookstore/user/srv/proto/user"
	repo "github.com/ahmadnurus/go-micro-bookstore/user/srv/repository"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.user"),
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

	// Register User Handler
	proto.RegisterUserHandler(service.Server(), &handler.Handler{Session: session})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
