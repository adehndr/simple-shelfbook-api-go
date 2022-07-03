package controller

import (
	"encoding/json"
	"net/http"

	"example.com/shelfbook-api/helper"
	"example.com/shelfbook-api/model/domain"
	"example.com/shelfbook-api/service"
	"github.com/julienschmidt/httprouter"
)

type BookShelfControllerImpl struct {
	service service.BookshelfService
}

func NewBookshelfRepository(bookshelfService service.BookshelfService) BookShelfController {
	return &BookShelfControllerImpl{
		service: bookshelfService,
	}
}

func (controller *BookShelfControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bookRequested := domain.Book{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bookRequested)
	helper.PanicIfError(err)
	webResponse, err := controller.service.Create(bookRequested)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		if err.Error() == "Buku gagal ditambahkan" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err.Error() == "Gagal menambahkan buku. Mohon isi nama buku" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err.Error() == "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			helper.PanicIfError(err)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)

}

func (controller *BookShelfControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	webResponse := controller.service.FindAll()
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (controller *BookShelfControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams := params.ByName("bookId")
	webResponse, err := controller.service.FindById(idParams)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		if err.Error() == "Buku tidak ditemukan" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			helper.PanicIfError(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (controller *BookShelfControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams := params.ByName("bookId")
	bookBodyRequest := domain.Book{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bookBodyRequest)
	if err != nil {
		helper.PanicIfError(err)
	}

	webResponse, err := controller.service.Update(idParams, bookBodyRequest)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		if err.Error() == "Gagal memperbarui buku. Mohon isi nama buku" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err.Error() == "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount" {
			w.WriteHeader(http.StatusBadRequest)
		} else if err.Error() == "Gagal memperbarui buku. Id tidak ditemukan" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			helper.PanicIfError(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (controller *BookShelfControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idParams := params.ByName("bookId")
	webResponse, err := controller.service.Delete(idParams)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		if err.Error() == "Buku gagal dihapus. Id tidak ditemukan" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			helper.PanicIfError(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}
