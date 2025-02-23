package routes

import (
	"github.com/gorilla/mux"
	"github.com/milan-kovac/packages/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
  router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
  router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
  router.HandleFunc("book/{bookId}", controllers.GetBookById).Methods("GET")
  router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}