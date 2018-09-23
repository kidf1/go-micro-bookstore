package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"

	proto "github.com/ahmadnurus/go-micro-bookstore/auth/srv/proto/auth"
)

// Key JWT Secret
var (
	Key = []byte("secret")
)

// Handler struct
type Handler struct {
	Redis *redis.Pool
}

// Authorize method
func (h *Handler) Authorize(ctx context.Context, req *proto.AuthorizeRequest, res *proto.AuthorizeResponse) error {
	conn := h.Redis.Get()
	defer conn.Close()

	if len(req.AccessToken) == 0 {
		return errors.BadRequest("go.micro.bookstore.srv.auth", "require access_token")
	}

	var claims jwt.StandardClaims
	token, err := jwt.ParseWithClaims(req.AccessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if !token.Valid {
		return errors.Unauthorized("go.micro.bookstore.srv.auth", "invalid access_token")
	} else if claims.ExpiresAt < time.Now().Unix() {
		return errors.Unauthorized("go.micro.bookstore.srv.auth", "expired access_token")
	}

	var result []byte
	result, err = redis.Bytes(conn.Do("GET", claims.Subject))
	if err != nil {
		if err == redis.ErrNil {
			return errors.Unauthorized("go.micro.bookstore.srv.auth", "expired access_token")
		}
		return err
	}
	if string(result[:]) != claims.Id {
		return errors.Unauthorized("go.micro.bookstore.srv.auth", "invalid access_token")
	}

	res.Code = 200
	res.Detail = "authorized"
	res.ClientId = claims.Subject
	return nil
}

// Token method
func (h *Handler) Token(ctx context.Context, req *proto.TokenRequest, res *proto.TokenResponse) error {
	conn := h.Redis.Get()
	defer conn.Close()

	jti := uuid.New().String()
	ttl := time.Hour * 2
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Id:        jti,
		Subject:   req.ClientId,
		Issuer:    req.Issuer,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(ttl).Unix(),
	})

	accessToken, err := claims.SignedString(Key)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETEX", req.ClientId, ttl.Seconds(), jti)
	if err != nil {
		return err
	}

	token := &proto.Token{
		AccessToken: accessToken,
		TokenType:   "bearer",
		ExpiresAt:   time.Now().Add(ttl).Unix(),
	}

	res.Code = 200
	res.Token = token
	return nil
}

// Revoke method
func (h *Handler) Revoke(ctx context.Context, req *proto.RevokeRequest, res *proto.RevokeResponse) error {
	conn := h.Redis.Get()
	defer conn.Close()

	if len(req.AccessToken) == 0 {
		return errors.BadRequest("go.micro.bookstore.srv.auth", "require access_token")
	}

	var claims jwt.StandardClaims
	token, err := jwt.ParseWithClaims(req.AccessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if !token.Valid {
		return errors.Unauthorized("go.micro.bookstore.srv.auth", "invalid access_token")

	}

	_, err = conn.Do("DEL", claims.Subject)
	if err != nil {
		return err
	}

	res.Code = 200
	res.Detail = "token revoked"
	return nil
}
