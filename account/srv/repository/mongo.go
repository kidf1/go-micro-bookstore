package repository

import "gopkg.in/mgo.v2"

// Host mongo default
var (
	Host = "127.0.0.1:27017"
)

// Database and Collection
const (
	database   = "book-store"
	collection = "account"
)

// CreateSession mongo
func CreateSession() (*mgo.Session, error) {
	session, err := mgo.Dial(Host)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
