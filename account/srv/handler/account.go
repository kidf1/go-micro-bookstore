package handler

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"

	"github.com/micro/go-micro/errors"

	user "github.com/ahmadnurus/go-micro-bookstore/user/srv/proto/user"

	proto "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"
	repo "github.com/ahmadnurus/go-micro-bookstore/account/srv/repository"
)

// Handler struct
type Handler struct {
	Session     *mgo.Session
	UserService user.UserService
}

// Repo method
func (h *Handler) Repo() repo.Account {
	return &repo.Repository{Session: h.Session.Clone()}
}

// CreateAccount method
func (h *Handler) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest, res *proto.CreateAccountResponse) error {
	defer h.Repo().Close()

	profile, err := h.UserService.CreateProfile(context.Background(), &user.CreateProfileRequest{
		Profile: &user.Profile{
			Username: req.Account.Username,
			Email:    req.Account.Email,
		},
	})
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Account.Username = strings.Replace(strings.ToLower(req.Account.Username), " ", "", -1)
	req.Account.Email = strings.Replace(strings.ToLower(req.Account.Email), " ", "", -1)
	req.Account.Password = string(password[:])
	req.Account.IsVerified = false
	req.Account.ProfileId = profile.Profile.Id
	if err := h.Repo().CreateAccount(req.Account); err != nil {
		return err
	}

	res.Code = 201
	res.Detail = "created"
	res.Account = req.Account
	return nil
}

// GetAccount method
func (h *Handler) GetAccount(ctx context.Context, req *proto.GetAccountRequest, res *proto.GetAccountResponse) error {
	defer h.Repo().Close()

	account, err := h.Repo().GetAccount(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.account", "account not found")
		}
		return err
	}

	res.Code = 200
	res.Account = account
	return nil
}

// UpdateAccount method
func (h *Handler) UpdateAccount(ctx context.Context, req *proto.UpdateAccountRequest, res *proto.UpdateAccountResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().UpdateAccount(req.Account.Id, req.Account); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.account", "account not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "updated"
	res.Account = req.Account
	return nil
}

// DeleteAccount method
func (h *Handler) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest, res *proto.DeleteAccountResponse) error {
	defer h.Repo().Close()

	account, err := h.Repo().GetAccount(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.account", "account not found")
		}
		return err
	}

	if err := h.Repo().DeleteAccount(req.Id); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.account", "account not found")
		}
		return err
	}

	_, err = h.UserService.DeleteProfile(context.Background(), &user.DeleteProfileRequest{
		Id: account.ProfileId,
	})
	if err != nil {
		return err
	}

	res.Code = 200
	res.Detail = "deleted"
	return nil
}

// Login method
func (h *Handler) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {
	defer h.Repo().Close()

	account, err := h.Repo().GetAccountByEmail(req.Email)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.account", "account not found")
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.Unauthorized("go.micro.bookstore.srv.account", "invalid password")
		}
		return err
	}

	res.Code = 200
	res.Detail = "authenticated"
	return nil
}

// Logout method
func (h *Handler) Logout(ctx context.Context, req *proto.LogoutRequest, res *proto.LogoutResponse) error {
	res.Code = 200
	res.Detail = "session ended"
	return nil
}
