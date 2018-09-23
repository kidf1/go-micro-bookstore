package repository

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	proto "github.com/ahmadnurus/go-micro-bookstore/account/srv/proto/account"
)

// Repository struct
type Repository struct {
	Session *mgo.Session
}

// Account interface
type Account interface {
	CreateAccount(account *proto.Credential) error
	GetAccount(id string) (*proto.Credential, error)
	GetAccountByEmail(email string) (*proto.Credential, error)
	UpdateAccount(id string, account *proto.Credential) error
	DeleteAccount(id string) error
	Close()
}

// Collection method for define database and collection
func (c *Repository) Collection() *mgo.Collection {
	return c.Session.DB(database).C(collection)
}

// CreateAccount method
func (c *Repository) CreateAccount(account *proto.Credential) error {
	account.Id = bson.NewObjectId().Hex()
	account.CreatedAt = time.Now().Unix()
	account.UpdatedAt = time.Now().Unix()
	return c.Collection().Insert(account)
}

// GetAccount method
func (c *Repository) GetAccount(id string) (*proto.Credential, error) {
	var result *proto.Credential
	if err := c.Collection().Find(bson.M{"id": id}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountByEmail method
func (c *Repository) GetAccountByEmail(email string) (*proto.Credential, error) {
	var result *proto.Credential
	if err := c.Collection().Find(bson.M{"email": email}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAccount method
func (c *Repository) UpdateAccount(id string, account *proto.Credential) error {
	account.UpdatedAt = time.Now().Unix()
	update := bson.M{
		"isverified": account.IsVerified,
		"authorid":   account.AuthorId,
		"lastlogin":  account.LastLogin,
		"updatedat":  account.UpdatedAt,
	}
	return c.Collection().Update(bson.M{"id": id}, bson.M{"$set": update})
}

// DeleteAccount method
func (c *Repository) DeleteAccount(id string) error {
	return c.Collection().Remove(bson.M{"id": id})
}

//Close mongo session
func (c *Repository) Close() {
	c.Session.Close()
}
