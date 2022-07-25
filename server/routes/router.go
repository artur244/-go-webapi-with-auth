package routes

import (
	"github.com/artur244/first-go-rest-api/controllers"
	"github.com/artur244/first-go-rest-api/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

		users := main.Group("users")
		{
			users.POST("/", controllers.CreateUser)
		}

		main.POST("login", controllers.Login)
	}

	return router
}
