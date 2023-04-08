package routes

import (
	"github.com/atorahmad/go/crud/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	books := router.Group("/books")
	{
		books.POST("/", controllers.CreateBook)
		books.GET("/", controllers.GetBooks)
		books.GET("/:id", controllers.GetBookById)
		books.PUT("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}
