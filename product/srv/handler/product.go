package handler

import (
	"context"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/micro/go-micro/errors"

	proto "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"
	repo "github.com/ahmadnurus/go-micro-bookstore/product/srv/repository"
)

// Handler struct
type Handler struct {
	Session *mgo.Session
}

// Repo method
func (h *Handler) Repo() repo.Product {
	return &repo.Repository{Session: h.Session.Clone()}
}

// CreateBook method
func (h *Handler) CreateBook(ctx context.Context, req *proto.CreateBookRequest, res *proto.CreateBookResponse) error {
	defer h.Repo().Close()

	req.Book.DisplayName = req.Book.Name
	req.Book.Name = strings.Replace(strings.ToLower(req.Book.Name), " ", "-", -1)
	if err := h.Repo().CreateBook(req.Book); err != nil {
		return err
	}

	res.Code = 201
	res.Detail = "created"
	res.Book = req.Book
	return nil
}

// GetBook method
func (h *Handler) GetBook(ctx context.Context, req *proto.GetBookRequest, res *proto.GetBookResponse) error {
	defer h.Repo().Close()

	book, err := h.Repo().GetBook(req.Name)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.product", "book not found")
		}
		return err
	}

	res.Code = 200
	res.Book = book
	return nil
}

// GetMultipleBook method
func (h *Handler) GetMultipleBook(ctx context.Context, req *proto.GetMultipleBookRequest, res *proto.GetMultipleBookResponse) error {
	defer h.Repo().Close()
	books, err := h.Repo().GetMultipleBook(req.Id)
	if len(books) == 0 {
		return errors.NotFound("go.micro.bookstore.srv.product", "book not found")
	}
	if err != nil {
		return err
	}

	res.Code = 200
	res.Books = books
	return nil
}

// ListBook method
func (h *Handler) ListBook(ctx context.Context, req *proto.ListBookRequest, res *proto.ListBookResponse) error {
	defer h.Repo().Close()

	books, err := h.Repo().ListBook()
	if err != nil {
		return err
	}

	res.Code = 200
	res.Books = books
	return nil
}

// UpdateBook method
func (h *Handler) UpdateBook(ctx context.Context, req *proto.UpdateBookRequest, res *proto.UpdateBookResponse) error {
	defer h.Repo().Close()

	req.Book.DisplayName = req.Book.Name
	req.Book.Name = strings.Replace(strings.ToLower(req.Book.Name), " ", "-", -1)
	if err := h.Repo().UpdateBook(req.Book.Id, req.Book); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.product", "book not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "updated"
	res.Book = req.Book
	return nil
}

// DeleteBook method
func (h *Handler) DeleteBook(ctx context.Context, req *proto.DeleteBookRequest, res *proto.DeleteBookResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().DeleteBook(req.Id); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.product", "book not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "deleted"
	return nil
}

// SearchBook method
func (h *Handler) SearchBook(ctx context.Context, req *proto.SearchBookRequest, res *proto.SearchBookResponse) error {
	defer h.Repo().Close()
	books, err := h.Repo().SearchBook(req.Name)
	if len(books) == 0 {
		return errors.NotFound("go.micro.bookstore.srv.product", "book not found")
	}
	if err != nil {
		return err
	}

	res.Code = 200
	res.Books = books
	return nil
}
