package models

import "time"

// List is the model for a single list item
type List struct {
	id      int
	name    string
	created time.Time
	deleted time.Time
}
