package model

import "time"

type Todo struct {
	Id          int64
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}
