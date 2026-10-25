package routers

import (
	controller "golang_quiz3/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	// route category
	router.POST("/api/category", controller.CreateCategory)
	router.GET("/api/category", controller.GetAllCategory)
	router.PUT("/api/category/:id", controller.UpdateCategory)
	router.DELETE("/api/category/:id", controller.DeleteCategory)

	// route users
	router.POST("/api/users", controller.CreateUser)
	router.GET("/api/users", controller.GetAllUsers)
	router.PUT("/api/users/:id", controller.UpdateUser)
	router.DELETE("/api/users/:id", controller.DeleteUser)

	// route books
	router.POST("/api/books", controller.CreateBook)
	router.GET("/api/books", controller.GetAllBooks)
	router.PUT("/api/books/:id", controller.UpdateBook)
	router.DELETE("/api/books/:id", controller.DeleteBook)

	return router
}
