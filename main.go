package main

import (
	"github.com/artur244/first-go-rest-api/database"
	"github.com/artur244/first-go-rest-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
