package repositories

import "golang-todo-list/app/models"

// TodoItemRepository manages data access for the todo items
type TodoItemRepository struct{}

// Get returns a specific todo item
func (t *TodoItemRepository) Get(todoID int) *models.Todo {
	return new(models.Todo)
}
