package models

import (
	"github.com/lib/pq"
)

// Todo is the model for a single todo item
type Todo struct {
	ID         int
	ListID     int
	Name       string
	IsComplete bool
	Created    pq.NullTime
	Deleted    pq.NullTime
	Sort       int
}
