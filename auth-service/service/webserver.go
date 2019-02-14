package service

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//NewServer creates a new instance of a server
func NewServer(port string) {
	log.Println("Server is running at: " + port)
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateUserEndpoint)
	router.HandleFunc("/login", LoginEndpoint)
	router.HandleFunc("/users", JwtAuthentication(GetUsersEndpoint))
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router)))
}
