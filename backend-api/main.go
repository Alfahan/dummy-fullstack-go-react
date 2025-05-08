package main

import (
	"dummy-fullstack-go-react/backend-api/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load config .env
	config.LoadEnv()

	//inisialiasai Gin
	router := gin.Default()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//mulai server
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))

}
