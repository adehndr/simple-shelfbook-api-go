package main

import (
	"net/http"

	"example.com/shelfbook-api/app"
	"example.com/shelfbook-api/controller"
	"example.com/shelfbook-api/helper"
	"example.com/shelfbook-api/repository"
	"example.com/shelfbook-api/service"
)

func main() {
	bookShelfCategory := repository.NewBookshelfRepositoryImpl()
	bookShelfService := service.NewBookshelfRepository(bookShelfCategory)
	bookShelfController := controller.NewBookshelfRepository(bookShelfService)
	bookShelfRouter := app.NewRouter(bookShelfController)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: bookShelfRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
