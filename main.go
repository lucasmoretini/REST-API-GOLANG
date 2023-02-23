package main

import (
	"rest-api-go/database"
	"rest-api-go/routes"
)

func main() {
	database.DatabaseConnection()
	routes.HandleRequest()
}
