package handler

import (
	"context"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/micro/go-micro/errors"

	product "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"

	proto "github.com/ahmadnurus/go-micro-bookstore/catalog/srv/proto/catalog"
	repo "github.com/ahmadnurus/go-micro-bookstore/catalog/srv/repository"
)

// Handler struct
type Handler struct {
	Session        *mgo.Session
	ProductService product.ProductService
}

// Repo method
func (h *Handler) Repo() repo.Catalog {
	return &repo.Repository{Session: h.Session.Clone()}
}

// CreateCategory method
func (h *Handler) CreateCategory(ctx context.Context, req *proto.CreateCategoryRequest, res *proto.CreateCategoryResponse) error {
	defer h.Repo().Close()

	req.Category.DisplayName = req.Category.Name
	req.Category.Name = strings.Replace(strings.ToLower(req.Category.Name), " ", "-", -1)
	if err := h.Repo().CreateCategory(req.Category); err != nil {
		return err
	}

	res.Code = 201
	res.Detail = "created"
	res.Category = req.Category
	return nil
}

// GetCategory method
func (h *Handler) GetCategory(ctx context.Context, req *proto.GetCategoryRequest, res *proto.GetCategoryResponse) error {
	defer h.Repo().Close()

	category, err := h.Repo().GetCategory(req.Name)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Category = category
	return nil
}

// ListCategory method
func (h *Handler) ListCategory(ctx context.Context, req *proto.ListCategoryRequest, res *proto.ListCategoryResponse) error {
	defer h.Repo().Close()

	categories, err := h.Repo().ListCategory()
	if err != nil {
		return err
	}

	res.Code = 200
	res.Categories = categories
	return nil
}

// UpdateCategory method
func (h *Handler) UpdateCategory(ctx context.Context, req *proto.UpdateCategoryRequest, res *proto.UpdateCategoryResponse) error {
	defer h.Repo().Close()

	req.Category.DisplayName = req.Category.Name
	req.Category.Name = strings.Replace(strings.ToLower(req.Category.Name), " ", "-", -1)
	if err := h.Repo().UpdateCategory(req.Category.Id, req.Category); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "updated"
	res.Category = req.Category
	return nil
}

// DeleteCategory method
func (h *Handler) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest, res *proto.DeleteCategoryResponse) error {
	defer h.Repo().Close()

	err := h.Repo().DeleteCategory(req.Id)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "deleted"
	return nil
}

// AddBook method
func (h *Handler) AddBook(ctx context.Context, req *proto.AddBookRequest, res *proto.AddBookResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().AddBook(req.Category, req.BookId); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "added"
	res.Category = req.Category
	res.BookId = req.BookId
	return nil
}

// GetByBook method
func (h *Handler) GetByBook(ctx context.Context, req *proto.GetByBookRequest, res *proto.GetByBookResponse) error {
	defer h.Repo().Close()

	category, err := h.Repo().GetByBook(req.BookId)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Category = category
	return nil
}

// ListBook method
func (h *Handler) ListBook(ctx context.Context, req *proto.ListBookRequest, res *proto.ListBookResponse) error {
	defer h.Repo().Close()

	bookIDs, err := h.Repo().ListBook(req.Category)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
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
	res.Category = req.Category
	res.Books = books.Books
	return nil
}

// RemoveBook method
func (h *Handler) RemoveBook(ctx context.Context, req *proto.RemoveBookRequest, res *proto.RemoveBookResponse) error {
	defer h.Repo().Close()

	if err := h.Repo().RemoveBook(req.Category, req.BookId); err != nil {
		if err == mgo.ErrNotFound {
			return errors.NotFound("go.micro.bookstore.srv.catalog", "category not found")
		}
		return err
	}

	res.Code = 200
	res.Detail = "removed"
	return nil
}
