package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Naman15032001/go-sql-bookstore/pkg/models"
	"github.com/Naman15032001/go-sql-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error parsing book id")
	}
	newbook, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(newbook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the request create book")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	fmt.Println("body", b)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error parsing book id")
	}
	deletedbook := models.DeleteBook(bookId)
	res, _ := json.Marshal(deletedbook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	id := mux.Vars(r)["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error parsing book id")
	}
	booksDetails, db := models.GetBookById(bookId)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
