// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package go_micro_bookstore_srv_auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	Token
	AuthorizeRequest
	AuthorizeResponse
	TokenRequest
	TokenResponse
	RevokeRequest
	RevokeResponse
*/
package go_micro_bookstore_srv_auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error)
	Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error)
	Revoke(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.bookstore.srv.auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*AuthorizeResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Authorize", in)
	out := new(AuthorizeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Token", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Revoke(ctx context.Context, in *RevokeRequest, opts ...client.CallOption) (*RevokeResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Revoke", in)
	out := new(RevokeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Authorize(context.Context, *AuthorizeRequest, *AuthorizeResponse) error
	Token(context.Context, *TokenRequest, *TokenResponse) error
	Revoke(context.Context, *RevokeRequest, *RevokeResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Authorize(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error
		Token(ctx context.Context, in *TokenRequest, out *TokenResponse) error
		Revoke(ctx context.Context, in *RevokeRequest, out *RevokeResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Authorize(ctx context.Context, in *AuthorizeRequest, out *AuthorizeResponse) error {
	return h.AuthHandler.Authorize(ctx, in, out)
}

func (h *authHandler) Token(ctx context.Context, in *TokenRequest, out *TokenResponse) error {
	return h.AuthHandler.Token(ctx, in, out)
}

func (h *authHandler) Revoke(ctx context.Context, in *RevokeRequest, out *RevokeResponse) error {
	return h.AuthHandler.Revoke(ctx, in, out)
}
