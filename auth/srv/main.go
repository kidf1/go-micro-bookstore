package main

import (
	"github.com/ahmadnurus/go-micro-bookstore/auth/srv/handler"
	"github.com/ahmadnurus/go-micro-bookstore/auth/srv/redis"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	proto "github.com/ahmadnurus/go-micro-bookstore/auth/srv/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.bookstore.srv.auth"),
		micro.Version("latest"),
		micro.Flags(
			cli.StringFlag{
				Name:   "jwt_secret",
				EnvVar: "JWT_SECRET",
			},
			cli.StringFlag{
				Name:   "redis_host",
				EnvVar: "REDIS_HOST",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("jwt_secret")) > 0 {
				handler.Key = []byte(c.String("jwt_secret"))
			}
			if len(c.String("redis_host")) > 0 {
				redis.Host = c.String("redis_host")
			}
		}),
	)

	// Initialise service
	service.Init()

	pool := redis.NewPool()
	defer pool.Close()

	// Register Handler
	proto.RegisterAuthHandler(service.Server(), &handler.Handler{Redis: pool})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
