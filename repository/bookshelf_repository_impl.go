package repository

import (
	"errors"
	"strings"
	"time"

	"example.com/shelfbook-api/helper"
	"example.com/shelfbook-api/model/domain"
	"example.com/shelfbook-api/model/web"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var BookShelfDB []web.BookResponse = []web.BookResponse{}

type BookshelfRepositoryImpl struct {
}

func NewBookshelfRepositoryImpl() BookshelfRepository {
	return &BookshelfRepositoryImpl{}
}

func (repository *BookshelfRepositoryImpl) Create(book domain.Book) (web.BookResponse, error) {
	id, err := gonanoid.New()
	helper.PanicIfError(err)
	isFinished := book.PageCount == book.ReadPage
	currentTime := time.Now()
	tempBook := web.BookResponse{
		Id:         id,
		Name:       book.Name,
		Year:       book.Year,
		Author:     book.Author,
		Summary:    book.Summary,
		Publisher:  book.Publisher,
		PageCount:  book.PageCount,
		ReadPage:   book.ReadPage,
		Finished:   isFinished,
		Reading:    book.Reading,
		InsertedAt: currentTime,
		UpdatedAt:  currentTime,
	}
	BookShelfDB = append(BookShelfDB, tempBook)
	isSuccess := func() bool {
		isContained := func() bool {
			for _, br := range BookShelfDB {
				if br.Id == id {
					return true
				}
			}
			return false
		}()
		if len(BookShelfDB) > 0 && isContained {
			return true
		} else {
			return false
		}
	}()
	if isSuccess {
		return tempBook, nil
	}
	return web.BookResponse{}, errors.New("Buku gagal ditambahkan")

}

func (repository *BookshelfRepositoryImpl) FindAll(queryParam domain.QueryParam) []web.WebResponseGet {
	tempBookShelfDB := []web.WebResponseGet{}
	for _, br := range BookShelfDB {
		if queryParam.NameParam != "" && strings.Contains(strings.ToLower(br.Name), strings.ToLower(queryParam.NameParam)) {
			tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
		} else {
			if queryParam.IsReading == nil && queryParam.IsFinished == nil && queryParam.NameParam == "" {
				tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
			}
		}
		if queryParam.IsFinished != nil {
			if *queryParam.IsFinished == true && br.Finished == true {
				tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
			}
			if *queryParam.IsFinished == false && br.Finished == false {
				tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
			}
		}

		if queryParam.IsReading != nil {
			if *queryParam.IsReading == true && br.Reading == true {
				tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
			}
			if *queryParam.IsReading == false && br.Reading == false {
				tempBookShelfDB = append(tempBookShelfDB, web.WebResponseGet{BookId: br.Id, BookName: br.Name, BookPublisher: br.Publisher})
			}
		}
	}
	return tempBookShelfDB
}

func (repository *BookshelfRepositoryImpl) FindById(bookId string) (web.BookResponse, error) {
	tempBookFound := web.BookResponse{}
	isBookFound := func(bookId string) bool {
		for _, item := range BookShelfDB {
			if item.Id == bookId {
				tempBookFound = item
				return true
			}
		}
		return false
	}(bookId)
	if isBookFound {
		return tempBookFound, nil
	}
	return web.BookResponse{}, errors.New("Buku tidak ditemukan")
}

func (repository *BookshelfRepositoryImpl) Update(bookId string, book domain.Book) (web.BookResponse, error) {
	var tempBook *web.BookResponse
	isFound := false
	for index, bookReal := range BookShelfDB {
		if bookReal.Id == bookId {
			tempBook = &BookShelfDB[index]
			isFound = true
		}
	}
	if isFound == false {
		return web.BookResponse{}, errors.New("Gagal memperbarui buku. Id tidak ditemukan")
	}
	(*tempBook).Name = book.Name
	(*tempBook).Year = book.Year
	(*tempBook).Author = book.Author
	(*tempBook).Summary = book.Summary
	(*tempBook).Publisher = book.Publisher
	(*tempBook).PageCount = book.PageCount
	(*tempBook).ReadPage = book.ReadPage
	(*tempBook).Reading = book.Reading

	return web.BookResponse{}, nil
}

func (repository *BookshelfRepositoryImpl) Delete(bookId string) error {
	isFound := false
	var indexDeletedBook int
	for index, bookReal := range BookShelfDB {
		if bookReal.Id == bookId {
			indexDeletedBook = index
			isFound = true
		}
	}
	if len(BookShelfDB) > 0 {
		deleteBookById(&BookShelfDB, indexDeletedBook)
	}

	if isFound == false {
		return errors.New("Buku gagal dihapus. Id tidak ditemukan")
	}
	return nil
}

func deleteBookById(books *[]web.BookResponse, idx int) {
	(*books)[idx] = (*books)[len(*books)-1]
	(*books) = (*books)[:len(*books)-1]
}
