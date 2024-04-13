package todos

import "github.com/gofiber/fiber/v3"

type TodoController struct {
}

func NewController() (*TodoController, error) {
	var controller TodoController

	return &controller, nil
}

func (controller *TodoController) Index(ctx fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{})
}
