// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package gen

import (
	"time"
)

type Todo struct {
	ID          int64
	Description string
	Done        bool
	Created     time.Time
	Updated     time.Time
}
