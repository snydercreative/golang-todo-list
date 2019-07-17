package models

import (
	"github.com/lib/pq"
)

// List is the model for a single list item
type List struct {
	ID      int
	Name    string
	Created pq.NullTime
	Deleted pq.NullTime
}
