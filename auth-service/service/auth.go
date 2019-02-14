package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Name  string
	Email string
	Role  string
	jwt.StandardClaims
}

var MySigningKey = "secret"

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

var JwtAuthentication = func(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/create", "/api/login"}

		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Message(false, "Invalid auth Token"))
			return
		}

		token := strings.Split(tokenHeader, " ")
		tk := &Token{}

		if len(token) != 2 {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Message(false, "Invalid/ malformed auth Token"))
			return
		}
		tokenPart := token[1]

		tokenString, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(MySigningKey), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Message(false, "Invalid/ malformed auth Token"))
		}

		if !tokenString.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Message(false, "Invalid/ malformed auth Token"))
		}

		log.Println(fmt.Sprintf("User %s", tk.Name)) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
