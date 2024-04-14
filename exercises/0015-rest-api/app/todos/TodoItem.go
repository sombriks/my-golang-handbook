package todos

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
	"time"
)

type TodoItem struct {
	Id          uint64 `gorm:"primaryKey"`
	Description string
	Done        bool
	Created     time.Time `gorm:"autoCreateTime"`
}

// https://gorm.io/docs/conventions.html
func (TodoItem) TableName() string {
	return "todos"
}

func TodoFromRequest(ctx fiber.Ctx) TodoItem {
	var todo TodoItem // TODO again the form parser issue
	todo.Description = ctx.FormValue("description")
	todo.Id, _ = strconv.ParseUint(ctx.FormValue("id"), 10, 32)
	todo.Done, _ = strconv.ParseBool(ctx.FormValue("done"))
	return todo
}
