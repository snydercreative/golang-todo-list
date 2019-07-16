package handlers

import (
	"net/http"
)

// TodoListsHandler is the collection of handler methods for the todo list
type TodoListsHandler struct{}

// Get handles the GET requests and lists the todo lists
func (t *TodoListsHandler) Get(w http.ResponseWriter, r *http.Request) {

}

// Post handles the POST requests and creates a new todo list
func (t *TodoListsHandler) Post(w http.ResponseWriter, r *http.Request) {

}

// Patch handles the PATCH requests and updates the list item order
func (t *TodoListsHandler) Patch(w http.ResponseWriter, r *http.Request) {

}

// Delete handles the DELETE requests and removes a list
func (t *TodoListsHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
