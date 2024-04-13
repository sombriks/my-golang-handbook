package app

import (
	"0015-rest-api/app/todos"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/pug/v2"
)

type Server struct {
	app            *fiber.App
	todoService    *todos.TodoService
	todoController *todos.TodoController
}

// Start - Quick setup fiber app and start to listen
func (s *Server) Start() error {

	var err error

	s.todoController, err = todos.NewController()
	if err != nil {
		return err
	}

	// setup fiber
	s.app = fiber.New(fiber.Config{
		Views: pug.New("./app/templates", ".pug"),
	})
	s.app.Use(logger.New())

	// routes configuration
	s.app.Static("/", "./app/assets")
	s.app.Get("/", s.todoController.Index)

	return s.app.Listen(":3000")

}
