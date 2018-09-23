package repository

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	proto "github.com/ahmadnurus/go-micro-bookstore/user/srv/proto/user"
)

// Repository struct
type Repository struct {
	Session *mgo.Session
}

// User interface
type User interface {
	CreateProfile(profile *proto.Profile) error
	GetProfile(username string) (*proto.Profile, error)
	UpdateProfile(id string, profile *proto.Profile) error
	DeleteProfile(id string) error
	SearchProfile(name string) ([]*proto.Profile, error)
	Close()
}

// Collection method for define database and collection
func (c *Repository) Collection() *mgo.Collection {
	return c.Session.DB(database).C(collection)
}

// CreateProfile method
func (c *Repository) CreateProfile(profile *proto.Profile) error {
	profile.Id = bson.NewObjectId().Hex()
	profile.CreatedAt = time.Now().Unix()
	profile.UpdatedAt = time.Now().Unix()
	return c.Collection().Insert(profile)
}

// GetProfile method
func (c *Repository) GetProfile(username string) (*proto.Profile, error) {
	var result *proto.Profile
	if err := c.Collection().Find(bson.M{"username": username}).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateProfile method
func (c *Repository) UpdateProfile(id string, profile *proto.Profile) error {
	profile.UpdatedAt = time.Now().Unix()
	update := bson.M{
		"fullname":  profile.Fullname,
		"image":     profile.Image,
		"updatedat": profile.UpdatedAt,
	}
	return c.Collection().Update(bson.M{"id": id}, bson.M{"$set": update})
}

// DeleteProfile method
func (c *Repository) DeleteProfile(id string) error {
	return c.Collection().Remove(bson.M{"id": id})
}

// SearchProfile method
func (c *Repository) SearchProfile(name string) ([]*proto.Profile, error) {
	var results []*proto.Profile
	if err := c.Collection().Find(bson.M{"$text": bson.M{"$search": name}}).All(&results); err != nil {
		return nil, err
	}
	return results, nil
}

//Close mongo session
func (c *Repository) Close() {
	c.Session.Close()
}
