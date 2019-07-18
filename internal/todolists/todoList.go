package todolists

import (
	"golang-todo-list/internal/util"

	"github.com/lib/pq"
)

// TodoList is the model for a single list item
type TodoList struct {
	ID      int
	Name    string
	Created pq.NullTime
	Deleted pq.NullTime
}

// GetTodoLists returns a all todo lists
func GetTodoLists() ([]*TodoList, error) {
	var lists []*TodoList

	rows, err := util.Database.Query("SELECT id, name, created, deleted FROM todolist")

	util.Check(err)

	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			name    string
			created pq.NullTime
			deleted pq.NullTime
		)

		err := rows.Scan(&id, &name, &created, &deleted)

		util.Check(err)

		lists = append(lists, &TodoList{
			ID:      id,
			Name:    name,
			Created: created,
			Deleted: deleted,
		})
	}

	return lists, nil
}

// GetTodoList returns a specific todo list
func GetTodoList(ID int) ([]*TodoList, error) {
	var lists []*TodoList

	rows, err := util.Database.Query("SELECT id, name, created, deleted FROM todolist WHERE id = ?", ID)

	util.Check(err)

	defer rows.Close()

	for rows.Next() {
		var (
			id      int
			name    string
			created pq.NullTime
			deleted pq.NullTime
		)

		err := rows.Scan(&id, &name, &created, &deleted)

		util.Check(err)

		lists = append(lists, &TodoList{
			ID:      id,
			Name:    name,
			Created: created,
			Deleted: deleted,
		})
	}

	return lists, nil
}
