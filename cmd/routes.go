package cmd

import (
	"golang-todo-list/app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Authenticate	POST 	/v1/api/authentication

List todo lists		GET     /v1/api/lists
New todo list		POST    /v1/api/lists
Update order		PATCH   /v1/api/lists
Delete todo list	DELETE	/v1/api/lists/{list id}

List my todos		GET     /v1/api/lists/{list id}/todos
Add individual todo	POST    /v1/api/lists/{list id}/todos

Get individual todo		GET     /v1/api/lists/{list id}/todos/{todoId}
Update individual todo	PATCH   /v1/api/lists/{list id}/todos/{todoId}
Delete individual todo	DELETE  /v1/api/lists/{list id}/todos/{todoId}
*/

// Routes - mappings for the endpoints
func routes() {
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
