package services

import (
	"golang-todo-list/app/models"
	"golang-todo-list/app/repositories"
	"golang-todo-list/app/util"
)

// TodoItemService manages the todo items
type TodoItemService struct{}

// Get returns a specific todo item
func (t *TodoItemService) Get(todoID int) (*models.Todo, error) {
	repo := repositories.TodoItemRepository{}

	todo, err := repo.Get(1)

	util.Check(err)

	return todo, nil
}
