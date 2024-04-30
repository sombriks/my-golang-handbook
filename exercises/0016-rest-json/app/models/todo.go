package models

import "time"

type Todo struct {
	Id          uint64
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}
