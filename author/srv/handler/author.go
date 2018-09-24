package handler

import (
	"context"

	"gopkg.in/mgo.v2"

	"github.com/micro/go-micro/errors"

	product "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"

	proto "github.com/ahmadnurus/go-micro-bookstore/author/srv/proto/author"
	repo "github.com/ahmadnurus/go-micro-bookstore/author/srv/repository"
)

// Handler struct
type Handler struct {
	Session        *mgo.Session
	ProductService product.ProductService
}

// Repo method
func (h *Handler) Repo() repo.Author {
	return &repo.Repository{Session: h.Session.Clone()}
}

// CreateAuthor method
func (h *Handler) CreateAuthor(ctx context.Context, req *proto.CreateAuthorRequest, res *proto.CreateAuthorResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().CreateAuthor(req.Author); err != nil {
		return err
	}

	res.Code = 201
	res.Detail = "created"
	res.Author = req.Author
	return nil
}

// GetAuthor method
func (h *Handler) GetAuthor(ctx context.Context, req *proto.GetAuthorRequest, res *proto.GetAuthorResponse) error {
	defer h.Repo().Close()

	author, err := h.Repo().GetAuthor(req.Username)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Author = author
	return nil
}

// GetAuthorByID method
func (h *Handler) GetAuthorByID(ctx context.Context, req *proto.GetAuthorByIDRequest, res *proto.GetAuthorByIDResponse) error {
	defer h.Repo().Close()

	author, err := h.Repo().GetAuthorByID(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Author = author
	return nil
}

// UpdateAuthor method
func (h *Handler) UpdateAuthor(ctx context.Context, req *proto.UpdateAuthorRequest, res *proto.UpdateAuthorResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().UpdateAuthor(req.Author.Id, req.Author); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "updated"
	res.Author = req.Author
	return nil
}

// DeleteAuthor method
func (h *Handler) DeleteAuthor(ctx context.Context, req *proto.DeleteAuthorRequest, res *proto.DeleteAuthorResponse) error {
	defer h.Repo().Close()

	err := h.Repo().DeleteAuthor(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "deleted"
	return nil
}

// AddAuthorBook method
func (h *Handler) AddAuthorBook(ctx context.Context, req *proto.AddAuthorBookRequest, res *proto.AddAuthorBookResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().AddAuthorBook(req.Author, req.BookId); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "added"
	res.Author = req.Author
	res.BookId = req.BookId
	return nil
}

// GetAuthorBook method
func (h *Handler) GetAuthorBook(ctx context.Context, req *proto.GetAuthorBookRequest, res *proto.GetAuthorBookResponse) error {
	defer h.Repo().Close()

	bookIDs, err := h.Repo().GetAuthorBook(req.Author)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	books, err := h.ProductService.GetMultipleBook(context.Background(), &product.GetMultipleBookRequest{
		Id: bookIDs,
	})
	if err != nil {
		return err
	}

	res.Code = 200
	res.Author = req.Author
	res.Books = books.Books
	return nil
}

// RemoveAuthorBook method
func (h *Handler) RemoveAuthorBook(ctx context.Context, req *proto.RemoveAuthorBookRequest, res *proto.RemoveAuthorBookResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().RemoveAuthorBook(req.Author, req.BookId); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.author", "author not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "removed"
	return nil
}
