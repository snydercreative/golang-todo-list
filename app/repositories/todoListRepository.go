package repositories

import (
	"golang-todo-list/app/models"
	"golang-todo-list/app/util"
)

// TodoListRepository manages data access for the todo lists
type TodoListRepository struct{}

// Get returns a specific todo item
func (t *TodoListRepository) Get(listID int) ([]*models.List, error) {
	var lists []*models.List
	var list *models.List

	rows, err := db.Query("SELECT id, name FROM todolist")

	util.Check(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(list.ID, list.Name)

		util.Check(err)

		lists = append(lists, list)
	}

	return lists, nil
}
