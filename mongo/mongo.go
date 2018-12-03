package config

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/lbrulet/web-app-golang/config"
	"github.com/lbrulet/web-app-golang/mongo/models"
)

type usersDAO struct {
	server     string
	database   string
	collection string
	db         *mgo.Database
}

var (
	// Users is ...
	Users usersDAO
)

func init() {
	Users.collection = "users"
	Users.server = config.MongoAddress
	Users.database = "users_db"
	Users.connect()
}

func (m *usersDAO) connect() {
	session, err := mgo.Dial(m.server)
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.database)
}

func (m *usersDAO) getCollection() *mgo.Collection {
	return m.db.C(m.collection)
}

func (m *usersDAO) FindAll() ([]models.User, error) {
	var users []models.User
	err := m.getCollection().Find(bson.M{}).All(&users)
	return users, err
}

func (m *usersDAO) FindById(id string) (models.User, error) {
	var user models.User
	err := m.getCollection().FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *usersDAO) Insert(user models.User) error {
	err := m.getCollection().Insert(&user)
	return err
}

func (m *usersDAO) Delete(user models.User) error {
	err := m.getCollection().Remove(&user)
	return err
}

func (m *usersDAO) Update(user models.User) error {
	err := m.getCollection().UpdateId(user.ID, &user)
	return err
}
