package repository

import (
	"time"

	proto "github.com/ahmadnurus/go-micro-bookstore/product/srv/proto/product"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository struct
type Repository struct {
	Session *mgo.Session
}

// Product interface
type Product interface {
	CreateBook(book *proto.Book) error
	GetBook(name string) (*proto.Book, error)
	GetMultipleBook(id []string) ([]*proto.Book, error)
	ListBook() ([]*proto.Book, error)
	UpdateBook(id string, book *proto.Book) error
	DeleteBook(id string) error
	SearchBook(name string) ([]*proto.Book, error)
	Close()
}

// Collection method for define database and collection
func (c *Repository) Collection() *mgo.Collection {
	return c.Session.DB(database).C(collection)
}

// CreateBook method
func (c *Repository) CreateBook(book *proto.Book) error {
	book.Id = bson.NewObjectId().Hex()
	book.CreatedAt = time.Now().Unix()
	book.UpdatedAt = time.Now().Unix()
	return c.Collection().Insert(book)
}

// GetBook method
func (c *Repository) GetBook(name string) (*proto.Book, error) {
	var result *proto.Book
	if err := c.Collection().Find(bson.M{"name": name}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMultipleBook method
func (c *Repository) GetMultipleBook(id []string) ([]*proto.Book, error) {
	var result []*proto.Book
	if err := c.Collection().Find(bson.M{"id": bson.M{"$in": id}}).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListBook method
func (c *Repository) ListBook() ([]*proto.Book, error) {
	var results []*proto.Book
	if err := c.Collection().Find(nil).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateBook method
func (c *Repository) UpdateBook(id string, book *proto.Book) error {
	book.UpdatedAt = time.Now().Unix()
	update := bson.M{
		"name":        book.Name,
		"displayname": book.DisplayName,
		"description": book.Description,
		"cover":       book.Cover,
		"price":       book.Price,
		"updatedat":   book.UpdatedAt,
	}
	return c.Collection().Update(bson.M{"id": id}, bson.M{"$set": update})
}

// DeleteBook method
func (c *Repository) DeleteBook(id string) error {
	return c.Collection().Remove(bson.M{"id": id})
}

// SearchBook method
func (c *Repository) SearchBook(name string) ([]*proto.Book, error) {
	var result []*proto.Book
	if err := c.Collection().Find(bson.M{"$text": bson.M{"$search": name}}).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//Close mongo session
func (c *Repository) Close() {
	c.Session.Close()
}
