package repository

import "gopkg.in/mgo.v2"

// Host mongo default
var (
	Host = "127.0.0.1:27017"
)

// Database and collection
const (
	database   = "book-store"
	collection = "user"
)

// CreateSession mongo
func CreateSession() (*mgo.Session, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(database).C(collection)
	if err := c.EnsureIndex(mgo.Index{
		Key: []string{"$text:username", "$text:fullname"},
	}); err != nil {
		return nil, err
	}

	return session, nil
}
