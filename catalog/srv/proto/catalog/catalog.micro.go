// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/catalog/catalog.proto

/*
Package go_micro_bookstore_srv_catalog is a generated protocol buffer package.

It is generated from these files:
	proto/catalog/catalog.proto

It has these top-level messages:
	Category
	CreateCategoryRequest
	CreateCategoryResponse
	GetCategoryRequest
	GetCategoryResponse
	ListCategoryRequest
	ListCategoryResponse
	UpdateCategoryRequest
	UpdateCategoryResponse
	DeleteCategoryRequest
	DeleteCategoryResponse
	AddBookRequest
	AddBookResponse
	GetByBookRequest
	GetByBookResponse
	ListBookRequest
	ListBookResponse
	RemoveBookRequest
	RemoveBookResponse
*/
package go_micro_bookstore_srv_catalog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"

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

// Client API for Catalog service

type CatalogService interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...client.CallOption) (*GetCategoryResponse, error)
	ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...client.CallOption) (*ListCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error)
	AddBook(ctx context.Context, in *AddBookRequest, opts ...client.CallOption) (*AddBookResponse, error)
	GetByBook(ctx context.Context, in *GetByBookRequest, opts ...client.CallOption) (*GetByBookResponse, error)
	ListBook(ctx context.Context, in *ListBookRequest, opts ...client.CallOption) (*ListBookResponse, error)
	RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...client.CallOption) (*RemoveBookResponse, error)
}

type catalogService struct {
	c    client.Client
	name string
}

func NewCatalogService(name string, c client.Client) CatalogService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.bookstore.srv.catalog"
	}
	return &catalogService{
		c:    c,
		name: name,
	}
}

func (c *catalogService) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...client.CallOption) (*CreateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.CreateCategory", in)
	out := new(CreateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...client.CallOption) (*GetCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetCategory", in)
	out := new(GetCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...client.CallOption) (*ListCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.ListCategory", in)
	out := new(ListCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...client.CallOption) (*UpdateCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.UpdateCategory", in)
	out := new(UpdateCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...client.CallOption) (*DeleteCategoryResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.DeleteCategory", in)
	out := new(DeleteCategoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) AddBook(ctx context.Context, in *AddBookRequest, opts ...client.CallOption) (*AddBookResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.AddBook", in)
	out := new(AddBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) GetByBook(ctx context.Context, in *GetByBookRequest, opts ...client.CallOption) (*GetByBookResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.GetByBook", in)
	out := new(GetByBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) ListBook(ctx context.Context, in *ListBookRequest, opts ...client.CallOption) (*ListBookResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.ListBook", in)
	out := new(ListBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogService) RemoveBook(ctx context.Context, in *RemoveBookRequest, opts ...client.CallOption) (*RemoveBookResponse, error) {
	req := c.c.NewRequest(c.name, "Catalog.RemoveBook", in)
	out := new(RemoveBookResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Catalog service

type CatalogHandler interface {
	CreateCategory(context.Context, *CreateCategoryRequest, *CreateCategoryResponse) error
	GetCategory(context.Context, *GetCategoryRequest, *GetCategoryResponse) error
	ListCategory(context.Context, *ListCategoryRequest, *ListCategoryResponse) error
	UpdateCategory(context.Context, *UpdateCategoryRequest, *UpdateCategoryResponse) error
	DeleteCategory(context.Context, *DeleteCategoryRequest, *DeleteCategoryResponse) error
	AddBook(context.Context, *AddBookRequest, *AddBookResponse) error
	GetByBook(context.Context, *GetByBookRequest, *GetByBookResponse) error
	ListBook(context.Context, *ListBookRequest, *ListBookResponse) error
	RemoveBook(context.Context, *RemoveBookRequest, *RemoveBookResponse) error
}

func RegisterCatalogHandler(s server.Server, hdlr CatalogHandler, opts ...server.HandlerOption) error {
	type catalog interface {
		CreateCategory(ctx context.Context, in *CreateCategoryRequest, out *CreateCategoryResponse) error
		GetCategory(ctx context.Context, in *GetCategoryRequest, out *GetCategoryResponse) error
		ListCategory(ctx context.Context, in *ListCategoryRequest, out *ListCategoryResponse) error
		UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, out *UpdateCategoryResponse) error
		DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error
		AddBook(ctx context.Context, in *AddBookRequest, out *AddBookResponse) error
		GetByBook(ctx context.Context, in *GetByBookRequest, out *GetByBookResponse) error
		ListBook(ctx context.Context, in *ListBookRequest, out *ListBookResponse) error
		RemoveBook(ctx context.Context, in *RemoveBookRequest, out *RemoveBookResponse) error
	}
	type Catalog struct {
		catalog
	}
	h := &catalogHandler{hdlr}
	return s.Handle(s.NewHandler(&Catalog{h}, opts...))
}

type catalogHandler struct {
	CatalogHandler
}

func (h *catalogHandler) CreateCategory(ctx context.Context, in *CreateCategoryRequest, out *CreateCategoryResponse) error {
	return h.CatalogHandler.CreateCategory(ctx, in, out)
}

func (h *catalogHandler) GetCategory(ctx context.Context, in *GetCategoryRequest, out *GetCategoryResponse) error {
	return h.CatalogHandler.GetCategory(ctx, in, out)
}

func (h *catalogHandler) ListCategory(ctx context.Context, in *ListCategoryRequest, out *ListCategoryResponse) error {
	return h.CatalogHandler.ListCategory(ctx, in, out)
}

func (h *catalogHandler) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, out *UpdateCategoryResponse) error {
	return h.CatalogHandler.UpdateCategory(ctx, in, out)
}

func (h *catalogHandler) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, out *DeleteCategoryResponse) error {
	return h.CatalogHandler.DeleteCategory(ctx, in, out)
}

func (h *catalogHandler) AddBook(ctx context.Context, in *AddBookRequest, out *AddBookResponse) error {
	return h.CatalogHandler.AddBook(ctx, in, out)
}

func (h *catalogHandler) GetByBook(ctx context.Context, in *GetByBookRequest, out *GetByBookResponse) error {
	return h.CatalogHandler.GetByBook(ctx, in, out)
}

func (h *catalogHandler) ListBook(ctx context.Context, in *ListBookRequest, out *ListBookResponse) error {
	return h.CatalogHandler.ListBook(ctx, in, out)
}

func (h *catalogHandler) RemoveBook(ctx context.Context, in *RemoveBookRequest, out *RemoveBookResponse) error {
	return h.CatalogHandler.RemoveBook(ctx, in, out)
}