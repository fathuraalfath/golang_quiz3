package routers

import (
	controller "golang_quiz3/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	auth := controller.AuthMiddleware()

	// route category
	router.POST("/api/category", auth, controller.CreateCategory)
	router.GET("/api/category", auth, controller.GetAllCategory)
	router.PUT("/api/category/:id", auth, controller.UpdateCategory)
	router.DELETE("/api/category/:id", auth, controller.DeleteCategory)
	router.GET("/api/category/:id/books", auth, controller.GetAllBookByCategory)

	// route users
	router.POST("/api/users", controller.CreateUser)
	router.GET("/api/users", controller.GetAllUsers)
	router.PUT("/api/users/:id", controller.UpdateUser)
	router.DELETE("/api/users/:id", controller.DeleteUser)

	// route books
	router.POST("/api/books", auth, controller.CreateBook)
	router.GET("/api/books", auth, controller.GetAllBooks)
	router.PUT("/api/books/:id", auth, controller.UpdateBook)
	router.DELETE("/api/books/:id", auth, controller.DeleteBook)

	return router
}
