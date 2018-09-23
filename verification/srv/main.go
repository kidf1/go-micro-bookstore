package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/verification/srv/handler"
	"github.com/ahmadnurus/go-micro-bookstore/verification/srv/redis"
	"github.com/ahmadnurus/go-micro-bookstore/verification/srv/subscriber"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	account "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"

	proto "github.com/ahmadnurus/go-micro-bookstore/verification/srv/proto/verification"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.verification"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:   "redis_host",
				EnvVar: "REDIS_HOST",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("redis_host")) > 0 {
				redis.Host = c.String("redis_host")
			}
		}),
	)

	// Initialise service
	service.Init()

	// Call redis NewPool function
	pool := redis.NewPool()
	defer pool.Close()

	// NewAccountService
	accountService := account.NewAccountService("go.micro.bookstore.srv.account", service.Client())

	// Register Verification Handler
	proto.RegisterVerificationHandler(service.Server(), &handler.Handler{Redis: pool, AccountService: accountService})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.bookstore.srv.verification:account_created", service.Server(), &subscriber.AccountCreated{Redis: pool})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
