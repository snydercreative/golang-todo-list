package models

import "time"

// Todo is the model for a single todo item
type Todo struct {
	ID         int
	ListID     int
	Name       string
	IsComplete bool
	Created    time.Time
	Deleted    time.Time
	Sort       int
}
