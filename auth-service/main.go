package main

import (
	"github.com/martina-dev/authapp/auth-service/dbclient"
	"github.com/martina-dev/authapp/auth-service/service"
)

func main() {
	initializeDB()
	service.NewServer("8000")
}

func initializeDB() {
	service.DBClient = &dbclient.StormRepo{}
	service.DBClient.OpenDB()
}
