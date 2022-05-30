package main

import (
	"fmt"
	"net/http"
	"yamlrest/controller"
	"yamlrest/services/database"
	"yamlrest/utils/context"
)

func main() {

	// Using in memory DB here. If actual DB, initialize connection here
	database := database.CreateInMemDB()

	// Create app context
	appContext := context.AppContext{
		Database: database,
	}

	// Make server instance
	server := controller.CreateServer(&appContext)

	http.Handle("/", server.Routers)
	fmt.Println("Server hosted on port 4400")
	error := http.ListenAndServe("localhost:4400", server.Routers)
	if error != nil {
		fmt.Println(error.Error())
	}
}
