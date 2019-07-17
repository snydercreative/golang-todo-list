package repositories

import (
	"golang-todo-list/app/models"
	"golang-todo-list/app/util"

	"github.com/lib/pq"
)

// TodoListRepository manages data access for the todo lists
type TodoListRepository struct {
}

// Get returns a specific todo item
func (t *TodoListRepository) Get(listID int) ([]*models.List, error) {
	var lists []*models.List

	rows, err := db.Query("SELECT id, name, created, deleted FROM todolist")

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

		lists = append(lists, &models.List{
			ID:      id,
			Name:    name,
			Created: created,
			Deleted: deleted,
		})
	}

	return lists, nil
}
