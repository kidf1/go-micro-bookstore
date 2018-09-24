package repository

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	proto "github.com/ahmadnurus/go-micro-bookstore/author/srv/proto/author"
)

// Repository struct
type Repository struct {
	Session *mgo.Session
}

// Author interface
type Author interface {
	CreateAuthor(author *proto.AuthorProfile) error
	GetAuthor(username string) (*proto.AuthorProfile, error)
	GetAuthorByID(id string) (*proto.AuthorProfile, error)
	UpdateAuthor(id string, author *proto.AuthorProfile) error
	DeleteAuthor(id string) error
	AddAuthorBook(author string, bookid string) error
	GetAuthorBook(author string) ([]string, error)
	RemoveAuthorBook(author string, bookid string) error
	Close()
}

// Collection method for define database and collection
func (c *Repository) Collection() *mgo.Collection {
	return c.Session.DB(database).C(collection)
}

// CreateAuthor method
func (c *Repository) CreateAuthor(author *proto.AuthorProfile) error {
	author.Id = bson.NewObjectId().Hex()
	return c.Collection().Insert(author)
}

// GetAuthor method
func (c *Repository) GetAuthor(username string) (*proto.AuthorProfile, error) {
	var result *proto.AuthorProfile
	if err := c.Collection().Find(bson.M{"username": username}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAuthorByID method
func (c *Repository) GetAuthorByID(id string) (*proto.AuthorProfile, error) {
	var result *proto.AuthorProfile
	if err := c.Collection().Find(bson.M{"id": id}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAuthor method
func (c *Repository) UpdateAuthor(id string, author *proto.AuthorProfile) error {
	author.UpdatedAt = time.Now().Unix()
	update := bson.M{
		"username":  author.Username,
		"fullname":  author.Fullname,
		"updatedat": author.UpdatedAt,
	}
	return c.Collection().Update(bson.M{"id": id}, bson.M{"$set": update})
}

// DeleteAuthor method
func (c *Repository) DeleteAuthor(id string) error {
	return c.Collection().Remove(bson.M{"id": id})
}

// AddAuthorBook method
func (c *Repository) AddAuthorBook(author string, bookid string) error {
	add := bson.M{"books": bookid}
	return c.Collection().Update(bson.M{"username": author}, bson.M{"$push": add})
}

// GetAuthorBook method
func (c *Repository) GetAuthorBook(author string) ([]string, error) {
	var results []string
	var result *proto.AuthorProfile
	if err := c.Collection().Find(bson.M{"username": author}).One(&result); err != nil {
		return nil, err
	}
	for _, v := range result.Books {
		results = append(results, v)
	}
	return results, nil
}

// RemoveAuthorBook method
func (c *Repository) RemoveAuthorBook(author string, bookid string) error {
	remove := bson.M{"books": bookid}
	return c.Collection().Update(bson.M{"username": author}, bson.M{"$pull": remove})
}

// Close mongo session
func (c *Repository) Close() {
	c.Session.Close()
}
