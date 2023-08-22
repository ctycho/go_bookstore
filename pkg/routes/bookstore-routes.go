package routes

import (
	"github.com/gorilla/mux"
	"github.com/ctycho/go_bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}