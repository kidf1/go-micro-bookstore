package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	proto "github.com/ahmadnurus/go-micro-bookstore/catalog/srv/proto/catalog"
)

// Repository struct
type Repository struct {
	Session *mgo.Session
}

// Catalog interface
type Catalog interface {
	CreateCategory(book *proto.Category) error
	GetCategory(name string) (*proto.Category, error)
	ListCategory() ([]*proto.Category, error)
	UpdateCategory(id string, Category *proto.Category) error
	DeleteCategory(id string) error
	AddBook(category string, bookid string) error
	GetByBook(bookid string) (*proto.Category, error)
	ListBook(category string) ([]string, error)
	RemoveBook(category string, bookid string) error
	Close()
}

// Collection method for define database and collection
func (c *Repository) Collection() *mgo.Collection {
	return c.Session.DB(database).C(collection)
}

// CreateCategory method
func (c *Repository) CreateCategory(category *proto.Category) error {
	category.Id = bson.NewObjectId().Hex()
	return c.Collection().Insert(category)
}

// GetCategory method
func (c *Repository) GetCategory(name string) (*proto.Category, error) {
	var result *proto.Category
	if err := c.Collection().Find(bson.M{"name": name}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListCategory method
func (c *Repository) ListCategory() ([]*proto.Category, error) {
	var result []*proto.Category
	if err := c.Collection().Find(nil).Select(bson.M{"id": 1, "name": 1}).All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCategory method
func (c *Repository) UpdateCategory(id string, category *proto.Category) error {
	upadate := bson.M{
		"name":        category.Name,
		"displayname": category.DisplayName,
	}
	return c.Collection().Update(bson.M{"id": id}, bson.M{"$set": upadate})
}

// DeleteCategory method
func (c *Repository) DeleteCategory(id string) error {
	return c.Collection().Remove(bson.M{"id": id})
}

// AddBook method
func (c *Repository) AddBook(category string, bookid string) error {
	add := bson.M{"books": bookid}
	return c.Collection().Update(bson.M{"name": category}, bson.M{"$push": add})
}

// GetByBook method
func (c *Repository) GetByBook(bookid string) (*proto.Category, error) {
	var result *proto.Category
	if err := c.Collection().Find(bson.M{"books": bookid}).One(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListBook method return array of bookid for spesific category
func (c *Repository) ListBook(category string) ([]string, error) {
	var results []string
	var result *proto.Category
	if err := c.Collection().Find(bson.M{"name": category}).One(&result); err != nil {
		return nil, err
	}
	for _, v := range result.Books {
		results = append(results, v)
	}
	return results, nil
}

// RemoveBook method
func (c *Repository) RemoveBook(category string, bookid string) error {
	remove := bson.M{"books": bookid}
	return c.Collection().Update(bson.M{"name": category}, bson.M{"$pull": remove})
}

//Close mongo session
func (c *Repository) Close() {
	c.Session.Close()
}
