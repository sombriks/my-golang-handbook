package app

import (
	"0015-rest-api/app/configs"
	"0015-rest-api/app/todos"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/pug/v2"
)

type Server struct {
	app            *fiber.App
	database       *configs.Database
	todoService    *todos.TodoService
	todoController *todos.TodoController
}

// Start - Quick setup fiber app and start to listen
func (s *Server) Start(service any) error {
	var err error
	// configuration phase
	s.database, err = configs.NewDatabase()
	if err != nil {
		return err
	}
	s.todoService = todos.NewTodoService(s.database)
	s.todoController = todos.NewController(s.todoService)

	// setup fiber
	engine := pug.New("./app/templates", ".pug")
	engine.Reload(true)
	s.app = fiber.New(fiber.Config{
		Views: engine,
	})
	s.app.Use(logger.New())

	// routes configuration
	s.app.Static("/", "./app/assets")
	s.app.Get("/", s.todoController.Index)

	todo := s.app.Group("/todos")
	todo.Get("/", s.todoController.List)
	todo.Post("/", s.todoController.Insert)

	todoId := todo.Group("/:id")
	todoId.Get("/", s.todoController.Find)
	todoId.Put("/", s.todoController.Update)
	todoId.Delete("/", s.todoController.Delete)

	return s.app.Listen(":3000")

}
