package app

import (
	"example.com/shelfbook-api/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.BookShelfController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/books", categoryController.FindAll)
	router.GET("/books/:bookId", categoryController.FindById)
	router.POST("/books", categoryController.Create)
	router.PUT("/books/:bookId", categoryController.Update)
	router.DELETE("/books/:bookId", categoryController.Delete)
	// router.PanicHandler = exception.InternalServerError
	return router
}
