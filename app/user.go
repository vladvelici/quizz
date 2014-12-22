package app

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
)

const (
	// Datastore Entity kind for users
	UserKind = "User"
)

type User struct {
	Id          string
	Email       string
	DisplayName string
}

func (user *User) populate(from *user.User) {
	user.Email = from.Email
	user.Id = from.ID
	user.DisplayName = from.Email
}

// LogoutURL
func (user *User) LogoutURL() {
}

// Generate a datastore key for the user with the given string ID.
func Key(c appengine.Context, id string) *datastore.Key {
	return datastore.NewKey(c, UserKind, id, 0, nil)
}

// Fetch a user from the datastore or cancel its premises.
func UserFetchOrCreate(c appengine.Context, u *user.User) (*User, error) {
	k := Key(c, u.ID)
	var user User
	err := datastore.Get(c, k, &user)

	// Add user if does not exist
	if err == datastore.ErrNoSuchEntity {
		user.populate(u)
		_, err = datastore.Put(c, k, &user)
		c.Infof("Added new user %s. %s", user.Id, err)
	}

	if err != nil {
		return nil, err
	}

	// Update e-mail if changed
	if u.Email != user.Email {
		user.Email = u.Email
		_, err = datastore.Put(c, k, &user)
		c.Infof("Email changed for user %s. %s", user.Id, err)
	}

	return &user, err
}
