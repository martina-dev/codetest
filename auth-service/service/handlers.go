package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/martina-dev/authapp/auth-service/dbclient"
	"golang.org/x/crypto/bcrypt"
)

//DBClient reference to the DB
var DBClient dbclient.IRepository

//CreateUserEndpoint endpoint that creates a new user
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	auth := r.Header.Get("Authorization")
	log.Println(auth)
	user, err := DBClient.CreateUser(email, name, string(hashedPassword), role)
	if err != nil {
		log.Fatal("Could not add user ", err)
	}

	theToken := GenerateToken(user.Name, user.Email, user.Role)
	json.NewEncoder(w).Encode(theToken)
	fmt.Println(user)
	fmt.Println(theToken)

}

//LoginEndpoint logs a user in
func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := DBClient.GetUserByEmail(email)
	if err != nil {
		json.NewEncoder(w).Encode(Message(false, err.Error()))
	}

	if err == nil {
		fmt.Println(user)

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			json.NewEncoder(w).Encode(Message(false, "Invalid password"))
		}

		token := GenerateToken(user.Name, user.Email, user.Role)
		json.NewEncoder(w).Encode(token)
		fmt.Println(token)
	}
}

//GetUsersEndpoint returns all the users
func GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	log.Println(auth)
	users, err := DBClient.GetUsers()
	if err != nil {
		json.NewEncoder(w).Encode(Message(false, " Could not get users"))
	}
	json.NewEncoder(w).Encode(users)
}

//GenerateToken generates a new token
func GenerateToken(name, email, role string) string {
	tk := &Token{Name: name, Email: email, Role: role}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(MySigningKey))

	return tokenString
}
