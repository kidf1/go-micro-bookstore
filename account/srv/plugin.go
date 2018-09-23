package main

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	auth "github.com/ahmadnurus/go-micro-bookstore/auth/srv/proto/auth"
)

// AuthWrapper to wrapping auth
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.BadRequest("go.micro.bookstore.srv.account", "bad request")
		}
		accessToken := meta["Access-Token"]

		authService := auth.NewAuthService("go.micro.bookstore.srv.auth", client.DefaultClient)
		authResponse, err := authService.Authorize(ctx, &auth.AuthorizeRequest{
			AccessToken: accessToken,
		})
		if err != nil {
			return err
		}

		ctx = metadata.NewContext(context.Background(), map[string]string{
			"Access-Token": accessToken,
			"ClientID":     authResponse.ClientId,
		})
		err = fn(ctx, req, resp)
		return err
	}
}
