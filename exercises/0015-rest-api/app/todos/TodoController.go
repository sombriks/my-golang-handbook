package todos

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"strconv"
)

type TodoController struct {
	service *TodoService
}

func NewController(service *TodoService) *TodoController {
	var controller TodoController
	controller.service = service
	return &controller
}

func (controller *TodoController) Index(ctx fiber.Ctx) error {
	return controller.returnTodos(ctx, "index")
}

func (controller *TodoController) List(ctx fiber.Ctx) error {
	return controller.returnTodos(ctx, "todos/list")
}

func (controller *TodoController) Insert(ctx fiber.Ctx) error {
	newTodo := TodoFromRequest(ctx)
	id, err := controller.service.Insert(newTodo)
	if err != nil {
		return err
	}
	log.Infof("New ID: %d", id)
	return controller.returnTodos(ctx, "todos/list")
}

func (controller *TodoController) Find(ctx fiber.Ctx) error {
	qId := ctx.Params("id")
	id, err := strconv.Atoi(qId)
	if err != nil {
		return err
	}
	todo, err := controller.service.Find(uint(id))
	return ctx.Render("todos/detail", fiber.Map{"todo": todo})
}

func (controller *TodoController) Update(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return err
	}
	newTodo := TodoFromRequest(ctx)
	affected, err := controller.service.Update(id, newTodo)
	if err != nil {
		return err
	}
	log.Infof("Rows affected: %d", affected)
	return controller.returnTodos(ctx, "todos/list")
}

func (controller *TodoController) Delete(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return err
	}
	affected, err := controller.service.Delete(id)
	if err != nil {
		return err
	}
	log.Infof("Rows deleted: %d", affected)
	return controller.returnTodos(ctx, "todos/list")
}

func (controller *TodoController) returnTodos(ctx fiber.Ctx, template string) error {
	todos, err := controller.service.List(ctx.Query("q", ""))
	if err != nil {
		return err
	}
	return ctx.Render(template, fiber.Map{"todos": todos})
}
