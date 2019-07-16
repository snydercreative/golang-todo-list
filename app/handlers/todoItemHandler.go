package handlers

import (
	"encoding/json"
	"golang-todo-list/app/services"
	"golang-todo-list/app/util"
	"net/http"
)

// TodoItemHandler is a collection of verb handler methods
type TodoItemHandler struct{}

// Get handles GET requests
func (t *TodoItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	service := services.TodoItemService{}

	todo, err := service.Get(1)

	util.Check(err)

	json.NewEncoder(w).Encode(todo)
}

// Patch handles PATCH requests
func (t *TodoItemHandler) Patch(w http.ResponseWriter, r *http.Request) {

}

// Delete handles DELETE requests
func (t *TodoItemHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
