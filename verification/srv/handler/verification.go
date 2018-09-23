package handler

import (
	"context"
	"encoding/base64"

	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"

	"github.com/micro/go-micro/errors"

	account "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"

	proto "github.com/ahmadnurus/go-micro-bookstore/verification/srv/proto/verification"
)

// Handler struct
type Handler struct {
	Redis          *redis.Pool
	AccountService account.AccountService
}

// VerifyAccount method
func (h *Handler) VerifyAccount(ctx context.Context, req *proto.VerifyAccountRequest, res *proto.VerifyAccountResponse) error {
	// Get credential from account servcie
	credential, err := h.AccountService.GetAccount(context.Background(), &account.GetAccountRequest{
		Id: req.AccountId,
	})
	if err != nil {
		return err
	}
	if credential.Account.IsVerified == true {
		return errors.BadRequest("go.micro.bookstore.srv.verification", "account already verified")
	}

	// Decode token from base64 to []byte bcrypt
	token, _ := base64.StdEncoding.DecodeString(req.Token)
	if err != nil {
		return err
	}

	// Compare is token valid
	if err := bcrypt.CompareHashAndPassword(token, []byte(credential.Account.Email)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.BadRequest("go.micro.bookstore.srv.verification", "invalid token")
		}
		return err
	}

	conn := h.Redis.Get()
	defer conn.Close()

	// Get existing token in redis
	var result []byte
	result, err = redis.Bytes(conn.Do("GET", credential.Account.Email))

	if err != nil {
		// If token has expired
		if err == redis.ErrNil {
			return errors.NotFound("go.micro.bookstore.srv.verification", "token has expired")
		}
		return err
	}

	// If the sent token is not the same as the existing token.
	// As Verifying if the user has requested the verification code again
	// but request with old token
	if string(result[:]) != string(token[:]) {
		return errors.BadRequest("go.micro.bookstore.srv.verification", "invalid token")
	}

	// Delete token in redis if token is valid
	_, err = conn.Do("DEL", credential.Account.Email)
	if err != nil {
		return err
	}

	// Update account
	_, err = h.AccountService.UpdateAccount(context.Background(), &account.UpdateAccountRequest{
		Account: &account.Credential{
			Id:         req.AccountId,
			IsVerified: true,
			AuthorId:   credential.Account.AuthorId,
			LastLogin:  credential.Account.LastLogin,
		},
	})
	if err != nil {
		return err
	}

	res.Code = 200
	res.Detail = "verified"
	return nil
}
