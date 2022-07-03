package service

import (
	"errors"

	"example.com/shelfbook-api/helper"
	"example.com/shelfbook-api/model/domain"
	"example.com/shelfbook-api/model/web"
	"example.com/shelfbook-api/repository"
)

type BookshelfServiceImpl struct {
	bookshelfRepository repository.BookshelfRepository
}

func NewBookshelfRepository(bookshelfRepository repository.BookshelfRepository) BookshelfService {
	return &BookshelfServiceImpl{
		bookshelfRepository: bookshelfRepository,
	}
}

func (service *BookshelfServiceImpl) Create(book domain.Book) (web.WebResponse, error) {
	/*
		Validation from body request
	*/
	if book.Name == "" {
		errMsg := "Gagal menambahkan buku. Mohon isi nama buku"
		return web.WebResponse{
			Status:  "fail",
			Message: errMsg,
		}, errors.New(errMsg)
	} else if book.ReadPage > book.PageCount {
		errMsg := "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount"
		return web.WebResponse{
			Status:  "fail",
			Message: errMsg,
		}, errors.New(errMsg)
	}

	data, err := service.bookshelfRepository.Create(book)
	if err != nil {
		if err.Error() == "Buku gagal ditambahkan" {
			return web.WebResponse{
				Status:  "fail",
				Message: err.Error(),
			}, errors.New(err.Error())
		}
		helper.PanicIfError(err)
	}
	bookid := web.WebResponseId{
		BookId: data.Id,
	}
	return web.WebResponse{
		Status:  "success",
		Message: "Buku berhasil ditambahkan",
		Data:    bookid,
	}, nil
}

func (service *BookshelfServiceImpl) FindAll(queryParam domain.QueryParam) web.WebResponse {
	webResponse := web.WebResponse{
		Status: "success",
		Data:   web.WebResponseGetAll{Books: service.bookshelfRepository.FindAll(queryParam)},
	}
	return webResponse
}

func (service *BookshelfServiceImpl) FindById(bookId string) (web.WebResponse, error) {
	webResponse, err := service.bookshelfRepository.FindById(bookId)
	if err != nil {
		if err.Error() == "Buku tidak ditemukan" {
			return web.WebResponse{
				Status:  "fail",
				Message: err.Error(),
			}, err
		}
		helper.PanicIfError(err)
	}
	return web.WebResponse{Status: "success", Data: web.WebResponseGetById{Book: webResponse}}, nil
}

func (service *BookshelfServiceImpl) Update(bookId string, book domain.Book) (web.WebResponse, error) {

	if book.Name == "" {
		errMsg := "Gagal memperbarui buku. Mohon isi nama buku"
		return web.WebResponse{Status: "fail", Message: errMsg}, errors.New(errMsg)
	}

	if book.ReadPage > book.PageCount {
		errMsg := "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount"
		return web.WebResponse{Status: "fail", Message: errMsg}, errors.New(errMsg)
	}

	_, err := service.bookshelfRepository.Update(bookId, book)
	if err != nil {
		errMsg := "Gagal memperbarui buku. Id tidak ditemukan"
		if err.Error() == errMsg {
			return web.WebResponse{Status: "fail", Message: errMsg}, errors.New(errMsg)
		} else {
			helper.PanicIfError(err)
		}
	}
	return web.WebResponse{
		Status:  "success",
		Message: "Buku berhasil diperbarui",
	}, nil
}

func (service *BookshelfServiceImpl) Delete(bookId string) (web.WebResponse, error) {
	err := service.bookshelfRepository.Delete(bookId)
	if err != nil {
		if err.Error() == "Buku gagal dihapus. Id tidak ditemukan" {
			return web.WebResponse{Status: "fail" , Message: err.Error()},errors.New(err.Error())
		}else {
			helper.PanicIfError(err)
		}
	}
	return web.WebResponse{Status: "success", Message: "Buku berhasil dihapus"},nil
}
