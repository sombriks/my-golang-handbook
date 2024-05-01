package controllers

import (
	"0016-rest-json/app/models"
	"0016-rest-json/app/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
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
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200, result)
	}
	return err
}

func (controller *TodoController) Insert(ctx echo.Context) error {
	var newTodo models.Todo
	err := ctx.Bind(&newTodo)
	if err != nil {
		ctx.JSON(400, err.Error())
	}
	id, err := controller.service.Create(&newTodo)
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(201, id)
	}
	return err
}

func (controller *TodoController) Find(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}
	result, found, err := controller.service.Find(id)
	if !found {
		ctx.JSON(404, "Not found")
		return err
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}
	ctx.JSON(200, result)
	return nil
}

func (controller *TodoController) Update(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}
	var todo models.Todo
	err = ctx.Bind(&todo)
	if err != nil {
		ctx.JSON(400, err.Error())
	}
	count, err := controller.service.Update(id, &todo)
	if err != nil {
		ctx.JSON(500, err.Error())
	}
	ctx.JSON(200, fmt.Sprintf("%d rows updated", count))
	return nil
}

func (controller *TodoController) Delete(ctx echo.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}
	count, err := controller.service.Delete(id)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}
	ctx.JSON(200,fmt.Sprintf("%d rows deleted", count))
	return nil
}
