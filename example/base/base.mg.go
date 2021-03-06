// Code generated by mgen.
// source:
// base.yaml
// DO NOT EDIT

package base

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//go:generate goimports -w base.mg.go
var (
	db *DB
)

type DB struct {
	name    string
	session *mgo.Session
}

func NewDB(name string) *DB {
	if db == nil {
		db = &DB{
			name: name,
		}
	}

	return db
}

func (db *DB) InitDB(session *mgo.Session) {
	if session == nil {
		log.Fatalf("[FATAL] you must connect database\n")
	}
	db.session = session

	log.Printf("[INFO] %v connection succeeded\n", db.name)
}

func GetSessionAndCollection(collection string) (*mgo.Session, *mgo.Collection) {
	s := db.session.Copy()
	c := s.DB(db.name).C(collection)

	return s, c
}

func GetSessionAndGridFS(prefix string) (*mgo.Session, *mgo.GridFS) {
	s := db.session.Copy()
	f := s.DB(db.name).GridFS(prefix)
	return s, f
}

const (
	CollectionUser = "users"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	UserName  string        `bson:"user_name" json:"user_name,omitempty"`
	Email     string        `bson:"email" json:"email,omitempty"`
	Password  string        `bson:"password" json:"password,omitempty"`
	CreatedAt int64         `bson:"created_at" json:"created_at"`
	UpdatedAt int64         `bson:"updated_at" json:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) Insert() error {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now().UTC().Unix()

	return c.Insert(user)
}

func UpdateUserByID(id string, user *User) error {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user.UpdatedAt = time.Now().UTC().Unix()

	return c.UpdateId(bson.ObjectIdHex(id), bson.M{
		"$set": user,
	})
}

func UpdateUser(selector interface{}, user *User) error {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user.UpdatedAt = time.Now().UTC().Unix()

	return c.Update(selector, bson.M{
		"$set": user,
	})
}

func UpdateUserAll(selector interface{}, user *User) (*mgo.ChangeInfo, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user.UpdatedAt = time.Now().UTC().Unix()

	return c.UpdateAll(selector, bson.M{
		"$set": user,
	})
}

func FindUserByID(id string) (*User, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user := new(User)

	if bson.IsObjectIdHex(id) {
		return user, c.FindId(id).One(user)
	}

	return user, c.FindId(bson.ObjectIdHex(id)).One(user)
}

func FindUserByQuery(query interface{}) (*User, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user := new(User)

	return user, c.Find(query).One(user)
}

func FindAllUserByQuery(query interface{}) ([]*User, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user := make([]*User, 0)

	return user, c.Find(query).All(&user)
}

func ExistUserByID(id string) (bool, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user := new(User)

	if err := c.FindId(bson.ObjectIdHex(id)).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func ExistUserByQuery(query interface{}) (bool, error) {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	user := new(User)

	if err := c.Find(query).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func DeleteUserByID(id string) error {
	s, c := GetSessionAndCollection(CollectionUser)
	defer s.Close()

	return c.RemoveId(bson.ObjectIdHex(id))
}
