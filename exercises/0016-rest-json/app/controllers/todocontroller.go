package controllers

import (
	"0016-rest-json/app/services"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	service *services.TodoService
}

func NewTodoController(echo *echo.Echo, service *services.TodoService) *TodoController {
	var controller TodoController
	controller.service = service
	echo.GET("/todos", controller.List)
	echo.POST("/todos", controller.Insert)
	echo.GET("/todos/:id", controller.Find)
	echo.PUT("/todos/:id", controller.Update)
	echo.DELETE("/todos/:id", controller.Delete)
	return &controller
}

func (controller *TodoController) List(ctx echo.Context) error {
	q := ctx.QueryParam("q")
	result, err := controller.service.List(q)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(200, result)
	}
	return err
}

func (controller *TodoController) Insert(ctx echo.Context) error {
	return nil

}

func (controller *TodoController) Find(ctx echo.Context) error {
	return nil

}

func (controller *TodoController) Update(ctx echo.Context) error {
	return nil

}

func (controller *TodoController) Delete(ctx echo.Context) error {
	return nil

}
