package models

import "time"

// Todo is the model for a single todo item
type Todo struct {
	id         int
	listID     int
	name       string
	isComplete bool
	created    time.Time
	deleted    time.Time
	sort       int
}
