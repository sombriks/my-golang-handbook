package models

import "time"

// Todo - our todo item, full of annotations
type Todo struct {
	Id          uint64    `json:"id" db:"id" goqu:"skipinsert"`
	Description string    `json:"description" db:"description"`
	Done        bool      `json:"done" db:"done"`
	Created     time.Time `json:"created" db:"created" goqu:"skipinsert,skipupdate"`
	Updated     time.Time `json:"updated" db:"updated" goqu:"skipinsert"`
}
