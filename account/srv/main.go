package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/account/srv/handler"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	user "github.com/ahmadnurus/go-micro-bookstore/user/srv/proto/user"

	proto "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"
	repo "github.com/ahmadnurus/go-micro-bookstore/account/srv/repository"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.account"),
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

	// NewUserService
	userService := user.NewUserService("go.micro.bookstore.srv.user", service.Client())

	// Register Account Handler
	proto.RegisterAccountHandler(service.Server(), &handler.Handler{Session: session, UserService: userService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
