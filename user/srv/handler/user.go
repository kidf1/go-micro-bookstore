package handler

import (
	"context"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/micro/go-micro/errors"

	proto "github.com/ahmadnurus/go-micro-bookstore/user/srv/proto/user"
	repo "github.com/ahmadnurus/go-micro-bookstore/user/srv/repository"
)

// Handler struct
type Handler struct {
	Session *mgo.Session
}

// Repo method
func (h *Handler) Repo() repo.User {
	return &repo.Repository{Session: h.Session.Clone()}
}

// CreateProfile method
func (h *Handler) CreateProfile(ctx context.Context, req *proto.CreateProfileRequest, res *proto.CreateProfileResponse) error {
	defer h.Repo().Close()

	req.Profile.Username = strings.Replace(strings.ToLower(req.Profile.Username), " ", "", -1)
	req.Profile.Email = strings.Replace(strings.ToLower(req.Profile.Email), " ", "", -1)
	if err := h.Repo().CreateProfile(req.Profile); err != nil {
		return err
	}

	res.Code = 201
	res.Detail = "created"
	res.Profile = req.Profile
	return nil
}

// GetProfile method
func (h *Handler) GetProfile(ctx context.Context, req *proto.GetProfileRequest, res *proto.GetProfileResponse) error {
	defer h.Repo().Close()

	profile, err := h.Repo().GetProfile(req.Username)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.user", "user not found")
		}
		return err
	}

	res.Code = 200
	res.Profile = profile
	return nil
}

// UpdateProfile method
func (h *Handler) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest, res *proto.UpdateProfileResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().UpdateProfile(req.Profile.Id, req.Profile); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.user", "user not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "updated"
	res.Profile = req.Profile
	return nil
}

// DeleteProfile method
func (h *Handler) DeleteProfile(ctx context.Context, req *proto.DeleteProfileRequest, res *proto.DeleteProfileResponse) error {
	defer h.Repo().Close()

	err := h.Repo().DeleteProfile(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.user", "user not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "deleted"
	return nil
}

// SearchProfile method
func (h *Handler) SearchProfile(ctx context.Context, req *proto.SearchProfileRequest, res *proto.SearchProfileResponse) error {
	defer h.Repo().Close()

	profile, err := h.Repo().SearchProfile(req.Name)
	if len(profile) == 0 {
		return errors.NotFound("go.micro.bookstore.srv.user", "user not found")
	}
	if err != nil {
		return err
	}

	res.Code = 200
	res.Profile = profile
	return nil
}
