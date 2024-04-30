package controllers

import (
	"0016-rest-json/app/services"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	var controller TodoController
	controller.service = service
	return &controller
}

func (c *TodoController) List(ctx echo.Context) error {
	return nil
}

func (c *TodoController) Insert(ctx echo.Context) error {
	return nil

}

func (c *TodoController) Find(ctx echo.Context) error {
	return nil

}

func (c *TodoController) Update(ctx echo.Context) error {
	return nil

}

func (c *TodoController) Delete(ctx echo.Context) error {
	return nil

}
