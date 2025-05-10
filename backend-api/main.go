package main

import (
	"dummy-fullstack-go-react/backend-api/config"
	"dummy-fullstack-go-react/backend-api/database"
	"dummy-fullstack-go-react/backend-api/routes"
)

func main() {

	// Load config .env
	config.LoadEnv()

	database.InitDB()

	// setup router
	r := routes.SetupRoutes()

	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))

}
