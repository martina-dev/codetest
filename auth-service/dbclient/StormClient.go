package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/ksuid"

	"github.com/asdine/storm"
	"github.com/martina-dev/authapp/auth-service/schema"
)

//IRepository contract to be satisfied by all the DB implementations
type IRepository interface {
	OpenDB()
	CreateUser(name, email, password, role string) (schema.User, error)
	GetUserByEmail(email string) (schema.User, error)
	GetUsers() ([]schema.User, error)
}

//StormRepo reference to the in-memory DB {Storm DB}
type StormRepo struct {
	db *storm.DB
}

//OpenDB opens the DB
func (sr *StormRepo) OpenDB() {
	var err error
	sr.db, err = storm.Open("auth.db")
	err = sr.db.Init(&schema.User{})
	if err != nil {
		fmt.Println("Could not open DB", err)
	}
}

func genID(time time.Time) (string, error) {
	id, err := ksuid.NewRandomWithTime(time)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

//CreateUser adds a user to the Database
func (sr *StormRepo) CreateUser(name, email, password, role string) (schema.User, error) {
	createdAt := time.Now().UTC()
	Id, _ := genID(createdAt)
	user := schema.User{
		ID:       Id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}

	err := sr.db.Save(&user)
	if err != nil {
		return schema.User{}, fmt.Errorf("could not add user %s", err)
	}

	return user, nil
}

//GetUserByEmail returns a user with a matching email
func (sr *StormRepo) GetUserByEmail(email string) (schema.User, error) {
	var user schema.User

	err := sr.db.One("Email", email, &user)
	if err != nil {
		log.Printf("Im here %s", err)
		return schema.User{}, fmt.Errorf("Could not find user %s", err)
	}

	return user, nil
}

//GetUsers returns all users in the DB
func (sr *StormRepo) GetUsers() ([]schema.User, error) {
	var users []schema.User
	err := sr.db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
