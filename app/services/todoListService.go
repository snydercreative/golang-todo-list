package services

import (
	"golang-todo-list/app/models"
	"golang-todo-list/app/repositories"
	"golang-todo-list/app/util"
)

// TodoListService manages the todo lists
type TodoListService struct{}

// Get returns a specific todo item
func (t *TodoListService) Get(todoID int) ([]*models.List, error) {
	repo := repositories.TodoListRepository{}

	lists, err := repo.Get(1)

	util.Check(err)

	return lists, nil
}
