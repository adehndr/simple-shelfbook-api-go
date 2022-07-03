package repository

import (
	"example.com/shelfbook-api/model/domain"
	"example.com/shelfbook-api/model/web"
)

type BookshelfRepository interface {
	Create(book domain.Book) (web.BookResponse, error)
	FindAll() []web.WebResponseGet
	FindById(bookId string) (web.BookResponse, error)
	Update(bookId string, book domain.Book) (web.BookResponse, error)
	Delete(bookId string) error
}
