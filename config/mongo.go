package dao

import (
	"log"

	models "github.com/lbrulet/web-app-golang/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *UsersDAO) FindAll() ([]models.User, error) {
	var users []models.User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (m *UsersDAO) FindById(id string) (models.User, error) {
	var user models.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UsersDAO) Insert(user models.User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *UsersDAO) Delete(user models.User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *UsersDAO) Update(user models.User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
