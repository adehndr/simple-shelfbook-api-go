package service

import (
	"example.com/shelfbook-api/model/domain"
	"example.com/shelfbook-api/model/web"
)

type BookshelfService interface {
	Create(book domain.Book) (web.WebResponse, error)
	FindAll(queryParam domain.QueryParam) web.WebResponse
	FindById(bookId string) (web.WebResponse, error)
	Update(bookId string, book domain.Book) (web.WebResponse, error)
	Delete(bookId string) (web.WebResponse, error)
}
