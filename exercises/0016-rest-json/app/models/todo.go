package models

import "time"

// Todo - our todo item
type Todo struct {
	Id          uint64    `json:"id" db:"id"`
	Description string    `json:"description" db:"description"`
	Done        bool      `json:"done" db:"done"`
	Created     time.Time `json:"created" db:"created"`
	Updated     time.Time `json:"updated" db:"updated"`
}
