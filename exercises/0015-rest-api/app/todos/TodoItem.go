package todos

import "time"

type TodoItem struct {
	Id          uint
	Done        bool
	Description string
	CreatedAt   time.Time
}
