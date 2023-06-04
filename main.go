package main

import (
	"auto-reload/database"
	"auto-reload/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
