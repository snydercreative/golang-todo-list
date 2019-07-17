package handlers

import (
	"encoding/json"
	"golang-todo-list/app/services"
	"golang-todo-list/app/util"
	"net/http"
)

// TodoListHandler contains methods to deal with HTTP verbs
type TodoListHandler struct{}

// Get handles the GET requests
func (t *TodoListHandler) Get(w http.ResponseWriter, r *http.Request) {
	service := services.TodoListService{}

	lists, err := service.Get(1)

	util.Check(err)

	jsonEncoder := json.NewEncoder(w)

	jsonEncoder.Encode(lists)
}

// Post handles the POST requests
func (t *TodoListHandler) Post(w http.ResponseWriter, r *http.Request) {

}
