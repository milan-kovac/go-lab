package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/milan-kovac/packages/models"
	"github.com/milan-kovac/packages/utils"
)

var NewBook models.Book

func GetBooks(res http.ResponseWriter, req *http.Request){
	books := models.GetAllBooks()

	r,_ := json.Marshal(books)
	
	res.Header().Set("Content-Type", "aplication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r);
}

func GetBookById(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0);
	if err != nil{
		fmt.Println("Error while parsing.")
	}

	book, _ := models.GetBookById(ID)

	r, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "aplication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r);
}

func CreateBook(res http.ResponseWriter, req *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	book := CreateBook.CreateBoook()

	r, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "aplication/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(r);
}


func UpdateBook(res http.ResponseWriter, req *http.Request){
	
}

func DeleteBook(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0);

	if err != nil{
		fmt.Println("Error while parsing.")
	}

	book := models.DeleteBook(ID)

	r, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "aplication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(r);
}