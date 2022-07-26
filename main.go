package main

import (
	"fmt"

	"github.com/RaguSoftEng/inviteservice_api/src/controllers"
	"github.com/RaguSoftEng/inviteservice_api/src/repository"
	"github.com/RaguSoftEng/inviteservice_api/util"
)

func main() {
	//Init configuration variables
	util.InitVariables()

	//Make sure config file loaded correctly
	fmt.Printf("%+v\n", util.AppConfig)

	// Connect Database (mysql)
	db := util.Connect()

	// Initial Migration
	util.Migrate()

	// Seed user records
	repository.ImportUsers(db)

	var server = controllers.Server{}

	// Initialize the server
	server.Initialize()

	server.Run(":4000")
}
