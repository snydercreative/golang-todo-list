package models

import "time"

// List is the model for a single list item
type List struct {
	ID      int
	Name    string
	Created time.Time
	Deleted time.Time
}
