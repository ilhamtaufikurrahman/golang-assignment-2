package main

import (
	"assignment-2/database"
	"assignment-2/router"
)

var PORT = ":8080"

func main() {
	database.StartDB()

	router.StartServer().Run(PORT)
}
