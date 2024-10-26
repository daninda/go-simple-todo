package router

import (
	"simple-todo/internal/adapter/http/handler"

	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.TodoHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todos", handler.FindAll).Methods("GET")
	router.HandleFunc("/todos/{id:[0-9]+}", handler.FindOne).Methods("GET")
	router.HandleFunc("/todos", handler.Create).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", handler.Update).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", handler.Delete).Methods("DELETE")

	return router
}
