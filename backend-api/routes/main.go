package routes

import (
	"dummy-fullstack-go-react/backend-api/controllers"
	"dummy-fullstack-go-react/backend-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return router
}
