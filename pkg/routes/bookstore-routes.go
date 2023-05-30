package routes

import (
	"github.com/Sosuke-Uchiha/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
