package routes

import (
	"dummy-fullstack-go-react/backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	return router
}
