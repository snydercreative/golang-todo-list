package cmd

import (
	"golang-todo-list/app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/v1/api").Subrouter()

	todoListsHandler := new(handlers.TodoListsHandler)
	todoListHandler := new(handlers.TodoListHandler)
	todoItemHander := new(handlers.TodoItemHandler)

	subrouter.HandleFunc("/lists", todoListsHandler.Get).Methods("GET")
	subrouter.HandleFunc("/lists", todoListsHandler.Post).Methods("POST")
	subrouter.HandleFunc("/lists", todoListsHandler.Patch).Methods("PATCH")
	subrouter.HandleFunc("/lists/{listId}", todoListsHandler.Delete).Methods("DELETE")

	subrouter.HandleFunc("/lists/{listId}/todos", todoListHandler.Get).Methods("GET")
	subrouter.HandleFunc("/lists/{listId}/todos", todoListHandler.Post).Methods("POST")

	subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Get).Methods("GET")
	subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Patch).Methods("PATCH")
	subrouter.HandleFunc("/lists/{listId}/todos/{todoId}", todoItemHander.Delete).Methods("DELETE")

	http.Handle("/", router)
}
